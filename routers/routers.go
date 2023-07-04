package routers

import (
	"github.com/gin-gonic/gin"
)

func Router(e *gin.Engine) {
	api := e.Group("/api")
	{
		// /api/v1となる
		v1 := api.Group("/v1")
		{
			cr := v1.Group("/comments")
			InitCommentRouter(cr)
		}
	}
}
