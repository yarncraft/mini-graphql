package models


type User struct {
	ID        int
	Email     string
	Name	  string
	Posts     []Post
}