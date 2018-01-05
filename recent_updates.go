package main

import "time"

type Activity struct {
	Revisions []struct {
		URL     string `json:"url"`
		Message string `json:"message"`
		Author  struct {
			Email string `json:"email"`
			Name  string `json:"name"`
		} `json:"author"`
		ID        string    `json:"id"`
		Timestamp time.Time `json:"timestamp"`
	} `json:"revisions"`
	Before     string `json:"before"`
	After      string `json:"after"`
	Repository struct {
		URL         string `json:"url"`
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"repository"`
	Ref string `json:"ref"`
}
