#!/bin/bash
set -e
set -o pipefail
export blue="\e[1;34m"                                                         
export green="\e[92;1m"                                                           
export red="\e[1;31m"
export reset="\e[0m"

CONTAINER="ubuntu18-dev"


docker build --no-cache --force-rm -t $CONTAINER .
