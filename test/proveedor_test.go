package test

import (
	"context"
	"testing"

	db "github.com/AguileraFacundo/caja-simple/internal/db/sqlc"
	"github.com/AguileraFacundo/caja-simple/internal/util"
	"github.com/stretchr/testify/require"
)

func createRandomSupplier(t *testing.T) db.Supplier {
	name := util.RandomName()
	proveedor, err := TestQueries.CreateSupplier(context.Background(), name)
	require.NoError(t, err)
	require.NotEmpty(t, proveedor)
	return proveedor
}

func TestCreateProveedor(t *testing.T) {
	supplier := createRandomSupplier(t)
	require.NotEmpty(t, supplier)
}

func TestListProveedor(t *testing.T) {
	for range 10 {
		createRandomSupplier(t)
	}

	proveedores, err := TestQueries.ListSupplier(context.Background(), db.ListSupplierParams{
		Limit:  5,
		Offset: 5,
	})

	require.NoError(t, err)
	require.Len(t, proveedores, 5)
}

func TestGetProveedor(t *testing.T) {
	supplier := createRandomSupplier(t)
	check, err := TestQueries.GetSupplier(context.Background(), supplier.ID)
	require.NoError(t, err)
	require.NotEmpty(t, check)
	require.Equal(t, supplier.ID, check.ID)
	require.Equal(t, supplier.Name, check.Name)
	require.NotZero(t, supplier.ID)
	require.NotZero(t, check.ID)
	require.NotZero(t, supplier.CreationDate)
	require.NotZero(t, check.CreationDate)
}

func TestDeteleProveedor(t *testing.T) {
	proveedor := createRandomSupplier(t)
	err := TestQueries.DeleteSupplier(context.Background(), proveedor.ID)
	require.NoError(t, err)

	check, err := TestQueries.GetSupplier(context.Background(), proveedor.ID)
	require.Error(t, err)
	require.Empty(t, check)
}

func TestUpdateProveedor(t *testing.T) {
	proveedor := createRandomSupplier(t)
	arg := db.UpdateSupplierParams{
		ID:   proveedor.ID,
		Name: util.RandomName(),
	}

	check, err := TestQueries.UpdateSupplier(context.Background(), arg)
	require.NoError(t, err)
	require.NotEqual(t, proveedor.Name, check.Name)
}
