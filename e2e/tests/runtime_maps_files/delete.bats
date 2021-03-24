#!/usr/bin/env bats
#
# Copyright 2019 HAProxy Technologies
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http:#www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

load '../../libs/dataplaneapi'
load "../../libs/get_json_path"
load './haproxy_config_setup'

@test "runtime_maps_files: Remove all map entries from the map file" {
    run dpa_curl DELETE "/services/haproxy/runtime/maps/mapfile1.map?force_delete=true&force_sync=true"
    assert_success

    dpa_curl_status_body '$output'
    assert_equal $SC 204
}

@test "runtime_maps_files: Return an error when trying to delete non existing runtime map file" {
    run dpa_curl DELETE "/services/haproxy/runtime/maps/not-exists.map"
    assert_success

    dpa_curl_status_body '$output'
    assert_equal $SC 404
}
