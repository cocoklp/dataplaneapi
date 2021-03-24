// Copyright 2019 HAProxy Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package handlers

import (
	"fmt"
	"strconv"

	"github.com/go-openapi/runtime/middleware"
	"github.com/google/renameio"
	client_native "github.com/haproxytech/client-native/v2"
	"github.com/haproxytech/client-native/v2/models"

	"github.com/haproxytech/dataplaneapi/configuration"
	"github.com/haproxytech/dataplaneapi/haproxy"
	"github.com/haproxytech/dataplaneapi/operations/cluster"
	"github.com/haproxytech/dataplaneapi/operations/discovery"
)

// CreateClusterHandlerImpl implementation of the CreateClusterHandler interface
type CreateClusterHandlerImpl struct {
	Client      *client_native.HAProxyClient
	Config      *configuration.Configuration
	ReloadAgent haproxy.IReloadAgent
}

// GetClusterHandlerImpl implementation of the GetClusterHandler interface
type GetClusterHandlerImpl struct {
	Config *configuration.Configuration
}

// ClusterInitiateCertificateRefreshHandlerImpl implementation of the ClusterInitiateCertificateRefreshHandler interface
type ClusterInitiateCertificateRefreshHandlerImpl struct {
	Config *configuration.Configuration
}

// Handle executing the request and returning a response
func (h *ClusterInitiateCertificateRefreshHandlerImpl) Handle(params cluster.InitiateCertificateRefreshParams, principal interface{}) middleware.Responder {
	if h.Config.Mode.Load() != "cluster" {
		return cluster.NewInitiateCertificateRefreshForbidden()
	}
	h.Config.Notify.CertificateRefresh.Notify()
	return cluster.NewInitiateCertificateRefreshOK()
}

func (h *CreateClusterHandlerImpl) err500(err error, transaction *models.Transaction) middleware.Responder {
	if transaction != nil {
		_ = h.Client.Configuration.DeleteTransaction(transaction.ID)
	}
	msg := err.Error()
	code := int64(500)
	return cluster.NewPostClusterDefault(500).WithPayload(&models.Error{
		Code:    &code,
		Message: &msg,
	})
}

