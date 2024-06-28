-- +goose Up
CREATE TABLE cats (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),

    name varchar(128),
    years_of_experience integer,
    breed varchar(128),
    salary_in_cents bigint
);

-- +goose Down
DROP TABLE IF EXISTS cats;
