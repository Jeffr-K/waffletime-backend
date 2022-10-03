package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	bootstrap := Bootstrap{}
	bootstrap.initializedDatabase() //
	bootstrap.initializeRouter(router)

	router.Run(":8080")
}
