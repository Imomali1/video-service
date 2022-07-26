package main

import (
	"api/pkg/controller"
	"api/pkg/db"
	"api/pkg/service"
	"github.com/gin-gonic/gin"
)

var (
	videoService    = service.NewService()
	videoController = controller.NewController(videoService)
)

func main() {
	server := gin.Default()
	db.ConnDB()

	server.GET("/videos", videoController.GetAll)
	server.POST("/videos", videoController.Create)
	server.GET("/videos/:id", videoController.GetById)
	server.PUT("/videos/:id", videoController.Update)
	server.DELETE("/videos/:id", videoController.Delete)

	err := server.Run("localhost:8080")
	if err != nil {
		panic(err)
	}
	defer db.CloseDB(db.MyDB)
}
