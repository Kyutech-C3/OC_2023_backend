package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(e *gin.Engine) {
	api := e.Group("/api")
	{
		// /api/v1となる
		v1 := api.Group("/v1")
		{
			lr := v1.Group("/likes")
			InitLikeRouter(lr)

			cr := v1.Group("/comments")
			InitCommentRouter(cr)

			wr := v1.Group("/works")
			InitWorkRouter(wr)
		}
	}
}
