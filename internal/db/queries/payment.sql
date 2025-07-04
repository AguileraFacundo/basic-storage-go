-- name: CreatePayment :one
INSERT INTO payments (balance, supplier_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetPayment :one
SELECT * FROM payments
WHERE id = $1
FOR NO KEY UPDATE;

-- name: ListPayments :many
SELECT * FROM payments
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdatePayment :one
UPDATE payments
SET balance = $2
WHERE id = $1
RETURNING *;

-- name: DeletePayment :exec
DELETE FROM payments
WHERE id = $1;