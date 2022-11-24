package data

import (
	"math/rand"
	"strconv"
	"time"
)

type Review struct {
	ID          int64
	Title       string
	Description string
	Reviewer    User
}

var (
	Reviews = make(map[int64]Review)
)

func InitReviewers() {
	for i := 1; i <= 15; i++ {
		userID := randNumInt64()
		review := Review{
			ID:          int64(i),
			Title:       "title " + strconv.Itoa(i),
			Description: "desc " + strconv.Itoa(i),
			Reviewer: User{
				ID: userID,
			},
		}
		addUserReview(userID, 1, review)
		Reviews[int64(i)] = review
	}

	for i, review := range Reviews {
		if reviewer, ok := Reviewers[review.Reviewer.ID]; ok {
			review.Reviewer = reviewer
			Reviews[i] = review
		}

	}
}

func randNumInt64() int64 {
	var (
		min = 1
		max = 9
	)

	rand.Seed(time.Now().UnixNano())
	return int64(rand.Intn(max-min+1) + min)
}
