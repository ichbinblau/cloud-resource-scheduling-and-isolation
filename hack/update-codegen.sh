#!/usr/bin/env bash

# Copyright 2024 Intel Corporation
# 
# Licensed under the Apache License, Version 2.0 (the License);
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
# 
#    http://www.apache.org/licenses/LICENSE-2.0
# 
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

# corresponding to go mod init <module>
MODULE=github.com/intel/cloud-resource-scheduling-and-isolation
# api package
API_PKG=pkg/api
# generated output package
OUTPUT_PKG=pkg/generated
# group-version such as foo:v1alpha1
GROUP_VERSION=diskio:v1alpha1

SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
CODEGEN_PKG=${CODEGEN_PKG:-$(cd "${SCRIPT_ROOT}"; ls -d -1 ./vendor/k8s.io/code-generator 2>/dev/null || echo ../code-generator)}

bash "${CODEGEN_PKG}"/generate-groups.sh \
  "deepcopy,client,informer,lister" \
  ${MODULE}/pkg/generated \
  ${MODULE}/pkg/api \
  diskio:v1alpha1 \
  --trim-path-prefix ${MODULE} \
  --output-base "./" \
  --go-header-file "${SCRIPT_ROOT}"/hack/boilerplate.go.txt
