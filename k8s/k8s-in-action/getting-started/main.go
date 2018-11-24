package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	msg := fmt.Sprintf("%s - %s", r.RemoteAddr, hostname)

	log.Println(msg)
	fmt.Fprintf(w, msg)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
