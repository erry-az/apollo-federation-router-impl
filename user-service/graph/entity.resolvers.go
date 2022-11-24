package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"
	"time"

	"github.com/erry-az/go-gql-federation/user-service/data"
	"github.com/erry-az/go-gql-federation/user-service/graph/generated"
	"github.com/erry-az/go-gql-federation/user-service/graph/model"
)

// FindUserByID is the resolver for the findUserByID field.
func (r *entityResolver) FindUserByID(ctx context.Context, id int64) (*model.User, error) {
	counterUserByID++

	defer func() {
		log.Println("user ent: ", counterUserByID, ": ", time.Now().UnixMilli())
	}()
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

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
var counterUserByID int
