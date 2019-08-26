package models

import (
	"time"
)

type Post struct {
	ID        		int
	Title   		string
	Content   		string
	IsPublished     bool
	AuthorID     	int
	CreatedAt 		time.Time
	UpdatedAt 		time.Time
}