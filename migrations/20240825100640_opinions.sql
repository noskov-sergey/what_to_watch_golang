-- +goose Up
CREATE TABLE opinions (
    id SERIAL PRIMARY KEY,
    title varchar,
    text varchar,
    source varchar,
    added_by varchar,
    created_at timestamp not null default now()
);

-- +goose Down
DROP TABLE opinions;
