// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package src

type CreateDraftInput struct {
	Title       string `json:"title"`
	Content     string `json:"content"`
	AuthorEmail string `json:"authorEmail"`
}

type DeletePostInput struct {
	ID string `json:"id"`
}

type PublishInput struct {
	ID string `json:"id"`
}