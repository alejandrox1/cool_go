FROM ubuntu:18.04
MAINTAINER Jorge Alarcon <alarcj137@gmail.com>


RUN apt-get update -y && apt-get install -y \
    shellcheck \
    --no-install-recommends && \
    rm -rf /var/lib/apt/lists/* && \
    rm -rf /src/*.deb

WORKDIR /home

ENTRYPOINT [ "shellcheck" ]
