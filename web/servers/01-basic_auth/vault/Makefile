.PHONY: all clean

all:
	mkdir vault && \
		chmod g+rw vault/ && \
		docker-compose up --build

login:
	docker exec -it vault /bin/sh

clean:
	sudo rm -rf vault/
