package presentor

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jeffr-k/waffletime/internal/user/presentor/dto"
	"github.com/jeffr-k/waffletime/internal/user/usecase"
	"net/http"
)

type Controller struct {
	userCreateUseCase usecase.UserCreateUseCase
}

func (c Controller) Register(context *gin.Context) {
	var dto dto.CreateUserDto
	fmt.Println(dto)
	if err := context.ShouldBindJSON(&dto); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	usecase, err := c.userCreateUseCase.UserRegisterAsMember(&dto)
	if err != nil {
		fmt.Println(err)
	}

	context.JSON(http.StatusOK, gin.H{"result": usecase})
}

func (c Controller) Dropdown(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "successfully login.",
	})
}

func (c Controller) updateUserInfo(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "successfully login.",
	})
}

func (c Controller) getUserInfo(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "hi.",
	})
}
