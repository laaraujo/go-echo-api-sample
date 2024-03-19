-- +goose Up
-- +goose StatementBegin
ALTER TABLE exercises ADD routine_id INTEGER NOT NULL DEFAULT 0;
ALTER TABLE exercises ADD CONSTRAINT fk_exercises_routines FOREIGN KEY (routine_id) REFERENCES routines(id) ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE exercises
DROP CONSTRAINT fk_exercises_routines;

ALTER TABLE exercises
DROP COLUMN routine_id;
-- +goose StatementEnd
