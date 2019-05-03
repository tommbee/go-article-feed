package model

import "time"

// Article entity
type Article struct {
	Title     string    `json:"title" bson:"title"`
	URL       string    `json:"url" bson:"url"`
	Published time.Time `json:"published" bson:"published"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}
