package domain

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type User struct {
	Id string `db:"id"`
	FirstName string `db:"first_name"`
	LastName string `db:"last_name"`
	Email string `db:"email"`
	Password string `db:"password"`
	Type string `db:"type"`
}

type UserRepository interface {
	Find(id string) (*User, *error)
}

type UserRepositoryDb struct {
	client *sqlx.DB
}

func (r UserRepositoryDb) Find(id string) (*User, *error) {
	return nil, nil
}

func NewUserRepositoryDb(dbClient *sqlx.DB) UserRepositoryDb {
	return UserRepositoryDb{dbClient}
}
