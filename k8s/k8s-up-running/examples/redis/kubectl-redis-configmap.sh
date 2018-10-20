#!/usr/bin/env bash

kubectl create configmap \
    --from-file=slave.conf=./slave.conf \
    --from-file=master.conf=./master-conf \
    --from-file=sentinerl.conf=./sentinel.conf \
    --from-file=init.sh=./init.sh \
    --from-file=sentinel.sh=./sentinel.sh \
    redis-config
