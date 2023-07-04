package routers

import (
	handler "oc-2023/handlers"

	"github.com/gin-gonic/gin"
)

func InitLikeRouter(lr *gin.RouterGroup) {
	lr.POST("", handler.HandlePostLike)
	lr.DELETE("", handler.HandleDeleteLike)
}
