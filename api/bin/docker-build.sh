#!/bin/bash

set -eu

SCRIPT_DIR=$(cd $(dirname $0); pwd)

cd ${SCRIPT_DIR}
cd ../

IMAGE_NAME=tkitsunai/edinet-api-server
TAG=$(git rev-parse --short HEAD)
docker build --rm \
  -t ${IMAGE_NAME}:${TAG} \
  -t ${IMAGE_NAME}:latest \
  .
