package models

import (
	"time"
)

type Article struct {
	ID        int       `json:"id" gorm:"primary_key;auto_increment"`
	Author    string    `json:"author"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateArticle struct {
	Author string `json:"author"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type GetArticle struct {
	Query  string `form:"query"`
	Author string `form:"author"`
}
