#!/bin/bash

CONTAINER="alejandrox1/ubuntu18_python-dev"

docker build --no-cache --force-rm -t $CONTAINER . && docker push $CONTAINER
