package service

import (
	"CleanArchitecture/repo"
)

type UserService struct {
	repo repo.UserRepoInterface
}

func NewUserService(repo repo.UserRepoInterface) *UserService {
	return &UserService{repo: repo}
}

func (service *UserService) CreateUser(name string) repo.User {
	return service.repo.AddUser(name)
}

func (service *UserService) ListUsers() []repo.User {
	return service.repo.GetAllUsers()
}
