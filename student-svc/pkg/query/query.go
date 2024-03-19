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

	VillageOrRoad string `db:"village_or_road"`
	Union         string `db:"union"`
	Upzilla       string `db:"upzilla"`
	District      string `db:"district"`

	SessionName string `db:"session_name"`
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

// District struct represents the District table.
type District struct {
	Id        int32        `db:"id"`
	Name      string       `db:"name"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
	CreatedBy string       `db:"created_by"`
	UpdatedBy string       `db:"updated_by"`
}

// Upazilla struct represents the Upazilla table.
type Upazilla struct {
	Id         int32        `db:"id"`
	Name       string       `db:"name"`
	DistrictId int32        `db:"district_id"`
	CreatedAt  time.Time    `db:"created_at"`
	UpdatedAt  sql.NullTime `db:"updated_at"`
	CreatedBy  string       `db:"created_by"`
	UpdatedBy  string       `db:"updated_by"`
}

// Union struct represents the Union table.
type UnionInfo struct {
	Id         int32        `db:"id"`
	Name       string       `db:"name"`
	UpazillaId int32        `db:"upazilla_id"`
	CreatedAt  time.Time    `db:"created_at"`
	UpdatedAt  sql.NullTime `db:"updated_at"`
	CreatedBy  string       `db:"created_by"`
	UpdatedBy  string       `db:"updated_by"`
}

// VillageOrRoad struct represents the VillageOrRoad table.
type VillageOrRoad struct {
	Id        int32        `db:"id"`
	Name      string       `db:"name"`
	UnionId   int32        `db:"union_id"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
	CreatedBy string       `db:"created_by"`
	UpdatedBy string       `db:"updated_by"`
}

type Address struct {
	Id              int32  `db:"id"`
	Info            string `db:"info"`
	StudentId       string `db:"student_id"`
	VillageOrRoadId int32  `db:"village_or_road_id"`
	UnionID         int32  `db:"union_id"`
	UpazillaId      int32  `db:"upazilla_id"`
	DistrictId      int32  `db:"district_id"`
	IsPresent       bool   `db:"is_present"`
	IsPermanent     bool   `db:"is_permanent"`
}

type SearchData struct {
	District     int32   `json:"district"`
	Upazillas    []int32 `json:"upazillas"`
	Unions       []int32 `json:"unions"`
	VillageRoads []int32 `json:"villageRoads"`
	SessionId    []int32 `json:"sessionId"`
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

	GetStudentProfileById(studentId string) (Student, error)

	InsertAddress(u Address) (int32, error)
	InsertDistrict(u *District) (int32, error)
	InsertUpazilla(u *Upazilla) (int32, error)
	InsertUnion(u *UnionInfo) (int32, error)
	InsertVillageOrRoad(u *VillageOrRoad) (int32, error)

	GetAddressByStudentID(studentID string) (Address, error)
	GetDistrictByName(districtName string) (District, error)
	GetUpazillaByName(upazillaName string) (Upazilla, error)
	GetUnionByName(unionName string) (UnionInfo, error)
	GetVillageOrRoadByName(villageOrRoadName string) (VillageOrRoad, error)
	GetDistrictByID(districtID int32) (District, error)
	GetUpazillaByID(upazillaID int32) (Upazilla, error)
	GetUnionByID(unionID int32) (UnionInfo, error)
	GetVillageOrRoadByID(villageOrRoadID int32) (VillageOrRoad, error)
	GetUpazillasByDistrictID(districtID int32) ([]Upazilla, error)
	GetVillagesOrRoadsByUnionID(unionID int32) ([]VillageOrRoad, error)
	GetUnionsByUpazillaID(upazillaID int32) ([]UnionInfo, error)
	AddressListByStudentID(studentID string) ([]*Address, error)

	GetDistrictList() ([]*District, error)
	SearchStudent(SearchData) ([]*Student, error)
	// GetDistrictList()([]*District, error)
}
