package presentor

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	userController Controller
}

func (r Router) Routes(route *gin.Engine) {
	user := route.Group("user")
	{
		user.POST("/", r.userController.Register)
		user.DELETE("/", r.userController.Dropdown)
		user.PUT("/", r.userController.updateUserInfo)
		user.GET("/", r.userController.getUserInfo)
	}
}
