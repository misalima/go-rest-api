package services

import (
	"rest-api/src/core/domain"
	"rest-api/src/core/interfaces/primary"
	"rest-api/src/core/interfaces/secondary"
)

type UserService struct {
	repo secondary.UserRepositoryPort
}

var _ primary.UserServicePort = (*UserService)(nil)

func (u *UserService) Create(user domain.User) (domain.User, error) {
	return u.repo.Insert(user)
}

func (u UserService) GetByID(id int) (domain.User, error) {
	return u.repo.FetchByID(id)
}

func (u UserService) GetAll() ([]domain.User, error) {
	return u.repo.FetchAll()
}

func (u UserService) Update(user domain.User) (domain.User, error) {
	return u.repo.Update(user)
}

func (u UserService) Delete(id int) error {
	return u.repo.Delete(id)
}

func NewUserService(repo secondary.UserRepositoryPort) *UserService {
	return &UserService{repo}
}
