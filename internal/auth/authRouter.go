package auth

import (
	"github.com/gin-gonic/gin"
	"waffletime/internal/auth/presentor"
)

type Router struct {
	authController presentor.Controller
}

func (ar Router) Routes(route *gin.Engine) {
	auth := route.Group("auth")
	{
		auth.POST("/login", ar.authController.Login)
		auth.DELETE("/logout", ar.authController.Logout)
	}
}
