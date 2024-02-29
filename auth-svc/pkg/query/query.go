package query

import (
	"database/sql"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
)

type QueryInit struct {
	Db *sqlx.DB
}

var NotFound = errors.New("not found")

type User struct {
	ID       string       `db:"id"`
	Email    string       `db:"email"`
	Password string       `db:"password"`
	Phone    string       `db:"phone"`
	IsActive string       `db:"is_active"`
	Created  time.Time    `db:"created"`
	Updated  sql.NullTime `db:"updated"`
}

type Query interface {
	InsertUser(u User) error
	GetUserByEmail(email string) (*User, error)
	GetUserByPhone(phone string) (*User, error)
	GetUserByID(userID string) (User, error)
}
