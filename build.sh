#!/bin/bash

NAME=${1:-joelncornett/readit-ocr-server}

echo $NAME
docker rmi $NAME
docker build -t $NAME .
