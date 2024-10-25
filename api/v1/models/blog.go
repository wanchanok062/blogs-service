package models

import (
	"time"
)

type BlogPost struct {
	PostID    string    `json:"id" bson:"id"`
	Title     string    `json:"title" validate:"required"`
	Content   string    `json:"content"  validate:"required"`
	AuthorID  string    `json:"authorId"  validate:"required"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type BlogUpdate struct {
	PostID    string    `json:"id" bson:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	AuthorID  string    `json:"authorId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type FilterBlogsOptions struct {
	AuthorID string
	Search   string
	Sort     string
}
