package query

import (
	"database/sql"
	"strconv"
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
	if err != nil {
		return err

	}
	defer stmt.Close()

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
	qc := `SELECT session_id, year, academic_class.name as class FROM academic_session
	left join academic_class on academic_class.id=academic_session.academic_class_id
	`
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

func (q *QueryInit) GetStudentProfileById(studentId string) (Student, error) {
	var student Student

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
    student.updated_by,
    
    vor.name as village_or_road, 
	ui.name as union, 
	u.name as upzilla, 
	d.name as district
    
	FROM
		student 
	left join address a on a.student_id  = student.id  
	left join village_or_road vor ON
	a.village_or_road_id =vor.id 
	left join union_info ui on ui.id=a.union_id 
	left join upazilla u on u.id = a.upazilla_id 
	left join district d on d.id =a.district_id 

	where student.id=$1;`

	if err := q.Db.Get(&student, qc, studentId); err != nil {
		if err == sql.ErrNoRows {
			return student, NotFound
		}
		return student, err
	}

	return student, nil
}

func (q *QueryInit) SearchStudent(searchData SearchData) ([]*Student, error) {
	var students []*Student

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
    student.updated_by,
    
    vor.name as village_or_road, 
	ui.name as union, 
	u.name as upzilla, 
	d.name as district
    
	FROM
		student 
	left join address a on a.student_id  = student.id  
	left join village_or_road vor ON
	a.village_or_road_id =vor.id 
	left join union_info ui on ui.id=a.union_id 
	left join upazilla u on u.id = a.upazilla_id 
	left join district d on d.id =a.district_id
	left join student_academic_session sae on sae.student_id=a.student_id
    WHERE
    d.id = $1`

	// Declare slices to hold parameters
	var upazillaParams, unionParams, villageRoadParams []interface{}

	// Add filters for Upazillas
	if len(searchData.Upazillas) > 0 {
		qc += " AND u.id IN ("
		upazillaParams = make([]interface{}, len(searchData.Upazillas))
		for i, upazilla := range searchData.Upazillas {
			qc += "$" + strconv.Itoa(i+2) + ", "
			upazillaParams[i] = upazilla
		}
		qc = qc[:len(qc)-2] + ")" // Remove the trailing comma and space
	}

	// Add filters for Unions
	if len(searchData.Unions) > 0 {
		qc += " AND ui.id IN ("
		unionParams = make([]interface{}, len(searchData.Unions))
		for i, union := range searchData.Unions {
			qc += "$" + strconv.Itoa(i+len(searchData.Upazillas)+2) + ", "
			unionParams[i] = union
		}
		qc = qc[:len(qc)-2] + ")" // Remove the trailing comma and space
	}

	// Add filters for VillageRoads
	if len(searchData.VillageRoads) > 0 {
		qc += " AND vor.id IN ("
		villageRoadParams = make([]interface{}, len(searchData.VillageRoads))
		for i, villageRoad := range searchData.VillageRoads {
			qc += "$" + strconv.Itoa(i+len(searchData.Upazillas)+len(searchData.Unions)+2) + ", "
			villageRoadParams[i] = villageRoad
		}
		qc = qc[:len(qc)-2] + ")" // Remove the trailing comma and space
	}

	if len(searchData.SessionId) > 0 {
		qc += " AND sae.session_id IN ("
		villageRoadParams = make([]interface{}, len(searchData.SessionId))
		for i, sessionId := range searchData.SessionId {
			qc += "$" + strconv.Itoa(i+len(searchData.Upazillas)+len(searchData.Unions)+2) + ", "
			villageRoadParams[i] = sessionId
		}
		qc = qc[:len(qc)-2] + ")" // Remove the trailing comma and space
	}

	qc += ";"

	// Prepare parameter slice
	params := make([]interface{}, 1)
	params[0] = searchData.District
	params = append(params, upazillaParams...)
	params = append(params, unionParams...)
	params = append(params, villageRoadParams...)

	if err := q.Db.Select(&students, qc, params...); err != nil {
		if err == sql.ErrNoRows {
			return students, NotFound
		}
		return students, err
	}

	return students, nil
}
