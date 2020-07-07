package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/fk/gqlplayground/auth"
	"github.com/fk/gqlplayground/db"
	"github.com/fk/gqlplayground/graph/generated"
	"github.com/fk/gqlplayground/graph/model"
	"github.com/fk/gqlplayground/pkg/jwt"
	"github.com/pkg/errors"
)

func (r *mutationResolver) CreateLink(ctx context.Context, input model.NewLink) (*model.Link, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return &model.Link{}, fmt.Errorf("access denied")
	}
	var link model.Link
	link.Address = input.Address
	link.Title = input.Title
	link.User = user
	id, err := r.Db.CreateLink(&link)
	if err != nil {
		return &model.Link{}, err
	}
	newLink, err := r.Db.GetLinkById(id.Hex())
	return newLink, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	var user model.NewUser
	user.Username = input.Username
	user.Password = input.Password
	//check if already exists
	err := r.Db.CreateUser(&user)
	token, err := jwt.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	var user model.Login
	user.Username = input.Username
	user.Password = input.Password
	pwdFromDb, err := r.Db.GetPasswordByUsername(user.Username)
	if err != nil {
		return "", errors.Errorf("wrong username or password")
	}
	ok := db.CheckPasswordHash(user.Password, pwdFromDb)
	if !ok {
		return "", errors.Errorf("wrong username or password")
	}
	token, err := jwt.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	username, err := jwt.ParseToken(input.Token)
	if err != nil {
		return "", fmt.Errorf("access denied")
	}
	token, err := jwt.GenerateToken(username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *queryResolver) Links(ctx context.Context) ([]*model.Link, error) {
	result, err := r.Db.GetAllLinks()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
