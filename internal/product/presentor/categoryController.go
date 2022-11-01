package presentor

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"waffletime/internal/product/application/usecase/product"
)

type ICategory interface {
	CreateCategory(context *gin.Context)
	DeleteCategory(context *gin.Context)
	UpdateCategory(context *gin.Context)
	GetCategory(context *gin.Context)
	GetCategoryList(context *gin.Context)
}

type categoryController struct {
	productPublishUseCase product.IProductPublishUseCase
	productDeleteUseCase  product.IProductDeleteUseCase
	productSearchUseCase  product.IProductSearchUseCase
	productUpdateUseCase  product.IProductUpdateUseCase
}

func NewProductController() ICategory {
	return &categoryController{
		categoryCreateUseCase: product.NewProductPublishUseCase(),
		categoryDeleteUseCase: product.NewProductDeleteUseCase(),
		categorySearchUseCase: product.NewProductSearchUseCase(),
		categoryUpdateUseCase: product.NewProductUpdateUseCase(),
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
