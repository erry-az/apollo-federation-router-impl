package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/erry-az/go-gql-federation/review-service/data"
	"github.com/erry-az/go-gql-federation/review-service/graph/generated"
	"github.com/erry-az/go-gql-federation/review-service/graph/model"
)

// ReviewList is the resolver for the ReviewList field.
func (r *queryResolver) ReviewList(ctx context.Context) (*model.ReviewListResponse, error) {
	reviewResponse := make([]*model.Review, 0, len(data.Reviews))

	for _, review := range data.Reviews {
		reviewResponse = append(reviewResponse, newReviewModel(review))
	}

	return &model.ReviewListResponse{Reviews: reviewResponse}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
