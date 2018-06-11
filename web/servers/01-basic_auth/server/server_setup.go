package main

import (
	"encoding/json"
    "errors"
	"fmt"
	"os"
	"strings"

	"github.com/hashicorp/vault/api"
)

type Vault struct {
	Addr    string `json:"vault_address"`
	Service string `json:"vault_service"`
	Token   string `json:"vault_token"`
}

func readVaultConfig(filename string) (vaultConfig Vault, err error) {
	contents, err := os.Open(filename)
	if err != nil {
		fmt.Printf("readVaultConfig error while opening file: %s\n", err.Error())
		return
	}

	err = json.NewDecoder(contents).Decode(&vaultConfig)
	if err != nil {
		fmt.Printf("readVaultConfig error while decoding json: %s\n", err.Error())
		return
	}

	return
}

func getDBCreds(configFile string) (string, error) {
	vaultConfig, err := readVaultConfig(configFile)
	if err != nil {
        fmt.Printf("getDBCreds error while getting configuration information from %s: %s\n", configFile, err)
		panic(err)
	}

	client, err := api.NewClient(&api.Config{
		Address: vaultConfig.Addr,
	})
	if err != nil {
        fmt.Printf("getDBCreds error while instantiating vault client: %s\n", err)
		os.Exit(1)
	}
	client.SetToken(vaultConfig.Token)

	secretValues, err := client.Logical().Read(vaultConfig.Service)
    if err != nil {
        fmt.Printf("getDBCreds error while trying to read from the specified secret: %s\n", err)
		panic(err)
	}
    if secretValues == nil {
        panic(fmt.Sprintf("getDBCreds error while reading from secret: %s doesn't seem to exist", vaultConfig.Service))
    }

	var creds strings.Builder
    if val, ok := secretValues.Data["username"]; ok {
        creds.WriteString(fmt.Sprintf("user=%s ", val))
    } else {
        return "", errors.New("getDBCreds error while reading from secret: no username field on vault path")
    }
    if val, ok := secretValues.Data["password"]; ok {
        creds.WriteString(fmt.Sprintf("password=%s ", val))
    } else {
        return "", errors.New("getDBCreds error while reading from secret: no password field on vault path")
    }
	creds.WriteString("host=db dbname=db sslmode=disable")
	return creds.String(), nil
}
