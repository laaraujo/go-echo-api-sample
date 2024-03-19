-- name: GetRoutine :one
SELECT * FROM routines
WHERE id = $1 LIMIT 1;

-- name: ListRoutines :many
SELECT * FROM routines
ORDER BY name;

-- name: CreateRoutine :one
INSERT INTO routines (
  name
) VALUES (
  $1
)
RETURNING *;

-- name: UpdateRoutine :exec
UPDATE routines
  set name = $2
WHERE id = $1
RETURNING *;

-- name: DeleteRoutine :exec
DELETE FROM routines
WHERE id = $1;
