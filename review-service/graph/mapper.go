package graph

import (
	"github.com/erry-az/go-gql-federation/review-service/data"
	"github.com/erry-az/go-gql-federation/review-service/graph/model"
)

func newUserModel(user data.User) *model.User {
	return &model.User{
		ID:          user.ID,
		TotalReview: &user.TotalReview,
		Reviews: func() []*model.Review {
			reviews := make([]*model.Review, 0, len(data.Reviewers[user.ID].Reviews))
			for _, review := range data.Reviewers[user.ID].Reviews {
				reviews = append(reviews, &model.Review{
					ID:          review.ID,
					Title:       &review.Title,
					Description: &review.Description,
				})

				// if query call this it will be loop
				//reviews = append(reviews, newReviewModel(review))
			}

			return reviews
		}(),
	}
}

func newReviewModel(review data.Review) *model.Review {
	return &model.Review{
		ID:          review.ID,
		Title:       &review.Title,
		Description: &review.Description,
		Reviewer:    newUserModel(review.Reviewer),
	}
}
