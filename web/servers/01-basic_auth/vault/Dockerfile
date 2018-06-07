FROM vault:0.10.1

RUN mkdir config
ADD ./configs/local.json config/

CMD ["server", "-config=config/local.json"]
