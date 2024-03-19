package query

import "database/sql"

func (q *QueryInit) GetAddressByStudentID(studentID string) (Address, error) {
	var address Address
	const query = `SELECT * FROM address WHERE student_id=$1`
	if err := q.Db.Get(&address, query, studentID); err != nil {
		if err == sql.ErrNoRows {
			return address, NotFound
		}
		return address, err
	}
	return address, nil
}

func (q *QueryInit) GetDistrictByName(districtName string) (District, error) {
	var district District
	const query = `SELECT * FROM district WHERE name=$1`
	if err := q.Db.Get(&district, query, districtName); err != nil {
		if err == sql.ErrNoRows {
			return district, NotFound
		}
		return district, err
	}
	return district, nil
}

func (q *QueryInit) GetUpazillaByName(upazillaName string) (Upazilla, error) {
	var upazilla Upazilla
	const query = `SELECT * FROM upazilla WHERE name=$1`
	if err := q.Db.Get(&upazilla, query, upazillaName); err != nil {
		if err == sql.ErrNoRows {
			return upazilla, NotFound
		}
		return upazilla, err
	}
	return upazilla, nil
}

func (q *QueryInit) GetUnionByName(unionName string) (UnionInfo, error) {
	var unionInfo UnionInfo
	const query = `SELECT * FROM union_info WHERE name=$1`
	if err := q.Db.Get(&unionInfo, query, unionName); err != nil {
		if err == sql.ErrNoRows {
			return unionInfo, NotFound
		}
		return unionInfo, err
	}
	return unionInfo, nil
}

func (q *QueryInit) GetVillageOrRoadByName(villageOrRoadName string) (VillageOrRoad, error) {
	var villageOrRoad VillageOrRoad
	const query = `SELECT * FROM village_or_road WHERE name=$1`
	if err := q.Db.Get(&villageOrRoad, query, villageOrRoadName); err != nil {
		if err == sql.ErrNoRows {
			return villageOrRoad, NotFound
		}
		return villageOrRoad, err
	}
	return villageOrRoad, nil
}

func (q *QueryInit) GetDistrictByID(districtID int32) (District, error) {
	var district District
	const query = `SELECT * FROM district WHERE id=$1`
	if err := q.Db.Get(&district, query, districtID); err != nil {
		if err == sql.ErrNoRows {
			return district, NotFound
		}
		return district, err
	}
	return district, nil
}
func (q *QueryInit) GetUpazillaByID(upazillaID int32) (Upazilla, error) {
	var upazilla Upazilla
	const query = `SELECT * FROM upazilla WHERE id=$1`
	if err := q.Db.Get(&upazilla, query, upazillaID); err != nil {
		if err == sql.ErrNoRows {
			return upazilla, NotFound
		}
		return upazilla, err
	}
	return upazilla, nil
}

func (q *QueryInit) GetUnionByID(unionID int32) (UnionInfo, error) {
	var unionInfo UnionInfo
	const query = `SELECT * FROM union_info WHERE id=$1`
	if err := q.Db.Get(&unionInfo, query, unionID); err != nil {
		if err == sql.ErrNoRows {
			return unionInfo, NotFound
		}
		return unionInfo, err
	}
	return unionInfo, nil
}
func (q *QueryInit) GetVillageOrRoadByID(villageOrRoadID int32) (VillageOrRoad, error) {
	var villageOrRoad VillageOrRoad
	const query = `SELECT * FROM village_or_road WHERE id=$1`
	if err := q.Db.Get(&villageOrRoad, query, villageOrRoadID); err != nil {
		if err == sql.ErrNoRows {
			return villageOrRoad, NotFound
		}
		return villageOrRoad, err
	}
	return villageOrRoad, nil
}

func (q *QueryInit) GetUpazillasByDistrictID(districtID int32) ([]Upazilla, error) {
	var upazillas []Upazilla
	const query = `SELECT * FROM upazilla WHERE district_id=$1`
	if err := q.Db.Select(&upazillas, query, districtID); err != nil {
		if err == sql.ErrNoRows {
			return upazillas, NotFound
		}
		return upazillas, err
	}
	return upazillas, nil
}
func (q *QueryInit) GetVillagesOrRoadsByUnionID(unionID int32) ([]VillageOrRoad, error) {
	var villagesOrRoads []VillageOrRoad
	const query = `SELECT * FROM village_or_road WHERE union_id=$1`
	if err := q.Db.Select(&villagesOrRoads, query, unionID); err != nil {
		if err == sql.ErrNoRows {
			return villagesOrRoads, NotFound
		}
		return villagesOrRoads, err
	}
	return villagesOrRoads, nil
}
func (q *QueryInit) GetUnionsByUpazillaID(upazillaID int32) ([]UnionInfo, error) {
	var unions []UnionInfo
	const query = `SELECT * FROM union_info WHERE upazilla_id=$1`
	if err := q.Db.Select(&unions, query, upazillaID); err != nil {
		if err == sql.ErrNoRows {
			return unions, NotFound
		}
		return unions, err
	}
	return unions, nil
}

func (q *QueryInit) AddressListByStudentID(studentID string) ([]*Address, error) {
	var addresses []*Address
	query := `SELECT id,
        info,
        student_id,
        village_or_road_id,
        union_id,
        upazilla_id,
        district_id,
        is_present,
        is_permanent
    FROM address
    WHERE student_id = $1`
	if err := q.Db.Select(&addresses, query, studentID); err != nil {
		if err == sql.ErrNoRows {
			return nil, NotFound
		}
		return nil, err
	}
	return addresses, nil
}

func (q *QueryInit) GetDistrictList() ([]*District, error) {
	var districts []*District
	const query = `SELECT * FROM district order by name ASC`
	if err := q.Db.Select(&districts, query); err != nil {
		if err == sql.ErrNoRows {
			return districts, NotFound
		}
		return districts, err
	}
	return districts, nil
}
