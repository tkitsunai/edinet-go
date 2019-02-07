#!/bin/sh

IMAGE_NAME=tkitsunai/edinet-go
TAG=$(git rev-parse --short HEAD)
docker build \
  -t ${IMAGE_NAME}:${TAG} \
  -t ${IMAGE_NAME}:latest \
  .
