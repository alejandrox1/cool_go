IMAGE := alejandrox1/fortune:v1.2.0

.PHONY: all build push

all: push

build:
	docker build --force-rm --no-cache -t $(IMAGE) .

push: build
	docker push $(IMAGE)
