package dto

import "github.com/psocietyyy/go-gin-kafka/service/auth_service/internal/model"

type User struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func ToUserResponse(u *model.User) User {
	return User{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}
}

func ToUserResponses(u []model.User) []User {
	var users []User
	for _, user := range u {
		users = append(users, ToUserResponse(&user))
	}
	return users
}