// The servver program issues Google search requests and demonstrates the use
// of go.net Context API. It serves on port 8080.
//
// The /search endpoint accepts these query params:
//  q= the Google search query
//  timeout= a timeout for the requests, in time.Duration format
//
// For example, http://localhost:8080/search?q=golang&timeout=1s serves the
// first few Google search results for "golang: or a "deadline exceeded" error
// if the timeout expires.
package main

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/alejandrox1/cool_go/web/context/go-blog/google"
	"github.com/alejandrox1/cool_go/web/context/go-blog/userip"
)

func main() {
	http.HandleFunc("/search", handleSearch)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// handleSearch handles URLS like /search?q=golang&timeout=1s by forwarding the
// query to google.Search. If the query param includes timeout, the search is
// canceled after the duration elapses.
func handleSearch(w http.ResponseWriter, req *http.Request) {
	// ctx is the Context for this handler. Calling cancel closes the ctx.Done
	// channel, which is the cancellation signal for requests started by this
	// handler.
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)
	timeout, err := time.ParseDuration(req.FormValue("timeout"))
	if err == nil {
		ctx, cancel = context.WithTimeout(context.Background(), timeout)
	} else {
		ctx, cancel = context.WithCancel(context.Background())
	}
	defer cancel()

	// Check the search query.
	query := req.FormValue("q")
	if query == "" {
		http.Error(w, "no query", http.StatusBadRequest)
		return
	}

	// Store the user IP in ct for use by other packages.
	userIP, err := userip.FromRequest(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ctx = userip.NewContext(ctx, userIP)

	// Run the Google search and print the results.
	start := time.Now()
	results, err := google.Search(ctx, query)
	elapsed := time.Since(start)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := resultsTemplate.Execute(w, struct {
		Results          google.Results
		Timeout, Elapsed time.Duration
	}{
		Results: results,
		Timeout: timeout,
		Elapsed: elapsed,
	}); err != nil {
		log.Println(err)
		return
	}
}

var resultsTemplate = template.Must(template.New("results").Parse(`
<html>
<head/>
<body>
    <ol>
        {{range .Results}}
            <li>{{.Title}} - <a href="{{.URL}}">{{.URL}}</a></li>
        {{end}}
    </ol>
    <p>{{len .Results}} results in {{.Elapsed}}; timeout {{.Timeout}}</p>
</body>
</html>
`))
