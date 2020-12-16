package service

import "project/internal/data"

type UserService struct {

}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) ListUsers() []data.User {
	user := new(data.User)
	return user.FindAll()
}