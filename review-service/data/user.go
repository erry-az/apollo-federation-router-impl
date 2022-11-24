package data

type User struct {
	ID          int64
	TotalReview int
	Reviews     []Review
}

var Reviewers = make(map[int64]User)

func addUserReview(id int64, totalReview int, review Review) {
	reviewExist, ok := Reviewers[id]
	if ok {
		reviewExist.TotalReview += totalReview
		reviewExist.Reviews = append(reviewExist.Reviews, review)
		Reviewers[id] = reviewExist

		return
	}

	Reviewers[id] = User{
		ID:          id,
		TotalReview: totalReview,
		Reviews:     []Review{review},
	}
}
