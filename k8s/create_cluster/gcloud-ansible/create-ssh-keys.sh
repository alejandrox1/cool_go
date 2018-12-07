#!/usr/bin/env bash

mkdir -p ssh-keys/google_compute_engine

ssh-keygen -t rsa -b 4096 -f $PWD/ssh-keys/google_compute_engine/id_rsa -C "gcloud" -N ''
