package postgres

import (
	"context"
	"log"
	"rest-api/src/core/domain"
	"rest-api/src/core/interfaces/secondary"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

var _ secondary.UserRepositoryPort = (*UserRepository)(nil)

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Insert(user domain.User) (domain.User, error) {
	query := `INSERT INTO users (name, email, age) VALUES ($1, $2, $3) RETURNING id, created_at, updated_at`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	row := r.db.QueryRow(ctx, query, user.Name(), user.Email(), user.Age())

	var id int
	var createdAt, updatedAt time.Time

	err := row.Scan(&id, &createdAt, &updatedAt)
	if err != nil {
		log.Fatal("Error scanning from row: ", err)
		return domain.User{}, err
	}

	user.SetId(id)
	user.SetCreatedAt(createdAt)
	user.SetUpdatedAt(updatedAt)

	return user, nil
}

func (r *UserRepository) FetchByID(id int) (domain.User, error) {
	panic("unimplemented")
}
func (r *UserRepository) FetchAll() ([]domain.User, error) {
	panic("unimplemented")
}
func (r *UserRepository) Update(user domain.User) (domain.User, error) {
	panic("unimplemented")
}
func (r *UserRepository) Delete(id int) error {
	panic("unimplemented")
}
