-- name: GetExercise :one
SELECT * FROM exercises
WHERE id = $1 LIMIT 1;

-- name: ListExercises :many
SELECT * FROM exercises
ORDER BY name;

-- name: ListRoutineExercises :many
SELECT * FROM exercises
WHERE routine_id = $1
ORDER BY name;

-- name: CreateExercise :one
INSERT INTO exercises (
  name, sets, reps, routine_id
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: UpdateExercise :exec
UPDATE exercises
  set name = $2,
      sets = $3,
      reps = $4,
      routine_id = $5
WHERE id = $1
RETURNING *;

-- name: DeleteExercise :exec
DELETE FROM exercises
WHERE id = $1;
