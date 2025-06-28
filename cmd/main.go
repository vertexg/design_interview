package main

import (
	"example.com/design_interview/src/core/infra/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	controllers.Setup(router)

	router.Run(":8080")
}
