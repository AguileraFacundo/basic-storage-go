// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: sales.sql

package db

import (
	"context"
)

const createSale = `-- name: CreateSale :one
INSERT INTO sales (balance)
VALUES ($1)
RETURNING id, balance, date
`

func (q *Queries) CreateSale(ctx context.Context, balance int64) (Sale, error) {
	row := q.db.QueryRow(ctx, createSale, balance)
	var i Sale
	err := row.Scan(&i.ID, &i.Balance, &i.Date)
	return i, err
}

const deleteSale = `-- name: DeleteSale :exec
DELETE FROM sales
WHERE id = $1
`

func (q *Queries) DeleteSale(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteSale, id)
	return err
}

const getSale = `-- name: GetSale :one
SELECT id, balance, date FROM sales
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetSale(ctx context.Context, id int64) (Sale, error) {
	row := q.db.QueryRow(ctx, getSale, id)
	var i Sale
	err := row.Scan(&i.ID, &i.Balance, &i.Date)
	return i, err
}

const lastTenSales = `-- name: LastTenSales :many
SELECT id, balance, date FROM sales
ORDER BY date ASC
LIMIT 10
`

func (q *Queries) LastTenSales(ctx context.Context) ([]Sale, error) {
	rows, err := q.db.Query(ctx, lastTenSales)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Sale
	for rows.Next() {
		var i Sale
		if err := rows.Scan(&i.ID, &i.Balance, &i.Date); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listSales = `-- name: ListSales :many
SELECT id, balance, date FROM sales
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListSalesParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) ListSales(ctx context.Context, arg ListSalesParams) ([]Sale, error) {
	rows, err := q.db.Query(ctx, listSales, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Sale
	for rows.Next() {
		var i Sale
		if err := rows.Scan(&i.ID, &i.Balance, &i.Date); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateSale = `-- name: UpdateSale :one
UPDATE sales
SET balance = $2
WHERE id = $1
RETURNING id, balance, date
`

type UpdateSaleParams struct {
	ID      int64
	Balance int64
}

func (q *Queries) UpdateSale(ctx context.Context, arg UpdateSaleParams) (Sale, error) {
	row := q.db.QueryRow(ctx, updateSale, arg.ID, arg.Balance)
	var i Sale
	err := row.Scan(&i.ID, &i.Balance, &i.Date)
	return i, err
}
