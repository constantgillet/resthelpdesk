package models

import "time"

type Article struct {
	Id         int       `json:"id"`
	Title      string    `json:"title"`
	CategoryId int       `json:"category_id"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Articles []Article
