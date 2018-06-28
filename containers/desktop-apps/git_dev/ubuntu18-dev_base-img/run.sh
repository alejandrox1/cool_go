#!/bin/bash

CONTAINER="alejandrox1/ubuntu18-dev"

docker build --no-cache --force-rm -t $CONTAINER . && docker push $CONTAINER
