FROM vault:0.10.1

WORKDIR /work
RUN mkdir config
ADD local.json config/

CMD ["server", "-config=config/local.json"]
