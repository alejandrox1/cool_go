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

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("ok"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/health", healthHandler)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
