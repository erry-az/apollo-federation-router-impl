package data

import "strconv"

type User struct {
	ID    int64
	Name  string
	Email string
}

var Users = make(map[int64]User)

func InitSampleUser() {
	for i := 1; i < 10; i++ {
		Users[int64(i)] = User{
			ID:    int64(i),
			Name:  "Test " + strconv.Itoa(i),
			Email: "test" + strconv.Itoa(i) + "abc.jpg",
		}
	}
}
