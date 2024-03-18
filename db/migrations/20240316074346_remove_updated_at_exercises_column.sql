-- +goose Up
-- +goose StatementBegin
ALTER TABLE exercises
DROP COLUMN updated_at;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE exercises
ADD updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP;
-- +goose StatementEnd
