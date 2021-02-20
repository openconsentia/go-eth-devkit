# Project Layout

The layout of this project is based on the [standard Go project layout](https://github.com/golang-standards/project-layout).

## `./build`

This folder contains the following:

* [./build/package](../build/package) folder containers build scripts in the form of docker and docker-compose files to produce native macOS, Linux and Windows app or a Docker-based image;
* [./build/package/go-rice.sh](../build/package/go-rice.sh) script to pre-generate Go source files embedding the ReactJS artefacts in byte code form. DO NOT REMOVE THIS FILE.

If you wish to extend the build process please to do it in this folder [./build/package](../build/package).

When you execute the build scripts you will find a generated folder call `./build/native`, where you find native versions of executables (macOS, Linux and Windows).

## `./cmd`

This folder contains source code create the following cli applications known as [dapp](../cmd/dapp) a basic ReactJS and Go based decentralize application. 

## `./deployments`

This folder contains docker-compose networks to generate:

* a [local deployment network](../deployments/dev) to support UI and REST integration development effort;
* a [unit testing network](../deployments/unit) to support unit testing.

## `./internal`

This folder contains these Go packages:

* [server](../internal/server/) embedding a webserver and a RESTFul server;
* [infurawrap](../internal/infura) wrapper for infura API;
* [gethwrap](../internal/gethwrap) wrapper for geth based API;
* [shared](../internal/shared) shared models, interfaces and utilities for all Ethereum SDK API.

## `./web`

This folder contains a skeleton RectJS codes found in the sub-folder `./web/reactjs`.

## `./scripts`

This folder containers Bash scripts to trigger build processes and to execute various deployment scenarios. Out-of-the-box scaffold you will find the follow:

* [./scripts/container.sh](../scripts/container.sh) to trigger the build process for docker-images;
* [./scripts/dev.sh](../scripts/dev.sh) to trigger the build process for apps configure for development activities;
* [./scripts/native.sh](../scripts/native.sh) to trigger build processes to generate container apps.

## `./examples`

This folder contains example codes demonstrating the use of Go packages.

