package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/google/go-github/github"
)

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error webhookHandler - could not read the request body: %s\n", err)
	}
	defer r.Body.Close()

	event, err := github.ParseWebHook(github.WebHookType(r), payload)
	if err != nil {
		log.Printf("Error webhookHandler - error parsing webhook: %s\n", err)
	}

	switch e := event.(type) {
	case *github.PushEvent:
		var buffer bytes.Buffer
		var msg string

		if e.Ref != nil {
			msg = fmt.Sprintf("\t\t%s %s:\n", *e.Repo.Name, *e.Ref)
			buffer.WriteString(msg)

			for _, commit := range e.Commits {
				msg = fmt.Sprintf("%s\n%s\n", *commit.ID, *commit.Message)
				buffer.WriteString(msg)

				for _, a := range commit.Added {
					msg = fmt.Sprintf("\tA: %s\n", a)
					buffer.WriteString(msg)
				}
				for _, r := range commit.Removed {
					msg = fmt.Sprintf("\tR: %s\n", r)
					buffer.WriteString(msg)
				}
				for _, m := range commit.Modified {
					msg = fmt.Sprintf("\tM: %s\n", m)
					buffer.WriteString(msg)
				}

				msg = fmt.Sprintf("By: %s\n\n", *commit.Author.Name)
				buffer.WriteString(msg)
				fmt.Println(buffer.String())
			}
		}

	case *github.IssuesEvent:
		var buffer bytes.Buffer
		var msg string

		msg = fmt.Sprintf("%s %s issue %d on %s\n", *e.Sender.Login, *e.Action, *e.Issue.Number, *e.Repo.Name)
		buffer.WriteString(msg)
		msg = fmt.Sprintf("%s\n%s\n\n", *e.Issue.Title, *e.Issue.Body)
		buffer.WriteString(msg)
		fmt.Println(buffer.String())

	case *github.IssueCommentEvent:
		var buffer bytes.Buffer
		var msg string

		msg = fmt.Sprintf("%s %s a comment on Issue %d (%s) on %s\n", *e.Sender.Login, *e.Action, *e.Issue.Number, *e.Issue.Title, *e.Repo.Name)
		buffer.WriteString(msg)
		msg = fmt.Sprintf("%s\n", *e.Comment.Body)
		buffer.WriteString(msg)
		fmt.Println(buffer.String())
	}
}

func main() {
	http.HandleFunc("/webhook", webhookHandler)

	port := 8082
	log.Printf("Server starting on port: %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
