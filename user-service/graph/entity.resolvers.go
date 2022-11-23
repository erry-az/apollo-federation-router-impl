package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/erry-az/go-gql-federation/user-service/data"
	"github.com/erry-az/go-gql-federation/user-service/graph/generated"
	"github.com/erry-az/go-gql-federation/user-service/graph/model"
)

// FindUserByID is the resolver for the findUserByID field.
func (r *entityResolver) FindUserByID(ctx context.Context, id int64) (*model.User, error) {
	user, ok := data.Users[id]
	if !ok {
		return nil, nil
	}

	return &model.User{
		ID:    user.ID,
		Name:  &user.Name,
		Email: &user.Email,
	}, nil
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
