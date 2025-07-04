package api

import (
	"net/http"

	db "github.com/AguileraFacundo/caja-simple/internal/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func (server *Server) createDebtApi(ctx *gin.Context) {
	var req struct {
		Supplierid int64 `json:"supplier_id" binding:"required,min=1"`
		Balance    int64 `json:"balance" binding:"required,min=0"`
		Paid       bool  `json:"paid" binding:"required"`
	}

	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateDebtParams{
		Balance:    req.Balance,
		SupplierID: req.Supplierid,
		Paid:       req.Paid,
	}

	debt, err := server.db.CreateDebt(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, debt)
}

func (server *Server) getDebtApi(ctx *gin.Context) {
	var req struct {
		ID int64 `uri:"id" binding:"required,min=1"`
	}

	if err := ctx.BindUri(&req); err != nil {
		if err == pgx.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	debt, err := server.db.GetDebt(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, debt)
}

func (server *Server) deleteDebtApi(ctx *gin.Context) {
	var req struct {
		ID int64 `uri:"id" binding:"required,min=1"`
	}

	if err := ctx.BindUri(&req); err != nil {
		if err == pgx.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.db.DeleteDebt(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusNoContent, nil)

}

func (server *Server) updateDebtApi(ctx *gin.Context) {
	var req struct {
		Id      int64 `json:"id" binding:"required,min=1"`
		Balance int64 `json:"balance" biding:"required,min=1"`
	}

	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateDebtParams{
		ID:      req.Id,
		Balance: req.Balance,
	}

	debt, err := server.db.UpdateDebt(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusAccepted, debt)
}

func (server *Server) listDebtsApi(ctx *gin.Context) {
	var req struct {
		PageID   int32 `form:"page_id" binding:"required,min=1"`
		PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
	}

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListDebtsParams{
		Limit:  req.PageID,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	debts, err := server.db.ListDebts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, debts)
}
