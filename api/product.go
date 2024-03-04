package api

import (
	"net/http"

	db "github.com/bjamyl/begho/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createProductRequest struct {
	UserID      int64    `json:"user_id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	StartPrice  int64    `json:"start_price"`
	Images      []string `json:"images"`
	Watchers    []int64  `json:"watchers"`
}

// Add new product
func (server *Server) createProduct(ctx *gin.Context) {
	var req createProductRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	arg := db.CreateProductParams{
		UserID:      req.UserID,
		Name:        req.Name,
		Description: req.Description,
		StartPrice:  req.StartPrice,
		Images:      req.Images,
		Watchers:    req.Watchers,
	}

	product, err := server.store.CreateProduct(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, product)

}

type listProductsRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,max=20"`
}

// Get all products
func (server *Server) listProducts(ctx *gin.Context) {
	var req listProductsRequest
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	arg := db.ListProductsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	products, err := server.store.ListProducts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, products)
}

// List all products belonging to a particular user
type listUserProductsRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) listUserProducts(ctx *gin.Context) {
	var req listUserProductsRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	products, err := server.store.ListUserProducts(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, products)

}
