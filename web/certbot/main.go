package main

import (
    "fmt"
	"log"
	"net/http"
)


func indexHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "hello world")
}


func main() {
    http.HandleFunc("/", indexHandler)
    http.Handle("/.well-known/acme-challenge/",
        http.StripPrefix("/.well-known/acme-challenge/", http.FileServer(http.Dir("."))))

	log.Fatal(http.ListenAndServe(":80", nil))
}
