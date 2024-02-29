package query

import (
	"database/sql"
)

func (q *QueryInit) InsertUser(u User) error {
	stmt, err := q.Db.PrepareNamed("INSERT INTO users (email, password, phone) VALUES (:email, :password, :phone)")
	if err != nil {
		return err

	}
	_, err = stmt.Exec(&u)

	if err != nil {
		return err
	}
	return nil
}

func (q *QueryInit) GetUserByEmail(email string) (*User, error) {
	var user User
	qc := "SELECT * FROM users WHERE email=$1"
	if err := q.Db.Get(&user, qc, email); err != nil {
		if err == sql.ErrNoRows {
			return nil, NotFound
		}
		return nil, err
	}
	return &user, nil
}
func (q *QueryInit) GetUserByPhone(phone string) (*User, error) {
	var user User
	qc := "SELECT * FROM users WHERE phone=$1"
	if err := q.Db.Get(&user, qc, phone); err != nil {
		if err == sql.ErrNoRows {
			return nil, NotFound
		}
		return nil, err
	}
	return &user, nil
}

func (q *QueryInit) GetUserByID(userID string) (User, error) {

	var user User
	const qc = `SELECT * FROM users WHERE id=$1`
	if err := q.Db.Get(&user, qc, userID); err != nil {
		if err == sql.ErrNoRows {
			return user, NotFound
		}

		return user, err
	}
	return user, nil
}
