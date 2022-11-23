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

var Reviewers = make(map[int64]Review)

func InitReviewers() {
	for i := 1; i <= 15; i++ {
		Reviewers[int64(i)] = Review{
			ID:          int64(i),
			Title:       "title " + strconv.Itoa(i),
			Description: "desc " + strconv.Itoa(i),
			Reviewer: User{
				ID: randNumInt64(),
			},
		}
	}
}

func randNumInt64() int64 {
	var (
		min = 1
		max = 20
	)

	rand.Seed(time.Now().UnixNano())
	return int64(rand.Intn(max-min+1) + min)
}
