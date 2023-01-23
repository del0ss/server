package store

import (
	"fmt"
	"smth/internal/model"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	fmt.Println(u)
	err := r.store.db.QueryRow(
		"INSERT INTO users (login, email, enc_password) VALUES ($1, $2, $3) RETURNING id",
		u.Login,
		u.Email,
		u.EncPassword,
	).Scan(&u.ID)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *UserRepository) FindByLogin(login string) (*model.User, error) {
	return nil, nil
}
