package src

import (
	"context"
	dbm "github.com/yarncraft/mini-graphql/src/db"
	"github.com/yarncraft/mini-graphql/src/models"
)

var db = dbm.GetConnection()

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateDraft(ctx context.Context, input *CreateDraftInput) (*models.Post, error) {
    var author models.User
	db.Where("email = ?", input.AuthorEmail).First(&author)
	post := models.Post{
		Title: input.Title,
		Content: input.Content,
		AuthorID: author.ID,
	}
	db.Create(&post)
	return &post, nil
}

func (r *mutationResolver) DeletePost(ctx context.Context, input *DeletePostInput) (*models.Post, error) {
	panic("not implemented")
}
func (r *mutationResolver) Publish(ctx context.Context, input *PublishInput) (*models.Post, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Feed(ctx context.Context) ([]*models.Post, error) {
	panic("not implemented")
}
func (r *queryResolver) Drafts(ctx context.Context) ([]*models.Post, error) {
	panic("not implemented")
}
func (r *queryResolver) Post(ctx context.Context, id string) (*models.Post, error) {
	panic("not implemented")
}
