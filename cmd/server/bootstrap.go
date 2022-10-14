package server

import (
	"github.com/gin-gonic/gin"
	"waffletime/internal/auth/presentor"
	product "waffletime/internal/product/presentor"
	user "waffletime/internal/user/presentor"
	"waffletime/pkg/database"
	"waffletime/pkg/queue"
	"waffletime/pkg/redis"
)

type Bootstrap struct{}

func (b Bootstrap) InitializeRouter(router *gin.Engine) {
	user := user.Router{}
	auth := presentor.Router{}
	product := product.Router{}

	user.Routes(router)
	auth.Routes(router)
	product.Routes(router)
}

func (b Bootstrap) InitializeExternalPackage() {
	database.InitializedDatabase()
	redis.InitializeRedis()
	queue.InitializeKafka()
}

func (b Bootstrap) initializeMongoDB() {}
