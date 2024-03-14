package query

func (q *QueryInit) InsertAddress(a Address) (int32, error) {
	// Prepare the SQL statement
	stmt, err := q.Db.PrepareNamed(`
        INSERT INTO address (
            info,
            student_id,
            village_or_road_id,
            union_id,
            upazilla_id,
            district_id,
            is_present,
            is_permanent
        ) 
        VALUES (
            :info,
            :student_id,
            :village_or_road_id,
            :union_id,
            :upazilla_id,
            :district_id,
            :is_present,
            :is_permanent
        ) 
        RETURNING id;
    `)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	// Execute the SQL statement to insert the address and retrieve the ID
	var id int32
	err = stmt.QueryRowx(a).Scan(&id)
	if err != nil {
		return 0, err
	}

	// Return the ID of the newly inserted address
	return id, nil
}

func (q *QueryInit) InsertDistrict(d *District) (int32, error) {
	var id int32
	query := `INSERT INTO district (name) VALUES ($1) RETURNING id`
	err := q.Db.QueryRow(query, d.Name).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (q *QueryInit) InsertUpazilla(u *Upazilla) (int32, error) {
	var id int32
	query := `INSERT INTO upazilla (name, district_id) VALUES ($1, $2) RETURNING id`
	err := q.Db.QueryRow(query, u.Name, u.DistrictId).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (q *QueryInit) InsertUnion(u *UnionInfo) (int32, error) {
	var id int32
	query := `INSERT INTO union_info (name, upazilla_id) VALUES ($1, $2) RETURNING id`
	err := q.Db.QueryRow(query, u.Name, u.UpazillaId).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (q *QueryInit) InsertVillageOrRoad(v *VillageOrRoad) (int32, error) {
	var id int32
	query := `INSERT INTO village_or_road (name, union_id) VALUES ($1, $2) RETURNING id`
	err := q.Db.QueryRow(query, v.Name, v.UnionId).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
