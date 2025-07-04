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
	v1 := router.Group("/v1")
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
