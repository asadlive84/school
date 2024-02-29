-- +goose Up

CREATE TABLE IF NOT EXISTS users(
   id uuid DEFAULT uuid_generate_v4 (),
   password VARCHAR (300) NOT NULL,
   phone VARCHAR (20) NOT NULL DEFAULT '',
   email VARCHAR (300) UNIQUE NOT NULL,
   is_active BOOLEAN DEFAULT FALSE,
   created timestamp default current_timestamp,
   updated timestamp NULL
);

-- +goose Down
DROP TABLE users;