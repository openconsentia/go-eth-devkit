#!/bin/bash

# Copyright 2020 The Open Consentia Contributors
# 
# Licensed under the Apache License, Version 2.0 (the "License");
#  you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

go mod download && \
go build -o /usr/local/bin/${REST_SERVER_NAME} ./cmd/${REST_SERVER_NAME}

/usr/local/bin/${REST_SERVER_NAME} start noui -p 9000