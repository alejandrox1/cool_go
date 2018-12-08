FROM ubuntu:18.04

RUN apt-get update -y \
    && apt-get install -y fortune \
    && rm -rf /var/lib/apt/lists/* \
    && rm -rf /src/*.deb

ADD fortune-generator.sh /bin/fortune-generator.sh

ENTRYPOINT ["/bin/fortune-generator.sh"]
CMD ["15"]
