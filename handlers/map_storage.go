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
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	client_native "github.com/haproxytech/client-native/v2"
	models "github.com/haproxytech/client-native/v2/models"

	"github.com/haproxytech/dataplaneapi/haproxy"
	"github.com/haproxytech/dataplaneapi/misc"
	"github.com/haproxytech/dataplaneapi/operations/storage"
)

// StorageCreateRuntimeMapHandlerImpl implementation of the StorageCreateRuntimeMapHandler interface using client-native client
type StorageCreateRuntimeMapHandlerImpl struct {
	Client *client_native.HAProxyClient
}

func (h *StorageCreateRuntimeMapHandlerImpl) Handle(params storage.CreateRuntimeMapParams, principal interface{}) middleware.Responder {
	file, ok := params.FileUpload.(*runtime.File)
	if !ok {
		return storage.NewCreateRuntimeMapBadRequest()
	}

	filename, err := h.Client.MapStorage.Create(file.Header.Filename, params.FileUpload)
	if err != nil {
		status := misc.GetHTTPStatusFromErr(err)
		return storage.NewCreateRuntimeMapDefault(status).WithPayload(misc.SetError(status, err.Error()))
	}

	me := &models.Map{
		Description: "managed but not loaded map file (no runtime ID)",
		File:        filename,
		StorageName: filepath.Base(filename),
	}
	// no reload or force reload since this is just a file upload,
	// haproxy configuration has not been changed
	return storage.NewCreateRuntimeMapCreated().WithPayload(me)
}

// GetMapStorageHandlerImpl implementation of the StorageGetAllStorageMapFilesHandler interface
type GetAllStorageMapFilesHandlerImpl struct {
	Client *client_native.HAProxyClient
}

// Handle executing the request and returning a response
func (h *GetAllStorageMapFilesHandlerImpl) Handle(params storage.GetAllStorageMapFilesParams, principal interface{}) middleware.Responder {
	tempMaps := map[string]*models.Map{}

	// get filenames for files in storage
	filenames, err := h.Client.MapStorage.GetAll()
	if err != nil {
		e := misc.HandleError(err)
		return storage.NewGetAllStorageMapFilesDefault(int(*e.Code)).WithPayload(e)
	}

	for _, f := range filenames {
		tempMaps[f] = &models.Map{
			Description: "managed but not loaded map file (no runtime ID)",
			File:        f,
			ID:          "",
			StorageName: filepath.Base(f),
		}
	}

	// get Map model instances for runtime-loaded files
	runtimeMaps, err := h.Client.Runtime.ShowMaps()
	if err != nil {
		status := misc.GetHTTPStatusFromErr(err)
		return storage.NewGetAllStorageMapFilesDefault(status).WithPayload(misc.SetError(status, err.Error()))
	}

	// update (overwrite) info for on-disk files with runtime info
	for _, m := range runtimeMaps {
		// files outside of MapsDir are not managed, so shouldn't be returned
		if strings.HasPrefix(filepath.Dir(m.File), h.Client.Runtime.MapsDir) {
			tempMaps[m.File] = m
		}
	}

	// convert to a list to return
	var retMaps []*models.Map
	for _, v := range tempMaps {
		retMaps = append(retMaps, v)
	}

	return &storage.GetAllStorageMapFilesOK{Payload: retMaps}
}

// StorageGetOneStorageMapHandlerImpl implementation of the StorageGetOneStorageMapHandler interface
type GetOneStorageMapHandlerImpl struct {
	Client *client_native.HAProxyClient
}

func (h *GetOneStorageMapHandlerImpl) Handle(params storage.GetOneStorageMapParams, principal interface{}) middleware.Responder {
	filename, err := h.Client.MapStorage.Get(params.Name)
	if err != nil {
		e := misc.HandleError(err)
		return storage.NewGetOneStorageMapDefault(int(*e.Code)).WithPayload(e)
	}
	if filename == "" {
		return storage.NewGetOneStorageMapNotFound()
	}
	f, err := os.Open(filename)
	if err != nil {
		e := misc.HandleError(err)
		return storage.NewGetOneStorageMapDefault(int(*e.Code)).WithPayload(e)
	}
	return storage.NewGetOneStorageMapOK().WithPayload(f)
}

// StorageDeleteStorageMapHandlerImpl implementation of the StorageDeleteStorageMapHandler interface
type StorageDeleteStorageMapHandlerImpl struct {
	Client *client_native.HAProxyClient
}

func (h *StorageDeleteStorageMapHandlerImpl) Handle(params storage.DeleteStorageMapParams, principal interface{}) middleware.Responder {
	runningConf := strings.NewReader(h.Client.Configuration.Parser.String())

	filename, err := h.Client.MapStorage.Get(params.Name)
	if err != nil {
		e := misc.HandleError(err)
		return storage.NewDeleteStorageSSLCertificateDefault(int(*e.Code)).WithPayload(e)
	}

	// this is far from perfect but should provide a basic level of protection
	scanner := bufio.NewScanner(runningConf)

	lineNr := 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.Contains(line, filename) && !strings.HasPrefix(line, "#") {
			errCode := misc.ErrHTTPConflict
			errMsg := fmt.Sprintf("rejecting attempt to delete file %s referenced in haproxy conf at line %d: %s", filename, lineNr-1, line)
			e := &models.Error{Code: &errCode, Message: &errMsg}
			return storage.NewDeleteStorageSSLCertificateDefault(int(*e.Code)).WithPayload(e)
		}
		lineNr++
	}

	err = h.Client.MapStorage.Delete(params.Name)
	if err != nil {
		e := misc.HandleError(err)
		return storage.NewDeleteStorageMapDefault(int(*e.Code)).WithPayload(e)
	}
	return storage.NewDeleteStorageMapNoContent()
}

// StorageReplaceStorageMapFileHandlerImpl implementation of the StorageReplaceStorageMapFileHandler interface
type StorageReplaceStorageMapFileHandlerImpl struct {
	Client      *client_native.HAProxyClient
	ReloadAgent haproxy.IReloadAgent
}

func (h *StorageReplaceStorageMapFileHandlerImpl) Handle(params storage.ReplaceStorageMapFileParams, principal interface{}) middleware.Responder {
	_, err := h.Client.MapStorage.Replace(params.Name, params.Data)
	if err != nil {
		e := misc.HandleError(err)
		return storage.NewReplaceStorageMapFileDefault(int(*e.Code)).WithPayload(e)
	}

	if *params.ForceReload {
		err := h.ReloadAgent.ForceReload()
		if err != nil {
			e := misc.HandleError(err)
			return storage.NewReplaceStorageMapFileDefault(int(*e.Code)).WithPayload(e)
		}
		return storage.NewReplaceStorageMapFileNoContent()
	}
	rID := h.ReloadAgent.Reload()
	return storage.NewReplaceStorageMapFileAccepted().WithReloadID(rID)
}
