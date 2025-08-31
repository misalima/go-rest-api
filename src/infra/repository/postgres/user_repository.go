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
	query := `SELECT id, name, email, age, created_at, updated_at FROM users WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	row := r.db.QueryRow(ctx, query, id)

	var user domain.User
	var createdAt, updatedAt time.Time
	var name, email string
	var age int

	err := row.Scan(&id, &name, &email, &age, &createdAt, &updatedAt)
	if err != nil {
		log.Fatal("Error scanning from row: ", err)
		return domain.User{}, err
	}

	user.SetId(id)
	user.SetName(name)
	user.SetEmail(email)
	user.SetAge(age)
	user.SetCreatedAt(createdAt)
	user.SetUpdatedAt(updatedAt)
	return user, nil
}
func (r *UserRepository) FetchAll() ([]domain.User, error) {
	query := `SELECT id, name, email, age, created_at, updated_at FROM users`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		log.Fatal("Error querying rows: ", err)
		return nil, err
	}
	defer rows.Close()

	var users []domain.User

	for rows.Next() {
		var user domain.User
		var id int
		var name, email string
		var age int
		var createdAt, updatedAt time.Time

		err := rows.Scan(&id, &name, &email, &age, &createdAt, &updatedAt)
		if err != nil {
			log.Fatal("Error scanning from row: ", err)
			return nil, err
		}

		user.SetId(id)
		user.SetName(name)
		user.SetEmail(email)
		user.SetAge(age)
		user.SetCreatedAt(createdAt)
		user.SetUpdatedAt(updatedAt)

		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		log.Fatal("Error iterating over rows: ", err)
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) Update(user domain.User) (domain.User, error) {
	query := `UPDATE users SET name = $1, email = $2, age = $3, updated_at = NOW() WHERE id = $4 RETURNING updated_at, created_at`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	row := r.db.QueryRow(ctx, query, user.Name(), user.Email(), user.Age(), user.Id())

	var updatedAt time.Time
	var createdAt time.Time

	err := row.Scan(&updatedAt, &createdAt)
	if err != nil {
		log.Fatal("Error scanning from row: ", err)
		return domain.User{}, err
	}

	user.SetUpdatedAt(updatedAt)
	user.SetCreatedAt(createdAt)

	return user, nil
}
func (r *UserRepository) Delete(id int) error {
	query := `DELETE FROM users WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		log.Fatal("Error deleting user: ", err)
		return err
	}

	return nil
}
