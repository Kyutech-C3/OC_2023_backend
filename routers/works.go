package routers

import (
	handler "oc-2023/handlers"

	"github.com/gin-gonic/gin"
)

func InitWorkRouter(lr *gin.RouterGroup) {
	lr.GET("", handler.HandleGetWorks)
	lr.GET("/:id", handler.HandleGetWork)
}
