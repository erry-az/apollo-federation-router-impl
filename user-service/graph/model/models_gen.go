// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type User struct {
	ID    int64   `json:"id"`
	Name  *string `json:"name"`
	Email *string `json:"email"`
}

func (User) IsEntity() {}
