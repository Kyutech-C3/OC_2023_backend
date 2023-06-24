package main

import (
	// "log"
	// "oc-2023/config"
	// "oc-2023/db"
	// "oc-2023/routers"

	// "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// gin.Contextは、リクエストの (GETされた, PUTした) データやパラメータ、またはエラー情報などいろんなものを含んだもの
	r.GET("/", func(c *gin.Context) {
		
		// gin.Hはmap[string]interface{}と同義。
		c.JSON(200, gin.H{
			"message":"Hello World",
		})

	})
	r.Run()

	// r := gin.New()
	// cc := cors.DefaultConfig()
	// cc.AllowAllOrigins = true
	// cc.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	// r.Use(cors.New(cc))
	// config.LoadConfig()
	// db.InitDB()
	// routers.Router(r)

	// if err := r.Run(":8000"); err != nil {
	// 	log.Fatal("failed run app: ", err)
	// }
}
