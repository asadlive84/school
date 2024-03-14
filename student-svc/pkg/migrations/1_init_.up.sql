-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


CREATE TABLE IF NOT EXISTS academic_class (
    id SERIAL PRIMARY KEY,
    name VARCHAR (300) NOT NULL,
    numeric_name VARCHAR (10) NOT NULL,
    created_at timestamp default current_timestamp,
    updated_at timestamp NULL,
    created_by VARCHAR (300) NOT NULL DEFAULT '',
    updated_by VARCHAR (300) NOT NULL DEFAULT ''
    
);



CREATE TABLE IF NOT EXISTS academic_session (
    session_id SERIAL PRIMARY KEY,
    year INTEGER,
    academic_class_id int NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT FALSE,
    start_date timestamp NULL,
    end_date timestamp NULL,
    created_by VARCHAR (300) NOT NULL DEFAULT '',
    updated_by VARCHAR (300) NOT NULL DEFAULT '',
    FOREIGN KEY (academic_class_id) REFERENCES academic_class(id) ON DELETE CASCADE
);



CREATE TABLE IF NOT EXISTS student(
   id uuid DEFAULT uuid_generate_v4 () PRIMARY KEY ,
   name VARCHAR (300) NOT NULL,

   std_id VARCHAR (300) NOT NULL,
   name_bn VARCHAR (300) NOT NULL,

   email VARCHAR (300) NOT NULL DEFAULT '',

   fathers_name VARCHAR (300) NOT NULL,
   mothers_name VARCHAR (300) NOT NULL,
   dob TIMESTAMP WITHOUT TIME ZONE NULL,
   gender VARCHAR(6) CHECK (gender IN ('MALE', 'FEMALE')),
   blood_group VARCHAR(2) CHECK (blood_group IN ('A+', 'B+','-')) NOT NULL DEFAULT '-',


   mobile_number VARCHAR (20) NOT NULL DEFAULT '',
   session VARCHAR (300) NOT NULL DEFAULT '',

   religion VARCHAR(10) CHECK (religion IN ('ISLAM', 'HINDU', 'BUDDA')) NOT NULL DEFAULT '-',

   class_name VARCHAR (300) NOT NULL DEFAULT '',
   address TEXT  CHECK (length(address) <= 500) NULL DEFAULT '',
   class_section VARCHAR (300) NOT NULL DEFAULT '',


   is_active BOOLEAN DEFAULT FALSE,
   created_at timestamp default current_timestamp,
   updated_at timestamp NULL,
   created_by VARCHAR (300) NOT NULL DEFAULT '',
   updated_by VARCHAR (300) NOT NULL DEFAULT ''
);

-- Link academic_session to students (many-to-many relationship)
CREATE TABLE IF NOT EXISTS student_academic_session (
    student_id uuid REFERENCES student(id) ON DELETE CASCADE,
    session_id INTEGER REFERENCES academic_session(session_id),
    PRIMARY KEY (student_id, session_id)
);


-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE student_academic_session;
DROP TABLE academic_session;
DROP TABLE student;