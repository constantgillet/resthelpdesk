package models

import "time"

type Article struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Category  int       `json:"category"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Articles []Article
