-- name: GetExercise :one
SELECT * FROM exercises
WHERE id = $1 LIMIT 1;

-- name: ListExercises :many
SELECT * FROM exercises
ORDER BY name;

-- name: CreateExercise :one
INSERT INTO exercises (
  name, sets, reps
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateExercise :exec
UPDATE exercises
  set name = $2,
      sets = $3,
      reps = $4

WHERE id = $1
RETURNING *;

-- name: DeleteExercise :exec
DELETE FROM exercises
WHERE id = $1;
