package domain

import "time"

type User struct {
	id        int
	name      string
	email     string
	age       int
	createdAt time.Time
	updatedAt time.Time
}

func NewUser(name, email string, age int) *User {
	return &User{
		id:    0,
		name:  name,
		email: email,
		age:   age,
	}
}

func (u *User) Id() int {
	return u.id
}

func (u *User) Name() string {
	return u.name
}

func (u *User) Email() string {
	return u.email
}

func (u *User) Age() int {
	return u.age
}

func (u *User) CreatedAt() time.Time {
	return u.createdAt
}

func (u *User) UpdatedAt() time.Time {
	return u.updatedAt
}

func (u *User) SetId(id int) {
	u.id = id
}
func (u *User) SetName(name string) {
	u.name = name
}
func (u *User) SetEmail(email string) {
	u.email = email
}
func (u *User) SetAge(age int) {
	u.age = age
}

func (u *User) SetCreatedAt(createdAt time.Time) {
	u.createdAt = createdAt
}
func (u *User) SetUpdatedAt(updatedAt time.Time) {
	u.updatedAt = updatedAt
}
