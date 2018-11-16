#!/bin/sh

docker run \
    -d \
    -p 8080:8080 \
    --name edinet-goapi \
    tkitsunai/edinet-goapi:latest