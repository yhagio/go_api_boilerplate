package controllers

type User struct {
	us user.UserService
}

func NewUserController(us user.UserService) *User {
	return &User{
		us: us,
	}
}
