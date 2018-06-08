package main

import (
	"database/sql"
	"log"
	"net/http"
    "fmt"

	_ "github.com/lib/pq"
)

//var db *sql.DB

func initDB() {
	dataSource, err := getDBCreds("/vault/configs/vault.json")
    if err != nil {
        panic(err)
    }

	_, err = sql.Open("postgres", dataSource)
	if err != nil {
	    panic(err)
    }
    fmt.Println(dataSource)
}

func main() {
	initDB()

	//http.HandleFunc("/singin", Signin)
	//http.HandleFunc("/signup", Signup)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
