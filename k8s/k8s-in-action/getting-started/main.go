package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	// livenessProbes counts the number of times a request has been served by
	// healthHandler.
	livenessProbes = 0
	// probesBeforeFail is the number of times a liveness probe must be
	// conducted before /health returns a bad response code.
	probesBeforeFail = 10
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	msg := fmt.Sprintf("%s - %s", r.RemoteAddr, hostname)

	log.Println(msg)
	fmt.Fprintf(w, msg)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	livenessProbes += 1

	if livenessProbes > probesBeforeFail {
		log.Println("Failing health probes")
		http.Error(w, "failing for a bit...", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("ok"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/health", healthHandler)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
