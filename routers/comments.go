package routers

import (
	handler "oc-2023/handlers"

	"github.com/gin-gonic/gin"
)

func InitCommentRouter(cr *gin.RouterGroup) {
	cr.POST("", handler.HandlePostComment)
	cr.DELETE("", handler.HandleDeleteComment)
}
