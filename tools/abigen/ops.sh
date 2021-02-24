#!/bin/bash

COMMAND="$1"

export ABIGEN_IMAGE=oc/abigen:current

message="$0 build | clean | gen "

function build(){
    docker build -f ./tools/abigen/Dockerfile -t $ABIGEN_IMAGE .
}

function clean(){
    docker rmi -f $ABIGEN_IMAGE
    docker rmi -f $(docker images --filter "dangling=true" -q)
    rm -rf ./internal/token/*
}

function gen(){
    docker run --rm -v $PWD/solidity/tron.sol:/opt/solidity/tron.sol -v $PWD/internal/token/trontoken/:/opt/tokens/ $ABIGEN_IMAGE --sol /opt/solidity/tron.sol --pkg trontoken --out=/opt/tokens/trontoken.go
}

if [ "$#" -ne 1 ]; then
    echo $message
    exit 1
fi

case $COMMAND in
    "build")
        build
        ;;
    "clean")
        clean
        ;;
    "gen")
        gen
        ;;
    *)
        echo $message
        ;;
esac
