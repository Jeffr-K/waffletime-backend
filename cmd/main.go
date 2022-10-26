package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"waffletime/cmd/server"
)

func main() {
	gin.SetMode(gin.DebugMode)

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	bootstrap := server.Bootstrap{}
	router.Use(cors.New(bootstrap.InitializeCorsConfig()))
	bootstrap.InitializeExternalPackage()
	bootstrap.InitializeRouter(router)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, world!",
		})
	})

	router.Run(":6060")
}
