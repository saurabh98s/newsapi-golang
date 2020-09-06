package models

import "time"

// News implements the data recieved from the API
type News struct {
	Status       string `json:"status"`
	TotalResults int    `json:"totalResults"`
	Articles     []Articles `json:"articles"`
}

// Articles stores the data of type Articles
type  Articles   struct {
	Source struct {
		ID   interface{} `json:"id"`
		Name string      `json:"name"`
	} `json:"source"`
	Author      string    `json:"author"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	URLToImage  string    `json:"urlToImage"`
	PublishedAt time.Time `json:"publishedAt"`
	Content     string    `json:"content"`
}
