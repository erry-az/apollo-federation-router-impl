package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"
	"time"

	"github.com/erry-az/go-gql-federation/review-service/data"
	"github.com/erry-az/go-gql-federation/review-service/graph/generated"
	"github.com/erry-az/go-gql-federation/review-service/graph/model"
)

// FindReviewByID is the resolver for the findReviewByID field.
func (r *entityResolver) FindReviewByID(ctx context.Context, id int64) (*model.Review, error) {
	counterReviewByID++

	defer func() {
		log.Println("review ent: ", counterReviewByID, ": ", time.Now().UnixMilli())
	}()

	review, ok := data.Reviews[id]
	if !ok {
		return nil, nil
	}

	return &model.Review{
		ID:          review.ID,
		Title:       &review.Title,
		Description: &review.Description,
		Reviewer: &model.User{
			ID: review.Reviewer.ID,
			TotalReview: func() *int {
				totalReview := data.Reviewers[review.Reviewer.ID].TotalReview

				return &totalReview
			}(),
			Reviews: func() []*model.Review {
				reviews := make([]*model.Review, 0, len(data.Reviewers[review.Reviewer.ID].Reviews))
				for _, review := range data.Reviewers[review.Reviewer.ID].Reviews {
					reviews = append(reviews, &model.Review{
						ID:          review.ID,
						Title:       &review.Title,
						Description: &review.Description,
					})
				}

				return reviews
			}(),
		},
	}, nil
}

// FindUserByID is the resolver for the findUserByID field.
func (r *entityResolver) FindUserByID(ctx context.Context, id int64) (*model.User, error) {
	counterFindByUserID++

	defer func() {
		log.Println("review user ent: ", counterFindByUserID, ": ", time.Now().UnixMilli())
	}()

	userReview, ok := data.Reviewers[id]
	if !ok {
		return nil, nil
	}

	return &model.User{
		ID: userReview.ID,
		Reviews: func() []*model.Review {
			reviews := make([]*model.Review, 0, len(data.Reviewers[id].Reviews))
			for _, review := range data.Reviewers[id].Reviews {
				reviews = append(reviews, &model.Review{
					ID:          review.ID,
					Title:       &review.Title,
					Description: &review.Description,
				})
			}

			return reviews
		}(),
		TotalReview: &userReview.TotalReview,
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
var (
	counterReviewByID   int
	counterFindByUserID int
)
