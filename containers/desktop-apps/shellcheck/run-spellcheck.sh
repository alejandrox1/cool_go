#!/bin/bash
#
# SHELLCHECK all *.sh files in the current working directory.
#
# Another good utility to eventually put along with this would be 'file'.
#
set -e

GRE="\e[34;1;5m"
NOC="\e[0m"

IMG="shellchecker"
TAG="docker.io/alejandrox1/${IMG}:latest"

# Remove all exited shellcheck containers.
docker ps -aq -f name=shellchecker* -f status=exited | xargs -r docker rm

# Build your container.
docker build -t "${IMG}" . 

# If passed "push" then push a tagged image to a registry.
if [[ $# == 1 && $1 == "push" ]]; then
    echo -e "${GRE}Building... Tagging... and Pushing image...${NOC}" && \
        docker tag "${IMG}" "${TAG}" && \
        docker push "${TAG}"
fi

# Shellcheck all *.sh files in the current working directory.
for sh in *; do
    if [[ $(file -b "${sh}" | awk '{print $2}') == "shell" ]]; then
        echo -e "${GRE}Shellcheking ${sh}...${NOC}" && \
            docker run --name "${IMG}-${sh}" \
            -v $(pwd):/home \
            shellchecker "${sh}"
    fi
done
