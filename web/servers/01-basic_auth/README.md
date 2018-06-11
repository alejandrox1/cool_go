# Simple HTTP server with basicAuth 

## Getting Started
We are abusing GNUMake as a command-line utility here, so to start this thing
up we need to begin by running the vault container. Vault will contain all the
secrets and important information for all other services.
```
make vault
```

For details on how to the vault onfigured see [vault/README](vault/). The idea
is that you will have to login into the vault container, unseal it, login,
store your secrets, etc. 
To login into the vault container,
```
make vaultlogin
```

Once Vault in configured you can spin up the database instance along with
adminer by running
```
make db
```
you can check your database on `http://localhost:8080`.

Finally, to spin up the golang server, simply run:
```
make server
```
For more details see [server/README](server/).


## Notes
In case you have `postgreSQL` already running on your system then you may
either change the `ports`on `docker-compose.yaml` or 
```bash
sudo fuser -n tcp -k 5432
```

To run application simply run:
```
docker-compose up --build [-d]
```
and then visit `http://localhost:8080` to check `adminer` out.


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
