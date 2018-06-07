FROM vault:0.10.1

WORKDIR /work
RUN mkdir config
ADD local.json config/

RUN chmod -R go=rwx /work 

CMD ["server", "-config=config/local.json"]
