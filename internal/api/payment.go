package api

import (
	"net/http"

	db "github.com/AguileraFacundo/caja-simple/internal/db/sqlc"
	"github.com/gin-gonic/gin"
)

func (server *Server) createPaymentApi(ctx *gin.Context) {
	var req struct {
		Balance    int64 `json:"balance" binding:"required,min=1"`
		SupplierID int64 `json:"supplier_id" binding:"required,min=1"`
	}
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.CreatePaymentParams{
		Balance:    req.Balance,
		SupplierID: req.SupplierID,
	}
	payment, err := server.db.CreatePayment(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, payment)
}

func (server *Server) getPaymentApi(ctx *gin.Context) {
	var req struct {
		ID int64 `form:"id" binding:"required,min=1"`
	}

	if err := ctx.BindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	payment, err := server.db.GetPayment(ctx, req.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, payment)

}

func (server *Server) listPaymentsApi(ctx *gin.Context) {
	var req struct {
		PageID   int32 `form:"page_id" biding:"required,min=1"`
		PageSize int32 `form:"page_size" biding:"required,min=5,max=10"`
	}

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListPaymentsParams{
		Limit:  req.PageID,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	payments, err := server.db.ListPayments(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, payments)
}

func (server *Server) deletePaymentApi(ctx *gin.Context) {
	var req struct {
		ID int64 `form:"id" binding:"required,min=1"`
	}

	if err := ctx.BindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.db.DeletePayment(ctx, req.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

func (server *Server) updatePaymentApi(ctx *gin.Context) {
	var req struct {
		ID      int64 `json:"id" binding:"required,min=1"`
		Balance int64 `json:"balance" binding:"required, min=0" `
	}

	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdatePaymentParams{
		ID:      req.ID,
		Balance: req.Balance,
	}

	payment, err := server.db.UpdatePayment(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, payment)
}
