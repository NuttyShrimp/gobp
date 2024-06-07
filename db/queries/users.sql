-- name: GetUser :one
SELECT *
FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByUid :one
SELECT *
FROM users
WHERE uid = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (name, email, uid)
VALUES ($1, $2, $3)
RETURNING *;
