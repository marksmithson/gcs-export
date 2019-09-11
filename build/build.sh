#!/bin/bash

pushd ../

docker build -f build/package/Dockerfile -t gcs-export .