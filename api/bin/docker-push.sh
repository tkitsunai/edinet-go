#!/bin/bash

set -eu

IMAGE_NAME=tkitsunai/edinet-api-server
TAG=$(git rev-parse --short HEAD)
docker push ${IMAGE_NAME}:${TAG}