#!/usr/bin/env bash
#
# Copyright 2020 HAProxy Technologies
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

# dpa_curl is going to return the response status code along with the body
# these values can be easily read as following
#
# read -r SC BODY < <(api GET /services/haproxy/runtime/info)
# echo "Status Code: ${SC}"
# echo "Body: ${BODY}"
#
# Arguments:
# - HTTP verb
# - original URL
# - HTTP POST data
function dpa_curl() {
  verb=$1; shift
  endpoint=$1; shift
  data="@${BATS_TEST_DIRNAME}/$1";
  if [ -z "$1" ]; then
    data="/dev/null"
  fi
  response=$(curl -m 10 -s -H 'content-type: application/json' --user dataplaneapi:mypassword "-X${verb}" -w "\n%{http_code}" "-d${data}" "http://${LOCAL_IP_ADDRESS}:${E2E_PORT}${BASE_PATH}${endpoint}")
  status_code=$(tail -n1 <<< "$response")
  response=$(sed '$ d' <<< "$response")
  echo "$status_code $response"
}

# dpa_curl_text_plain behaves in same manner as api, but uses content-type: text/plain
function dpa_curl_text_plain() {
  verb=$1; shift
  endpoint=$1; shift
  data=${1:-"/dev/null"}
  response=$(curl -m 10 -s -H 'content-type: text/plain' --user dataplaneapi:mypassword "-X${verb}" -w "\n%{http_code}" --data-binary "${data}" "http://${LOCAL_IP_ADDRESS}:${E2E_PORT}${BASE_PATH}${endpoint}")
  status_code=$(tail -n1 <<< "$response")
  response=$(sed '$ d' <<< "$response")
  echo "$status_code $response"
}

function dpa_curl_file_upload() {
  verb=$1; shift
  endpoint=$1; shift
  data=${1:-"/dev/null"}
  response=$(curl -m 10 -s -H "Content-type: multipart/form-data" --user dataplaneapi:mypassword "-X${verb}" -w "\n%{http_code}" --form "file_upload=${data}" "http://${LOCAL_IP_ADDRESS}:${E2E_PORT}${BASE_PATH}${endpoint}")
  status_code=$(tail -n1 <<< "$response")
  response=$(sed '$ d' <<< "$response")
  echo "$status_code $response"
}

# function dpa_curl_download returns values differently to allow for multiline body contents, it should be used as follows:
# local BODY;
# local SC;
# dpa_curl_download GET "/services/haproxy/storage/maps/mapfile_example.map"
# echo "Status Code: ${SC}"
# echo "Body: ${BODY}"
function dpa_curl_download() {
  verb=$1; shift
  endpoint=$1; shift
  data=${1:-"/dev/null"}
  response=$(curl -m 10 -v -s -H 'content-type: application/json' --user dataplaneapi:mypassword "-X${verb}" -w "\n%{http_code}" "-d${data}" "http://${LOCAL_IP_ADDRESS}:${E2E_PORT}${BASE_PATH}${endpoint}" 2>/tmp/headers)
  #echo "$status_code $body"
  status_code=$(tail -n1 <<< "$response")
  body=$(head -n -1 <<< "$response")
  eval SC="'$status_code'"
  eval BODY="'$body'"
}

# version return the current HAProxy configuration file version, useful to
# avoid keeping track of it at each POST/PUT call.
#
# Any argument is required.
#
# Example:
# version
# >>> 10
function version() {
  read -r SC RES < <(dpa_curl GET "/services/haproxy/configuration/global")
  V="$(RES=${RES} jq -n 'env.RES | fromjson | ._version')"
  echo "$V"
}
