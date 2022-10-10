package presentor

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"waffletime/internal/user/application/usecase"
	"waffletime/internal/user/presentor/dto"
)

type Controller struct {
	userCreateUseCase usecase.UserCreateUseCase
	userDeleteUseCase usecase.UserDeleteUseCase
	userUpdateUseCase usecase.UserUpdateUseCase
	userSearchUseCase usecase.UserSearchUseCase
}

func (c Controller) Register(context *gin.Context) {
	dto := dto.CreateUserDto{}

	if err := context.ShouldBindJSON(&dto); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result := c.userCreateUseCase.UserRegisterAsMember(&dto)
	if result != nil {
		fmt.Println(result)
	}

	context.JSON(http.StatusOK, gin.H{"result": "successfully saved."})
}

func (c Controller) Dropdown(context *gin.Context) {
	userId := context.DefaultQuery("id", "")
	cast, _ := strconv.Atoi(userId)
	result := c.userDeleteUseCase.UserDropdownMembership(cast)
	if result != nil {
		fmt.Println(result)
	}

	context.JSON(http.StatusOK, gin.H{"message": "successfully deleted."})
}

func (c Controller) updateUserInfo(context *gin.Context) {
	dto := dto.UpdateUserDto{}

	if err := context.ShouldBindJSON(&dto); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result := c.userUpdateUseCase.UpdateUserInfo(&dto)
	if result != nil {
		fmt.Println(result)
		context.JSON(http.StatusOK, gin.H{"result": "successfully saved."})
	}

	context.JSON(http.StatusOK, gin.H{"result": result})
}

func (c Controller) getUserInfo(context *gin.Context) {
	userId := context.DefaultQuery("id", "")
	cast, _ := strconv.Atoi(userId)

	result, err := c.userSearchUseCase.GetUserById(cast)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"message": "Notfound."})
	}

	context.JSON(http.StatusOK, gin.H{"result": result})
}
