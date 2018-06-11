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

## Vault
* `vault/configs`
  * `local.json` - configuration file for the vault itself.
  * `vault.json` - configuration file for the server (see `server/`).

* `vault/policies`
  * `common.json` - policy to `list` and `read`.


## Database
* `db/config` - Go into the Dockerfile in `db/` and add whatever sql scripts
  you want.
  * `createDB.sql` - SQL script to create a table.
