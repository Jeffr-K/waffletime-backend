package server

import (
	"github.com/gin-gonic/gin"
	"waffletime/internal/auth"
	product "waffletime/internal/product/presentor"
	user "waffletime/internal/user/presentor"
	"waffletime/pkg/database"
	"waffletime/pkg/queue"
)

type Bootstrap struct{}

func (b Bootstrap) InitializeRouter(router *gin.Engine) {
	user := user.Router{}
	auth := auth.Router{}
	product := product.Router{}

	user.Routes(router)
	auth.Routes(router)
	product.Routes(router)
}

func (b Bootstrap) InitializeExternalPackage() {
	database.InitializedDatabase()
	queue.NewKafka()
}

func (b Bootstrap) initializeMongoDB() {}

//
