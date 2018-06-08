package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

//var db *sql.DB

func initDB() {
	dataSource := getDBCreds("/vault/configs/vault.json")

	_, err := sql.Open("postgres", dataSource)
	if err != nil {
		panic(err)
	}
}

func main() {
	initDB()

	//http.HandleFunc("/singin", Signin)
	//http.HandleFunc("/signup", Signup)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
