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

type Student struct {
	ID          string    `db:"id"`
	StdId       string    `db:"std_id"`
	Name        string    `db:"name"`
	Email       string    `db:"email"`
	NameBn      string    `db:"name_bn"`
	FathersName string    `db:"fathers_name"`
	MothersName string    `db:"mothers_name"`
	DOB         time.Time `db:"dob"`

	Gender       string       `db:"gender"`
	BloodGroup   string       `db:"blood_group"`
	MobileNumber string       `db:"mobile_number"`
	Session      string       `db:"session"`
	ClassName    string       `db:"class_name"`
	ClassSection string       `db:"class_section"`
	Address      string       `db:"address"`
	Religion     string       `db:"religion"`
	IsActive     bool         `db:"is_active"`
	CreatedAt    time.Time    `db:"created_at"`
	UpdatedAt    sql.NullTime `db:"updated_at"`
	CreatedBy    string       `db:"created_by"`
	UpdatedBy    string       `db:"updated_by"`
}

type Class struct {
	SessionId int32  `db:"session_id"`
	Year      string `db:"year"`
	Class     string `db:"class"`
}

type AcademicSession struct {
	StudentId string `db:"student_id"`
	SessionId int32  `db:"session_id"`
}

type Query interface {
	InsertUser(u User) error
	InsertStudent(u Student) (string, error)
	InsertAcademicSession(u AcademicSession) error
	StudentList() ([]*Student, error)
	GetStudentListBySession(sessionId int32) ([]*Student, error)
	GetClassList() ([]*Class, error)
	GetUserByEmail(email string) (*User, error)
	GetUserByPhone(phone string) (*User, error)
	GetUserByID(userID string) (User, error)
}
