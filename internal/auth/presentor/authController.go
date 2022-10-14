package presentor

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"waffletime/internal/auth/application/usecase"
	"waffletime/internal/auth/presentor/dto/req"
)

type Controller struct {
	userLoginUsecase usecase.UserLoginUsecase
}

func (c Controller) Login(context *gin.Context) {
	dto := req.UserLoginDto{}

	if err := context.ShouldBindJSON(&dto); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result, err := c.userLoginUsecase.Login(&dto)
	if err != nil {
		log.Fatalln("<Login>Error/LoginController", err)
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

func (c Controller) Logout(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "successfully logout",
	})
}
