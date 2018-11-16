#!/bin/sh

IMAGE_NAME=tkitsunai/edinet-goapi
TAG=$(git rev-parse --short HEAD)
docker push ${IMAGE_NAME}:${TAG}