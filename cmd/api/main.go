package main

import (
	"log"
	"oc-2023/pkg/config"
	commentrepository "oc-2023/pkg/infra/repository/comment"
	likerepository "oc-2023/pkg/infra/repository/like"
	commenthandler "oc-2023/pkg/interface/handler/comment"
	likehandler "oc-2023/pkg/interface/handler/like"
	workhandler "oc-2023/pkg/interface/handler/work"
	"oc-2023/pkg/usecase"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	conn := config.ConnectDB()

	//DI
	commentRepository := commentrepository.New(conn)
	likeRepository := likerepository.New(conn)
	interactor := usecase.NewInteractor(commentRepository, likeRepository)
	commentHandler := commenthandler.New(interactor)
	likeHandler := likehandler.New(interactor)
	workHandler := workhandler.New(interactor)

	//gin初期化,cors設定
	r := gin.Default()
	r.Use(cors.Default())

	//ルーティング設定
	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			lr := v1.Group("/likes")
			lr.POST("", likeHandler.Create)
			lr.DELETE("", likeHandler.Delete)

			cr := v1.Group("/comments")
			cr.POST("", commentHandler.Create)
			cr.DELETE("", commentHandler.Delete)

			wr := v1.Group("/works")
			wr.GET("", workHandler.Index)
			wr.GET("/:id", workHandler.GetByID)
		}
	}
	if err := r.Run(":8000"); err != nil {
		log.Fatal("failed run app: ", err)
	}
}
