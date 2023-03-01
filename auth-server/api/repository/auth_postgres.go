package repository

import (
	"auth-server/model"
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	userTable = "users"
)

type AuthDB struct {
	db *sqlx.DB
}

func NewAuthBD(db *sqlx.DB) *AuthDB {
	return &AuthDB{db: db}
}

func (r *AuthDB) CreateUser(user model.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (login, password_hash) values ($1, $2) RETURNING id", userTable)

	row := r.db.QueryRow(query, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthDB) GetUser(username string, password string) (model.User, error) {
	var usr model.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE login=$1 AND password_hash=$2 ", userTable)

	err := r.db.Get(&usr, query, username, password)

	return usr, err
}
