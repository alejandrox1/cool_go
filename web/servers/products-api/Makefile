-include db/config/conf.mk
.EXPORT_ALL_VARIABLES:

.PHONY: all db dblogin server clean


all: 
	docker-compose up -d --build

db:
	docker-compose up -d --build db adminer

dblogin:
	docker exec -it db /bin/sh

server:
	docker-compose up --build server

tests: db
	docker-compose run --rm server go test -v -cover

clean:
	docker rmi -f products-api_server
	docker-compose down -v
