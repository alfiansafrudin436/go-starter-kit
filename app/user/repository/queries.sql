-- name: GetAllUsers :many
SELECT id, name, email, role, is_active, created_at, updated_at
FROM users
WHERE is_active = true
ORDER BY created_at DESC;

-- name: GetUserByID :one
SELECT id, name, email, role, is_active, created_at, updated_at
FROM users
WHERE id = $1
LIMIT 1;

-- name: UpdateUserName :exec
UPDATE users
SET name = $2,
    updated_at = now()
WHERE id = $1;

-- name: DeactivateUser :exec
UPDATE users
SET is_active = false,
    updated_at = now()
WHERE id = $1;
