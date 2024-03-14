
-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE IF NOT EXISTS district (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    created_at timestamp default current_timestamp,
    updated_at timestamp NULL,
    created_by VARCHAR (300) NOT NULL DEFAULT '',
    updated_by VARCHAR (300) NOT NULL DEFAULT ''
);

CREATE TABLE IF NOT EXISTS upazilla (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    district_id INT,
    created_at timestamp default current_timestamp,
    updated_at timestamp NULL,
    created_by VARCHAR (300) NOT NULL DEFAULT '',
    updated_by VARCHAR (300) NOT NULL DEFAULT '',
    FOREIGN KEY (district_id) REFERENCES district(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS union_info (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    upazilla_id INT,
    created_at timestamp default current_timestamp,
    updated_at timestamp NULL,
    created_by VARCHAR (300) NOT NULL DEFAULT '',
    updated_by VARCHAR (300) NOT NULL DEFAULT '',
    FOREIGN KEY (upazilla_id) REFERENCES upazilla(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS village_or_road (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    union_id INT,
    created_at timestamp default current_timestamp,
    updated_at timestamp NULL,
    created_by VARCHAR (300) NOT NULL DEFAULT '',
    updated_by VARCHAR (300) NOT NULL DEFAULT '',
    FOREIGN KEY (union_id) REFERENCES union_info(id) ON DELETE CASCADE
);


CREATE TABLE IF NOT EXISTS address (
    id SERIAL PRIMARY KEY,
    info VARCHAR (300) NOT NULL DEFAULT '',
    student_id uuid NOT NULL,
    village_or_road_id INT,
    union_id INT,
    upazilla_id INT,
    district_id INT,
    is_present BOOLEAN NOT NULL DEFAULT FALSE,
    is_permanent BOOLEAN NOT NULL DEFAULT FALSE,
    FOREIGN KEY (student_id) REFERENCES student(id) ON DELETE CASCADE,
    FOREIGN KEY (district_id) REFERENCES district(id) ON DELETE CASCADE,
    FOREIGN KEY (upazilla_id) REFERENCES upazilla(id) ON DELETE CASCADE,
    FOREIGN KEY (union_id) REFERENCES union_info(id) ON DELETE CASCADE,
    FOREIGN KEY (village_or_road_id) REFERENCES village_or_road(id) ON DELETE CASCADE,
    created_at timestamp default current_timestamp,
    updated_at timestamp NULL,
    created_by VARCHAR (300) NOT NULL DEFAULT '',
    updated_by VARCHAR (300) NOT NULL DEFAULT ''
);


-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE address;
DROP TABLE district;
DROP TABLE upazilla;
DROP TABLE union_info;
DROP TABLE village_or_road;