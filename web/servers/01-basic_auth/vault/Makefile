.PHONY: all clean

all:
	mkdir -p vault/config && \
		chmod g+rw vault/ && \
		docker-compose up --build

login:
	docker exec -it vault /bin/sh

clean:
	sudo rm -rf vault/
