package primary

import "rest-api/src/core/domain"

type UserServicePort interface {
	Create(user domain.User) (domain.User, error)
	GetByID(id int) (domain.User, error)
	GetAll() ([]domain.User, error)
	Update(user domain.User) (domain.User, error)
	Delete(id int) error
}
