package api

import (
	db "github.com/AguileraFacundo/caja-simple/internal/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	db     *db.Queries
	router *gin.Engine
}

func NewServer(db *db.Queries) *Server {
	server := &Server{
		db: db,
	}
	router := gin.Default()
	server.router = router
	v1 := router.Group("/api/v1")
	//suppliers
	v1.POST("/supplier", server.createSupplierApi)
	v1.DELETE("/supplier/:id", server.deleteSupplierApi)
	v1.GET("/supplier/:id", server.getSupplierApi)
	v1.GET("/suppliers", server.listSupplierApi)
	v1.PUT("/supplier", server.updateProveedorApi)
	//debts
	v1.POST("/debt", server.createDebtApi)
	v1.GET("/debt/:id", server.getDebtApi)
	v1.GET("/debts", server.listDebtsApi)
	v1.DELETE("/debt/:id", server.deleteDebtApi)
	v1.PUT("/debt", server.updateDebtApi)
	//payments
	v1.POST("/payment", server.createPaymentApi)
	v1.GET("/payment/:id", server.getPaymentApi)
	v1.DELETE("/payment/:id", server.deletePaymentApi)
	v1.GET("/payments", server.listPaymentsApi)
	v1.PUT("/payment", server.updatePaymentApi)
	//sales
	v1.POST("/sale", server.createSaleApi)
	v1.DELETE("/sale/:id", server.deleteSaleApi)
	v1.GET("/sale/:id", server.getSaleApi)
	v1.GET("/sales", server.listSaleApi)
	v1.PUT("/sale", server.updateSaleApi)

	return server
}

func errorResponse(err error) gin.H {
	return gin.H{
		"Error": err.Error(),
	}
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
