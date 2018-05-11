#!/bin/bash

set -e

GRE="\e[32m"
NOC="\e[0m"

IMG="ubuntu-nvidia"
TAG="docker.io/alejandrox1/ubuntu-nvidia:latest"

docker build -t "${IMG}" .


if [[ "$#" == 1 && "$1" == "build" ]]; then
    echo -e "${GRE}Building... Tagging... and Pushing image...${NOC}" && \
        docker tag "${IMG}" "${TAG}" && \
        docker push "${TAG}"
fi

docker run -it \
    -v $(pwd):/home \
    --device /dev/nvidia0:/dev/nvidia0 \
    --device /dev/nvidiactl:/dev/nvidiactl \
    --device /dev/nvidia-uvm:/dev/nvidia-uvm \
    --device /dev/nvidia-uvm-tools:/dev/nvidia-uvm-tools \
    ubuntu-nvidia /bin/bash

