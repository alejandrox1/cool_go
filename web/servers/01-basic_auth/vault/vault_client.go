package main

import (
    "encoding/json"
    "fmt"
    "os"

    "github.com/hashicorp/vault/api"
)

type Vault struct {
    Token string `json:"vault_token"`
    Addr  string `json:"vault_address"`
}

func readVaultConfig(filename string) (vaultConfig Vault, err error) {
    contents, err := os.Open(filename)
    if err != nil {
        fmt.Printf("readVaultConfig error opening file: %s\n", err.Error())
        return
    }

    err = json.NewDecoder(contents).Decode(&vaultConfig)
    if err != nil {
        fmt.Printf("readVaultConfig error decoding json: %s\n", err.Error())
        return
    }

    return
}


func main() {
    vaultConfig, err := readVaultConfig("./configs/vault.json")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    client, err := api.NewClient(&api.Config{
        Address: vaultConfig.Addr,
    })
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    client.SetToken(vaultConfig.Token)

    secretValues, err := client.Logical().Read("secret/postgresql_creds")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    fmt.Printf("%+v\n", secretValues)
    fmt.Printf("%s - %s\n", secretValues.Data["username"], secretValues.Data["password"])
}
