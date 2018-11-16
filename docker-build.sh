#!/bin/sh

IMAGE_NAME=tkitsunai/edinet-goapi
TAG=$(git rev-parse --short HEAD)
docker build -t ${IMAGE_NAME}:${TAG} .