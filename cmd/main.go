package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"waffletime/cmd/server"
)

func main() {
	router := gin.Default()

	bootstrap := server.Bootstrap{} //
	bootstrap.InitializeExternalPackage()
	bootstrap.InitializeRouter(router)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, world!",
		})
	})

	router.Run(":6060")
}
