#!/usr/bin/env bash

docker build --force-rm -t gcloud . \
    && docker run --rm -it -v $PWD/ssh-keys:/root/ssh-keys gcloud bash
