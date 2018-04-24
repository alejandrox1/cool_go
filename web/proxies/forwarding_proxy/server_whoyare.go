package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func whoyareHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, err := json.Marshal(map[string]string{
		"addr": r.RemoteAddr,
	})
	if err != nil {
		return
	}

	fmt.Fprintf(w, "%s\n", body)
}

func main() {
	http.HandleFunc("/whoyare", whoyareHandler)

	http.ListenAndServe(":8080", nil)
}
