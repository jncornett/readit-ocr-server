#!/bin/bash

NAME=$(date +%s)
docker rmi $NAME
docker build -t $NAME . && docker run -it $NAME /bin/bash
