-- +goose Up
-- +goose StatementBegin
CREATE TABLE exercises (
 id SERIAL PRIMARY KEY,
 name VARCHAR(255) NOT NULL,
 sets SMALLINT NOT NULL,
 reps SMALLINT NOT NULL,
 created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
 updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE exercises;
-- +goose StatementEnd
