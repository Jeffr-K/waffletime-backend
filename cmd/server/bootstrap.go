package server

import (
	"github.com/gin-contrib/cors"
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

func (b Bootstrap) InitializeCorsConfig() cors.Config {
	return cors.Config{
		AllowAllOrigins:        true,
		AllowOrigins:           []string{"http://localhost:6060", "http://127.0.0.1:6060"},
		AllowOriginFunc:        nil,
		AllowMethods:           []string{"POST", "DELETE", "PUT", "PATCH", "GET"},
		AllowHeaders:           []string{"Origin", "Accept", "Content-Type", "X-Requested-With", "Multipart-FormData"},
		AllowCredentials:       false,
		ExposeHeaders:          nil,
		MaxAge:                 0,
		AllowWildcard:          false,
		AllowBrowserExtensions: false,
		AllowWebSockets:        false,
		AllowFiles:             false,
	}
}
