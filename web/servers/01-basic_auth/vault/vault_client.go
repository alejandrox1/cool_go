package main

import (
    "encoding/json"
    "fmt"
    "os"

//    "github.com/hashicorp/vault/api"
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
    vaultConfig, err := readVaultConfig("./vault.json")
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    encoder := json.NewEncoder(os.Stdout)
    encoder.SetIndent("", "\t")
    err = encoder.Encode(&vaultConfig)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

}
