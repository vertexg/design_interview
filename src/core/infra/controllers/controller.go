package controllers

import (
	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {
	healthController := newHealthController()
	healthController.routes(router)

}
