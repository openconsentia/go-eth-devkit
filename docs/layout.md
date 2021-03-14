# Project Layout

The layout of this project is based on the [standard Go project layout](https://github.com/golang-standards/project-layout).

## `./build`

This folder contains artefacts for building and packaging your DApp (i.e. macOS, Linux, Windows, etc) for deployment. Here you will find builders and packager mainly in the form of docker scripts [./build/package](../build/package).

When you execute the packager you will find your DApps in this folder [./build/package](../build/native), which are not checked into the remote repo. 

## `./cmd`

This folder contains source code create to (i.e. main entry point for) cli applications known as [dapp](../cmd/dapp) which embed ReactJS and Go based decentralize application. The codes integrate by chosing appropriate (via configuration logic) with the library codes found in `internal` folder.

In this folder, you can create different types of application to server different scenarios.

## `./deployments`

This folder contains docker-compose scripts representing different deployment scenarios (development, unit testing, end-to-end testing, etc) in a local setting and/or in a cloud setting.

## `./examples`

This folder contains example codes demonstrating the use of Go packages to create cli DApp

## `./internal`

This folder contains Go packages that will served as features and functionality (i.e. libraries) for packaging with the codes in `cmd` folder. The principal packages include:

* [server](../internal/server/) embedding a webserver and a RESTFul server;
* [infurawrap](../internal/infura) wrapper for infura API;
* [gethwrap](../internal/gethwrap) wrapper for geth based API;
* [shared](../internal/shared) shared models, interfaces and utilities for all Ethereum SDK API.

## `./scripts`

This folder containers Bash scripts to trigger build processes and to execute various deployment scenarios. Out-of-the-box scaffold you will find the follow:

* [./scripts/container.sh](../scripts/container.sh) to trigger the build process for docker-based application;
* [./scripts/dev.sh](../scripts/dev.sh) to trigger the build process for apps configure for development;
* [./scripts/native.sh](../scripts/native.sh) to trigger build processes to generate native (macOS, Linux and Windows) application.

## `./web`

This folder contains codes to create UI artefacts. In this context there is a basic `./web/reactjs`. You can add more if you wish.