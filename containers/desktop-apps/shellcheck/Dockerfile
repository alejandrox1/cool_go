FROM ubuntu:16.04
MAINTAINER Jorge Alarcon <alarcj137@gmail.com>


RUN apt-get update -y && apt-get install -y \
    file \
    shellcheck \
    --no-install-recommends && \
    rm -rf /var/lib/apt/lists/* && \
    rm -rf /src/*.deb

ENTRYPOINT [ "shellcheck" ]
