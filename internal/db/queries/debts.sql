-- name: CreateDebt :one
INSERT INTO debts (balance, supplier_id, paid)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetDebt :one
SELECT * FROM debts
WHERE id = $1
LIMIT 1;

-- name: ListDebts :many
SELECT * FROM debts
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateDebt :one
UPDATE debts
SET balance = $2
WHERE id = $1
RETURNING *;

-- name: DeleteDebt :exec
DELETE FROM supplier
WHERE id = $1;