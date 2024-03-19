-- +goose Up
-- +goose StatementBegin
CREATE TABLE routines (
 id SERIAL PRIMARY KEY,
 name VARCHAR(255) NOT NULL,
 created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE routines;
-- +goose StatementEnd
