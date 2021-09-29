package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/timc4662/go_graphql/graph/generated"
	"github.com/timc4662/go_graphql/graph/model"
	"github.com/timc4662/go_graphql/internal/auth"
	"github.com/timc4662/go_graphql/internal/links"
	"github.com/timc4662/go_graphql/internal/users"
	"github.com/timc4662/go_graphql/pkg/jwt"
)

func (r *mutationResolver) CreateLink(ctx context.Context, input model.NewLink) (*model.Link, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return nil, errors.New("access denied")
	}
	link := links.Link{Address: input.Address, Title: input.Title, User: user}
	linkID := link.Save()
	gUser := model.User{
		ID:   user.ID,
		Name: user.Username,
	}
	return &model.Link{ID: strconv.FormatInt(linkID, 10), Title: link.Title, Address: link.Address, User: &gUser}, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	user := users.User{Username: input.Username, Password: input.Password}
	user.Create()
	token, err := jwt.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	user := users.User{Username: input.Username, Password: input.Password}
	valid := user.Authenticate()
	if !valid {
		return "", &users.WrongUsernameOfPasswordError{}
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
	var resultLinks []*model.Link
	dbLinks := links.GetAll()
	for _, link := range dbLinks {
		gUser := model.User{
			ID:   link.User.ID,
			Name: link.User.Username,
		}
		resultLinks = append(resultLinks,
			&model.Link{ID: link.ID, Title: link.Title, Address: link.Address, User: &gUser})
	}
	return resultLinks, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
