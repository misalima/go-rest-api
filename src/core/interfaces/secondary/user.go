package secondary

import "rest-api/src/core/domain"

type UserRepositoryPort interface {
	Insert(user domain.User) (domain.User, error)
	FetchByID(id int) (domain.User, error)
	FetchAll() ([]domain.User, error)
	Update(user domain.User) (domain.User, error)
	Delete(id int) error
}
