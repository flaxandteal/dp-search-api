#!/bin/bash -eux

cwd=$(pwd)

export GOPATH=$cwd/go
echo $(pwd)
pushd dp-search-api
  make build && mv build/$(go env GOOS)-$(go env GOARCH)/* $cwd/build
  cp Dockerfile.concourse $cwd/build
  echo $cwd
  ls $cwd
  echo $cwd/build
  ls $cwd/build
popd

echo $(pwd)
ls $(pwd)
echo $(pwd)/build
ls $(pwd)/build
