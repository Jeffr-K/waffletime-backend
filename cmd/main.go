package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	bootstrap := Bootstrap{}
	bootstrap.initializedDatabase()
	bootstrap.initializeKafka()
	bootstrap.initializeRouter(router)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, world!",
		})
	}) //

	router.Run(":6060")
}
