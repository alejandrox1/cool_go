#!/usr/bin/env bash

mkdir -p ssh-keys

ssh-keygen -t rsa -b 4096 -f $PWD/ssh-keys/id_rsa -C "gcloud" -N ''
