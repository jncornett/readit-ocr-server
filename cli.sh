#!/bin/bash

NAME=$(date +%s)

./build.sh $NAME

docker run -it $NAME /bin/bash
