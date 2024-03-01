-- name: CreateUser :one
INSERT INTO users (
  username,
  full_name,
  email,
  password
) VALUES (
  $1, $2, $3, $4
)RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY username;