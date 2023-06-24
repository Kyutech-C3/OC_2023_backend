package routers

import (
	"oc-2023/handler"

	"github.com/gin-gonic/gin"
)

func InitLikeRouter(r *gin.RouterGroup) {
	r.POST("/", handler.HandlePostLike)
	r.DELETE("/", handler.HandleDeleteLike)
}