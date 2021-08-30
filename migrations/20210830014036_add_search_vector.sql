-- +goose Up
ALTER TABLE tips
    ADD COLUMN search_vector tsvector GENERATED ALWAYS AS (
        to_tsvector('russian', text) || to_tsvector('english', text)
    ) STORED;

CREATE INDEX ON tips USING gin(search_vector);

-- +goose Down
ALTER TABLE tips
    DROP COLUMN search_vector;
