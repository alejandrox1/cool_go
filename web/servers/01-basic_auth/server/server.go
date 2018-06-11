package main

import (
	"database/sql"
	"log"
	"net/http"
    "fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

func initDB() {
	dataSource, err := getDBCreds("/vault/configs/vault.json")
    if err != nil {
        panic(err)
    }

	db, err = sql.Open("postgres", dataSource)
	if err != nil {
	    panic(err)
    }
    fmt.Println(dataSource)
}

func main() {
	initDB()

    http.HandleFunc("/signup", Signup)
	http.HandleFunc("/signin", Signin)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
