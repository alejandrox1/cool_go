# Working with Vault

Based off [How to run HashiCorp Vault (Secrets Management) in
Docker](https://www.melvinvivas.com/secrets-management-using-docker-hashicorp-vault/).
The commands will deviate somewhat from the tutorial as we are using the latest
docker image, `0.10.1`.

For your future self, who already knows whats going on:
```
vault operator login

# Do this 3 times.
vault operator unseal

vault login
```

To store secrets
```
vault write secret/postgresql_creds username=admin password=admin
```

To create policies
```
vault policy write common /policies/common.json

vault token create -policy="common"
```

Once you are done
```
vault operator seal
```

## Seting it up
First, initialise the vault and check the server's status by running the
follwoing commands inside the vault container (`make login`):
```
vault operator init

vault status
```

To start guarding secrets we must first unseal the vault,
```
/ # vault operator unseal
Unseal Key (will be hidden): 
Key                Value
---                -----
Seal Type          shamir
Sealed             true
Total Shares       5
Threshold          3
Unseal Progress    1/3
Unseal Nonce       d798b4fd-ac68-42d3-89f7-2eef5bcd6c43
Version            0.10.1
HA Enabled         true

/ # vault operator unseal
Unseal Key (will be hidden): 
Key                Value
---                -----
Seal Type          shamir
Sealed             true
Total Shares       5
Threshold          3
Unseal Progress    2/3
Unseal Nonce       d798b4fd-ac68-42d3-89f7-2eef5bcd6c43
Version            0.10.1
HA Enabled         true

/ # vault operator unseal
Unseal Key (will be hidden): 
Key             Value
---             -----
Seal Type       shamir
Sealed          false
Total Shares    5
Threshold       3
Version         0.10.1
Cluster Name    vault-cluster-a2e4c8bb
Cluster ID      73b80d87-532b-e38b-c88e-2adda3a348c6
HA Enabled      false

```

Now that the vault is unsealed, authenticate by using your `initial root
token`:
```
/ # vault login
Token (will be hidden): 
Success! You are now authenticated. The token information displayed below
is already stored in the token helper. You do NOT need to run "vault login"
again. Future Vault requests will automatically use this token.

Key                Value
---                -----
token              5a35ef65-32ff-cbdd-fa04-c215ef8a6add
token_accessor     c5242ac0-0174-f576-0bbb-bf1ddb941a33
token_duration     ∞
token_renewable    false
token_policies     [root]
```

## Storing secrets
```
/ # vault write secret/postgresql_creds username=user password=pass
Success! Data written to: secret/postgresql_creds

/ # vault read secret/postgresql_creds
Key                 Value
---                 -----
refresh_interval    10h
password            pass
username            user

/ # vault read -format=json secret/postgresql_creds
{
  "request_id": "b4c87e8a-8d8a-4afa-9c55-b856de3429ce",
  "lease_id": "",
  "lease_duration": 36000,
  "renewable": false,
  "data": {
    "password": "pass",
    "username": "user"
  },
  "warnings": null
}

/ # vault operator seal
Success! Vault is sealed.
```

## Policies
```
/ # vault policy write common /policies/common.json 
Success! Uploaded policy: common

/ # vault token create -policy="common"
Key                Value
---                -----
token              a2bd9dbc-ebc5-4817-2f4a-93721b1a89a2
token_accessor     368f6213-5da9-2bdb-259d-84aa455e4cb6
token_duration     10h
token_renewable    true
token_policies     [common default]

```
