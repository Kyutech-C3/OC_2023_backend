package routers

import (
	"github.com/gin-gonic/gin"
)

func Router(e *gin.Engine) {
	// Routersのメインの部分
	e.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "OK",
		})
	})

	api := e.Group("/api")
	{
		// /api/v1となる
		v1 := api.Group("/v1")
		{
			like_router := v1.Group("/likes")
			InitLikeRouter(like_router)
		}
	}
}
