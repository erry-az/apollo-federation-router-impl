package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/erry-az/go-gql-federation/user-service/data"
	"github.com/erry-az/go-gql-federation/user-service/graph/generated"
	"github.com/erry-az/go-gql-federation/user-service/graph/model"
)

// UserGetByID is the resolver for the UserGetByID field.
func (r *queryResolver) UserGetByID(ctx context.Context, id int64) (*model.User, error) {
	user, ok := data.Users[id]
	if !ok {
		return nil, errors.New("user not found")
	}

	return &model.User{
		ID:    user.ID,
		Name:  &user.Name,
		Email: &user.Email,
	}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
