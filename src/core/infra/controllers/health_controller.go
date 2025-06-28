package controllers

import (
	"github.com/gin-gonic/gin"
)

// Controller struct definition
type HealthController struct{}

// 新しいHealthControllerのインスタンスを生成
func newHealthController() *HealthController {
	return &HealthController{}
}

func (c *HealthController) routes(router *gin.Engine) {
	router.GET("/health", c.HealthCheck)
}
func (c *HealthController) HealthCheck(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"status": "200 OK",
	})
}
