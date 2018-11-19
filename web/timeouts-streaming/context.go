package main

import (
	"context"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.TODO())
	timer := time.AfterFunc(5*time.Second, func() {
		cancel()
	})

	// Serve 256 bytes every second.
	url := "http://httpbin.org/range/2048?duration=8&chunk_size=256"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Error making requeust: %v\n", err)
	}
	req = req.WithContext(ctx)

	log.Println("Sending request...")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Error doing request: %v\n", err)
	}
	defer resp.Body.Close()

	log.Println("Reading body...")
	for i := 0; ; i++ {
		log.Println(i)

		if !timer.Stop() {
			<-timer.C
		}
		timer.Reset(2 * time.Second)

		_, err := io.CopyN(ioutil.Discard, resp.Body, 256)
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("Error reading: %v\n", err)
		}
	}
}
