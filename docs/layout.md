---
layout: default
title: Project Layout
nav_order: 2
---

# Project Layout

The layout of this template is based on the [standard Go project layout](https://github.com/golang-standards/project-layout).

## `./build`

This folder contains scripts for building and packaging your DApp (i.e. macOS, Linux, Windows, etc) and generated executables. 

* [./build/package](../build/package) contains docker scripts.

* [./build/package](../build/native) is generated folder containing executable binaries. 

## `./cmd`

This folder contains skeleon codes cli executables entrypoint -- i.e. commands and arguments processor -- based on [Go corbra cli development framework](https://github.com/spf13/cobra). The default code is [dapp](../cmd/dapp).

## `./deployments`

This folder contains docker-compose scripts representing different deployment scenarios (development, unit testing, end-to-end testing, etc) in a local setting and/or in a cloud setting.

## `./examples`

This folder contains example codes demonstrating the use of Geth packages to create cli DApp

## `./internal`

This folder contains Go packages that will served as features and functionality (i.e. libraries) of your app. The principal packages include:

* [server](../internal/server/) embedding a webserver and a RESTFul server.

## `./scripts`

This folder contains these Bash scripts. 

* [./scripts/container.sh](../scripts/container.sh) to trigger the build process for docker-based application.

* [./scripts/dev.sh](../scripts/dev.sh) to trigger the build process for apps configure for development.

* [./scripts/native.sh](../scripts/native.sh) to trigger build processes to generate native (macOS, Linux and Windows) application.

## `./web`

This folder contains a skeleton ReactJS framework -- i.e. `./web/reactjs`.