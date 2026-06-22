-- name: GetUserByEmail :one
SELECT id, name, email, password, role, is_active, created_at, updated_at
FROM users
WHERE email = $1 AND is_active = true
LIMIT 1;

-- name: GetUserByID :one
SELECT id, name, email, role, is_active, created_at, updated_at
FROM users
WHERE id = $1
LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (email, name, password, role, is_active, created_at)
VALUES ($1, $2, $3, $4, $5, now())
RETURNING id;

-- name: UpdateUserPassword :exec
UPDATE users
SET password = $2,
    reset_password_token = null,
    updated_at = now()
WHERE id = $1;

-- name: UpdateUserResetToken :exec
UPDATE users
SET reset_password_token = $2,
    updated_at = now()
WHERE id = $1;

-- name: GetUserByResetToken :one
SELECT id, name, email
FROM users
WHERE reset_password_token = $1 AND is_active = true
LIMIT 1;

-- name: DeactivateUser :exec
UPDATE users
SET is_active = false,
    updated_at = now()
WHERE id = $1;
