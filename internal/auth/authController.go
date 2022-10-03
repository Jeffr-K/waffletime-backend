package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct{}

func (c Controller) Login(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "successfully login.",
	})
}

func (c Controller) Logout(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "successfully logout",
	})
}
