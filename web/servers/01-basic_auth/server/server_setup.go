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

func getDBCreds(configFile string) string {
	vaultConfig, err := readVaultConfig(configFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	client, err := api.NewClient(&api.Config{
		Address: "http://vault", //vaultConfig.Addr,
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	client.SetToken(vaultConfig.Token)

	secretValues, err := client.Logical().Read(vaultConfig.Service)
    if err != nil {
		panic(err)
	}
    if secretValues == nil {
        panic(fmt.Sprintf("getDBCreds error:%s doesn't seem to exist", vaultConfig.Service))
    }

	var creds strings.Builder
    if val, ok := secretValues.Data["username"]; ok {
        creds.WriteString(fmt.Sprintf("user=%s ", val))
    } else {
        return "", errors.New("getDBCreds error: no username field on vault path")
    }
    if val, ok := secretValues.Data["password"]; ok {
        creds.WriteString(fmt.Sprintf("password=%s ", val))
    } else {
        return "", errors.New("getDBCreds error: no password field on vault path")
    }
	creds.WriteString("host=db dbname=db sslmode=disable")
	return creds.String(), nil
}
