-- name: CreateUser :one
INSERT INTO users(email, name)
VALUES (
    $1,
    $2
)
RETURNING *;

-- name: GetUsers :many
SELECT * FROM users;

-- name: GetUser :one
SELECT * FROM users WHERE email = $1;

-- name: DeleteUser :exec
DELETE FROM users WHERE email = $1;

-- name: UpdateUserNameByEmail :exec
UPDATE users 
SET name = $1
WHERE email = $2;