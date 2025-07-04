-- name: CreateSupplier :one
INSERT INTO supplier (name)
VALUES ($1)
RETURNING *;

-- name: GetSupplier :one
SELECT * FROM supplier
WHERE id = $1
LIMIT 1;

-- name: ListSupplier :many
SELECT * FROM supplier
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateSupplier :one
UPDATE supplier
SET name = $2
WHERE id = $1
RETURNING *;

-- name: DeleteSupplier :exec
DELETE FROM supplier
WHERE id = $1;