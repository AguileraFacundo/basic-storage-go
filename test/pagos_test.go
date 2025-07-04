package test

import (
	"context"
	"testing"

	db "github.com/AguileraFacundo/caja-simple/internal/db/sqlc"
	"github.com/AguileraFacundo/caja-simple/internal/util"
	"github.com/stretchr/testify/require"
)

func randomPago(t *testing.T) db.Payment {
	supplier := createRandomSupplier(t)
	arg := db.CreatePaymentParams{
		Balance:    util.RandomMoney(),
		SupplierID: supplier.ID,
	}

	payment, err := TestQueries.CreatePayment(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, payment)
	return payment
}

func TestSinglePago(t *testing.T) {
	payment := randomPago(t)
	require.NotEmpty(t, payment)

	check, err := TestQueries.GetPayment(context.Background(), payment.ID)
	require.NoError(t, err)
	require.NotEmpty(t, check)
	require.Equal(t, payment.ID, check.ID)
	require.Equal(t, payment.SupplierID, check.SupplierID)
	require.Equal(t, payment.Balance, check.Balance)
	require.NotZero(t, check.ID)
	require.NotZero(t, check.Balance)
}

func TestListPagos(t *testing.T) {
	for range 10 {
		randomPago(t)
	}

	pagos, err := TestQueries.ListPayments(context.Background(), db.ListPaymentsParams{
		Limit:  5,
		Offset: 5,
	})

	require.NoError(t, err)
	require.Len(t, pagos, 5)
}

func TestDelete(t *testing.T) {
	payment := randomPago(t)
	err := TestQueries.DeletePayment(context.Background(), payment.ID)
	require.NoError(t, err)

	check, err := TestQueries.GetPayment(context.Background(), payment.ID)
	require.Error(t, err)
	require.Empty(t, check)
}
