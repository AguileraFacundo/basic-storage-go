package api

import (
	"net/http"

	db "github.com/AguileraFacundo/caja-simple/internal/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func (server *Server) createSaleApi(ctx *gin.Context) {
	var req struct {
		Balance int64 `json:"balance" binding:"required,min=0"`
	}

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	sale, err := server.db.CreateSale(ctx, req.Balance)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, sale)
}

func (server *Server) deleteSaleApi(ctx *gin.Context) {
	var req struct {
		ID int64 `uri:"id" binding:"required,min=1"`
	}
	if err := ctx.BindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.db.DeleteSale(ctx, req.ID)
	if err != nil {
		if err == pgx.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func (server *Server) getSaleApi(ctx *gin.Context) {
	var req struct {
		ID int64 `uri:"id" binding:"required,min=1"`
	}
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	sale, err := server.db.GetSale(ctx, req.ID)
	if err != nil {
		if err == pgx.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, sale)
}

func (server *Server) listSaleApi(ctx *gin.Context) {
	var req struct {
		PageID   int32 `form:"page_id" binding:"required,min=1"`
		PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
	}
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListSalesParams{
		Limit:  req.PageID,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	sales, err := server.db.ListSales(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, sales)
}

func (server *Server) updateSaleApi(ctx *gin.Context) {
	var req struct {
		ID      int64 `json:"id" biding:"required"`
		Balance int64 `json:"balance" binding:"required,min=0"`
	}

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	arg := db.UpdateSaleParams{
		ID:      req.ID,
		Balance: req.Balance,
	}

	sale, err := server.db.UpdateSale(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusAccepted, sale)
}

func (server *Server) lastTenSalesApi(ctx *gin.Context) {
	lastTen, err := server.db.LastTenSales(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, lastTen)
}
