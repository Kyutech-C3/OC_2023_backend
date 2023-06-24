package routers

import (
	"oc-2023/handlers"

	"github.com/gin-gonic/gin"
)

func InitLikeRouter(r *gin.RouterGroup) {
	r.POST("/", handlers.HandlePostLike)
	r.DELETE("/", handlers.HandleDeleteLike)
}