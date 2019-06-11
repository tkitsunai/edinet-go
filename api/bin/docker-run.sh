#!/bin/bash

set -eu

${TAG:-latest}

docker run \
    -d \
    -p 8080:8080 \
    --name edinet-api-server \
    tkitsunai/edinet-api-server:${TAG}