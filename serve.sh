#!/bin/bash

./build.sh

PORT=$1

docker run \
    -p $PORT:$PORT \
    --rm \
    --env LISTEN=:$PORT \
    -it \
    joelncornett/readit-ocr-server \
    readit-ocr-server
