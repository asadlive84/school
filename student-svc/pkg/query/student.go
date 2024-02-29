package query

import (
	"database/sql"
)

func (q *QueryInit) InsertUser(u User) error {
	stmt, err := q.Db.PrepareNamed("INSERT INTO student (name, std_id, name_bn, fathers_name, mothers_name, dob, mobile_number, email, session, class_name, gender, is_active) VALUES (:name, :std_id, :name_bn, :fathers_name, :mothers_name, :dob, :mobile_number, :email, :session, :class_name, :gender, :is_active)")
	if err != nil {
		return err

	}
	_, err = stmt.Exec(&u)

	if err != nil {
		return err
	}
	return nil
}

func (q *QueryInit) InsertStudent(u Student) (string, error) {
	// Prepare the SQL statement
	stmt, err := q.Db.PrepareNamed(`
	INSERT INTO student (
		name, 
		std_id, 
		name_bn, 
		fathers_name, 
		mothers_name, 
		dob, 
		mobile_number, 
		email, 
		session, 
		class_name, 
		gender, 
		is_active,
		religion,
		address,
		created_at
	) 
	VALUES (
		:name, 
		:std_id, 
		:name_bn, 
		:fathers_name, 
		:mothers_name, 
		:dob, 
		:mobile_number, 
		:email, 
		:session, 
		:class_name, 
		:gender, 
		:is_active,
		:religion,
		:address,
		NOW()
	) 
	RETURNING id;
	`)
	if err != nil {

		return "", err
	}
	defer stmt.Close()

	// Execute the SQL statement to insert the student and retrieve the ID
	var stdId string
	err = stmt.QueryRow(&u).Scan(&stdId)
	if err != nil {
		return "", err
	}

	// Return the ID of the newly inserted student
	return stdId, nil
}

func (q *QueryInit) InsertAcademicSession(u AcademicSession) error {
	// Prepare the SQL statement
	stmt, err := q.Db.PrepareNamed(`INSERT INTO student_academic_session (student_id, session_id)
	VALUES (:student_id, :session_id);`)
	if err != nil {
		return err
	}
	defer stmt.Close()

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
func (q *QueryInit) StudentList() ([]*Student, error) {
	var user []*Student
	qc := `SELECT id,
    std_id,
    name,
    email,
    name_bn,
    fathers_name,
    mothers_name,
    dob,
    gender,
    blood_group,
    mobile_number,
    session,
    class_name,
    class_section,
    address,
    religion,
    is_active,
    created_at,
    updated_at,
    created_by,
    updated_by
	FROM student`
	if err := q.Db.Select(&user, qc); err != nil {
		if err == sql.ErrNoRows {
			return nil, NotFound
		}
		return nil, err
	}
	return user, nil
}

func (q *QueryInit) GetStudentListBySession(sessionId int32) ([]*Student, error) {
	var user []*Student
	qc := `SELECT
	student.id,
    student.std_id,
    student.name,
    student.email,
    student.name_bn,
    student.fathers_name,
    student.mothers_name,
    student.dob,
    student.gender,
    student.blood_group,
    student.mobile_number,
    student.session,
    student.class_name,
    student.class_section,
    student.address,
    student.religion,
    student.is_active,
    student.created_at,
    student.updated_at,
    student.created_by,
    student.updated_by
FROM
    academic_session
JOIN
    student_academic_session ON academic_session.session_id = student_academic_session.session_id
JOIN
    student ON student_academic_session.student_id = student.id
WHERE student_academic_session.session_id=$1;
`

	if err := q.Db.Select(&user, qc, sessionId); err != nil {
		if err == sql.ErrNoRows {
			return nil, NotFound
		}
		return nil, err
	}
	return user, nil
}

func (q *QueryInit) GetClassList() ([]*Class, error) {
	var class []*Class
	qc := `SELECT session_id, year, class FROM academic_session`
	if err := q.Db.Select(&class, qc); err != nil {
		if err == sql.ErrNoRows {
			return nil, NotFound
		}
		return nil, err
	}
	return class, nil
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
