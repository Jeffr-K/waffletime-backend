package presentor

import "github.com/gin-gonic/gin"

type Router struct {
	productController IProductController
}

func (r Router) Routes(route *gin.Engine) {
	controller := NewProductController()
	product := route.Group("product")
	{
		product.POST("/", controller.PublishProduct)
		product.DELETE("/", controller.DeleteProduct)
		product.PUT("/", controller.UpdateProduct)
		product.GET("/", controller.GetProductList)
	}
}
