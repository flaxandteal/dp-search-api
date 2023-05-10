#!/bin/bash -eux

cwd=$(pwd)

export GOPATH=$cwd/go

pushd dp-search-api
  make build && mv build/$(go env GOOS)-$(go env GOARCH)/* ../
  cp Dockerfile.concourse ../
popd