func (h *CreateClusterHandlerImpl) Handle(params cluster.PostClusterParams, principal interface{}) middleware.Responder {
	key := h.Config.Cluster.BootstrapKey.Load()
	if params.Data.BootstrapKey != "" && key != params.Data.BootstrapKey {
		h.Config.Mode.Store("cluster")
		h.Config.Cluster.BootstrapKey.Store(params.Data.BootstrapKey)
		h.Config.Cluster.Clear()
		h.Config.Notify.BootstrapKeyChanged.Notify()
	}
	if params.Data.Mode == "single" && h.Config.Mode.Load() != params.Data.Mode {
		// by default we are cleaning configuration
		if params.Configuration == nil || *params.Configuration != "keep" {
			version, errVersion := h.Client.Configuration.GetVersion("")
			if errVersion != nil || version < 1 {
				// silently fallback to 1
				version = 1
			}
			transaction, err := h.Client.Configuration.StartTransaction(version)
			if err != nil {
				return h.err500(err, transaction)
			}
			// delete backends
			_, backends, err := h.Client.Configuration.GetBackends(transaction.ID)
			if err != nil {
				return h.err500(err, transaction)
			}
			for _, backend := range backends {
				err = h.Client.Configuration.DeleteBackend(backend.Name, transaction.ID, 0)
				if err != nil {
					return h.err500(err, transaction)
				}
			}
			// delete all frontends
			_, frontends, err := h.Client.Configuration.GetFrontends(transaction.ID)
			if err != nil {
				return h.err500(err, transaction)
			}
			for _, frontend := range frontends {
				err = h.Client.Configuration.DeleteFrontend(frontend.Name, transaction.ID, 0)
				if err != nil {
					return h.err500(err, transaction)
				}
			}

			// now create dummy frontend so haproxy does not complain
			err = h.Client.Configuration.CreateFrontend(&models.Frontend{Name: "disabled"}, transaction.ID, 0)
			if err != nil {
				return h.err500(err, transaction)
			}
			err = h.Client.Configuration.CreateBind("disabled", &models.Bind{
				Name:    "tmp",
				Address: fmt.Sprintf("/tmp/dataplaneapi-%s.sock", h.Config.Name.Load()),
			}, transaction.ID, 0)
			if err != nil {
				return h.err500(err, transaction)
			}
			// now reset peer-id
			if h.Config.HAProxy.NodeIDFile != "" {
				err = renameio.WriteFile(h.Config.HAProxy.NodeIDFile, []byte("localhost"), 0644)
				if err != nil {
					return h.err500(err, transaction)
				}
				_, peerSections, errPeers := h.Client.Configuration.GetPeerSections(transaction.ID)
				if errPeers != nil {
					return h.err500(errPeers, transaction)
				}
				peerFound := false
				dataplaneID := h.Config.Cluster.ID.Load()
				for _, section := range peerSections {
					_, peerEntries, errPeersEntries := h.Client.Configuration.GetPeerEntries(section.Name, transaction.ID)
					if errPeersEntries != nil {
						return h.err500(errPeersEntries, transaction)
					}
					for _, peer := range peerEntries {
						if peer.Name == dataplaneID {
							peerFound = true
							peer.Name = "localhost"
							errPeerEntry := h.Client.Configuration.EditPeerEntry(dataplaneID, section.Name, peer, transaction.ID, 0)
							if errPeerEntry != nil {
								return h.err500(errPeerEntry, transaction)
							}
						}
					}
				}
				if !peerFound && dataplaneID != "" {
					return h.err500(fmt.Errorf("peer [%s] not found in HAProxy config", dataplaneID), transaction)
				}
			}
			_, err = h.Client.Configuration.CommitTransaction(transaction.ID)
			if err != nil {
				return h.err500(err, nil)
			}
			// we need to restart haproxy
			err = h.ReloadAgent.Restart()
			if err != nil {
				return h.err500(err, nil)
			}
		}

		h.Config.Cluster.BootstrapKey.Store("")
		h.Config.Mode.Store(params.Data.Mode)
		h.Config.Status.Store("active")
		h.Config.Cluster.Clear()
		defer func() {
			h.Config.Notify.Reload.Notify()
		}()
	}
	err := h.Config.Save()
	if err != nil {
		return h.err500(err, nil)
	}
	result := models.ClusterSettings{
		BootstrapKey: h.Config.Cluster.BootstrapKey.Load(),
		Mode:         h.Config.Mode.Load(),
		Status:       h.Config.Status.Load(),
	}
	return cluster.NewPostClusterOK().WithPayload(&result)
}

// Handle executing the request and returning a response
func (h *GetClusterHandlerImpl) Handle(params discovery.GetClusterParams, principal interface{}) middleware.Responder {
	portStr := h.Config.Cluster.Port.Load()
	p, err := strconv.Atoi(portStr)
	if err != nil {
		p = 0
	}
	port := int64(p)
	var clusterSettings *models.ClusterSettingsCluster
	if h.Config.Mode.Load() == "cluster" {
		clusterSettings = &models.ClusterSettingsCluster{
			Address:     h.Config.Cluster.URL.Load(),
			Port:        &port,
			APIBasePath: h.Config.Cluster.APIBasePath.Load(),
			Name:        h.Config.Cluster.Name.Load(),
			Description: h.Config.Cluster.Description.Load(),
		}
	}
	settings := &models.ClusterSettings{
		BootstrapKey: h.Config.Cluster.BootstrapKey.Load(),
		Cluster:      clusterSettings,
		Mode:         h.Config.Mode.Load(),
		Status:       h.Config.Status.Load(),
	}
	return discovery.NewGetClusterOK().WithPayload(settings)
}
