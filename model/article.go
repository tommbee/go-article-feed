package model

import "time"

// Article entity
type Article struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	URL       string    `json:"url"`
	Published time.Time `json:"published"`
	CreatedAt time.Time `json:"created_at"`
}
