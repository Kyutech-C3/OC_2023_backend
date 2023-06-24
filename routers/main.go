package routers

import (
	"github.com/gin-gonic/gin"
)

func Router(e *gin.Engine) {
	// Routersのメインの部分
	api := e.Group("/api")
	{
		// /api/v1となる
		v1 := api.Group("/v1")
		{
			// InitHelloWorld(v1)
			like_router := v1.Group("/likes")
			InitLikeRouter(like_router)
		}
	}
}
