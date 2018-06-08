.PHONY: vault vaultstop vaultstart vaultlogin db server clean

vault:
	mkdir -p vault/vault/config && \
		chmod g+rw vault/vault/ && \
		docker-compose up --build vault

vaultstop:
	docker-compose stop vault

vaultstart:
	docker-compose start vault

vaultlogin:
	docker exec -it vault /bin/sh

db:
	docker-compose up --build db adminer

server:
	( cd server; CGO_ENABLED=0 go build -a -installsuffix cgo -o $@ ) && \
		docker-compose up --build server

clean:
	sudo rm -rf vault/vault/
