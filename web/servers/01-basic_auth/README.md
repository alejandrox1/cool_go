# Simple HTTP server with basicAuth 

In case you have `postgreSQL` already running on your system then you may
either change the `ports`on `docker-compose.yaml` or 
```bash
sudo fuser -n tcp -k 5432
```

To run application simply run:
```
docker-compose up --build [-d]
```
and then visit `localhost:8080` to check `adminer` out.
