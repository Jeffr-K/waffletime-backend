package presentor

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"waffletime/internal/product/application/usecase/product"
)

type IProductController interface {
	PublishProduct(context *gin.Context)
	DeleteProduct(context *gin.Context)
	UpdateProduct(context *gin.Context)
	GetProduct(context *gin.Context)
	GetProductList(context *gin.Context)
}

type productController struct {
	productPublishUseCase product.IProductPublishUseCase
	productDeleteUseCase  product.IProductDeleteUseCase
	productSearchUseCase  product.IProductSearchUseCase
	productUpdateUseCase  product.IProductUpdateUseCase
}

func NewProductController() IProductController {
	return &productController{
		productPublishUseCase: product.NewProductPublishUseCase(),
		productDeleteUseCase:  product.NewProductDeleteUseCase(),
		productSearchUseCase:  product.NewProductSearchUseCase(),
		productUpdateUseCase:  product.NewProductUpdateUseCase(),
	}
}

func (pc productController) PublishProduct(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"result": "successfully saved."})
}

func (pc productController) DeleteProduct(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"result": "successfully saved."})
}

func (pc productController) UpdateProduct(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"result": "successfully saved."})
}

func (pc productController) GetProduct(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"result": "successfully saved."})
}

func (pc productController) GetProductList(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"result": "successfully saved."})
}
