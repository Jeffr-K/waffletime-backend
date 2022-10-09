package presentor

import (
	"github.com/gin-gonic/gin"
	"github.com/jeffr-k/waffletime/internal/product/application/usecase"
	"net/http"
)

type IProductController interface {
	PublishProduct(context *gin.Context)
	DeleteProduct(context *gin.Context)
	UpdateProduct(context *gin.Context)
	GetProduct(context *gin.Context)
	GetProductList(context *gin.Context)
}

type productController struct {
	productPublishUseCase usecase.IProductPublishUseCase
	productDeleteUseCase  usecase.IProductDeleteUseCase
	productSearchUseCase  usecase.IProductSearchUseCase
	productUpdateUseCase  usecase.IProductUpdateUseCase
}

func NewProductController() IProductController {
	return &productController{
		productPublishUseCase: usecase.NewProductPublishUseCase(),
		productDeleteUseCase:  usecase.NewProductDeleteUseCase(),
		productSearchUseCase:  usecase.NewProductSearchUseCase(),
		productUpdateUseCase:  usecase.NewProductUpdateUseCase(),
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
