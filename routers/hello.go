package routers

import (
	"github.com/gin-gonic/gin"
)

func InitHelloWorld(r *gin.RouterGroup) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})
}
