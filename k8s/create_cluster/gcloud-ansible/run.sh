#!/usr/bin/env bash


docker build --force-rm -t gcloud . \
    && docker run --rm -it \
    -v $PWD:/root/compute-video-demo-ansible \
    -v $PWD/ssh-keys:/root/.ssh \
    gcloud bash
