package main

import (
	"log"

	"oc-2023/config"
	"oc-2023/db"
	"oc-2023/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	db.InitDB()
	r := gin.Default()

	routers.Router(r)
	r.Run()
	if err := r.Run(":8000"); err != nil {
		log.Fatal("failed run app: ", err)
	}
}
