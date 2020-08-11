package main

import (
	"time"
)

type Blog struct {
	ID          int64     `json:"id"`
	Author      string    `json:"author"`
	Title       string    `json:"title"`
	Slug        string    `json:"slug"` // URL friendly representation of the link
	Body        string    `json:"body"`
	CreatedDate time.Time `json:"created_date"`
	Published   bool      `json:"published"`
}
