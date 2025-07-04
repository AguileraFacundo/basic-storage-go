-- name: CreateSale :one
INSERT INTO sales (balance)
VALUES ($1)
RETURNING *;

-- name: GetSale :one
SELECT * FROM sales
WHERE id = $1
LIMIT 1;

-- name: ListSales :many
SELECT * FROM sales
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateSale :one
UPDATE sales
SET balance = $2
WHERE id = $1
RETURNING *;

-- name: DeleteSale :exec
DELETE FROM sales
WHERE id = $1;