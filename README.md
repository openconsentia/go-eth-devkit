[![Go Report Card](https://goreportcard.com/badge/github.com/openconsentia/geth-tools)](https://goreportcard.com/report/github.com/openconsentia/geth-tools)

# Overview

`go-eth-devkit` is a template for an Ethereum based native (macOS, Linux and Windows) Decentralize Application (DApp) with Go and ReactJS as the development framework.

## Prerequisite

* Clone [https://github.com/openconsentia/gosol](https://github.com/openconsentia/gosol). This is an adjunct to this devkit. Use it to generate Go binding for your Solidity contract.

* Docker.

* Go tools 1.16

## How to use this project?

* STEP 1: Click the button named `Use this template`.

* STEP 2: Rename the module in [go.mod](./go.mod).

* STEP 3: Rename this folder [./cmd/dapp](./cmd/goreact) to one of your choice (`./cmd/<your-choice>`).

* STEP 4: Modify these scripts:
    * [./build/package/builder.yaml](./build/package/builder.yaml);
    * [./deployments/dev/docker-compose.yaml](./deployments/dev/docker-compose.yaml).

* STEP 5: Extends the [./web/reactjs](./web/reactjs).

* STEP 6: Extend the Go codes in `./cmd/<your-choice>` and [./internal](./internal).

* STEP 7: Build your project.

* STEP 8: To see you built artefacts (native or container) apps in action.

## Copyright Notice

Copyright 2020 Open Consentia Contributors.

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.