package services

import (
	"log"
	"rest-api/src/core/domain"
	"rest-api/src/core/interfaces/primary"
	"rest-api/src/core/interfaces/secondary"
)

type UserService struct {
	repo secondary.UserRepositoryPort
}

var _ primary.UserServicePort = (*UserService)(nil)

func (u *UserService) Create(user domain.User) (domain.User, error) {
	user, err := u.repo.Insert(user)
	if err != nil {
		log.Fatal("Error inserting user: ", err)
		return domain.User{}, err
	}
	return user, nil
}

func (u UserService) GetByID(id int) (domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserService) GetAll() ([]domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserService) Update(user domain.User) (domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserService) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}

func NewUserService(repo secondary.UserRepositoryPort) *UserService {
	return &UserService{repo}
}
