-- +goose Up
CREATE TABLE tips (
    id serial PRIMARY KEY,
    user_id integer NOT NULL CHECK (user_id > 0),
    problem_id integer NOT NULL CHECK (problem_id > 0),
    text text NOT NULL
);

-- +goose Down
DROP TABLE tips;
