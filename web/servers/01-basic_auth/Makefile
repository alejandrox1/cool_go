.PHONY: vault vaultstop vaultstart vaultlogin db server clean

vault:
	mkdir -p vault/vault/config && \
		chmod g+rw vault/vault/ && \
		docker-compose up -d --build vault

vaultstop:
	docker-compose stop vault

vaultstart:
	docker-compose start vault

vaultlogin:
	docker exec -it vault /bin/sh

db:
	docker-compose up -d --build db adminer

dblogin:
	docker exec -it db /bin/sh

server:
	( cd server; CGO_ENABLED=0 go build -a -installsuffix cgo -o $@ ) && \
		docker-compose up --build server

clean:
	sudo rm -rf vault/vault/ && \
		rm -rf server/server && \
		docker-compose down
