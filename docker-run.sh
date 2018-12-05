#!/bin/sh

docker run \
    -d \
    -p 8080:8080 \
    --name edinet-go \
    tkitsunai/edinet-go:latest