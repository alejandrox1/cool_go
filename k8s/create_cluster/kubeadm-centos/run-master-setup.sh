#!/usr/bin/env bash

ansible-playbook \
    --private-key=$HOME/.ssh/gcloud_vm \
    master-node.yml