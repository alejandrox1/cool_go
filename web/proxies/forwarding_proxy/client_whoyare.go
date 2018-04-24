package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type whoiam struct {
	Addr string
}

func main() {
	url := "http://localhost:8080"
	log.Printf("Tagert %s.", url)

	resp, err := http.Get(url + "/whoyare")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("You are " + string(body))
}
