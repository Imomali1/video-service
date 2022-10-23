package main

import (
	"api/pkg/controller"
	"api/pkg/db"
	"api/pkg/service"
	"api/pkg/utils/reader"
	"github.com/gin-gonic/gin"
)

var (
	videoService    = service.NewService()
	videoController = controller.NewController(videoService)
)

func init() {
	db.ConnDB()

	initTable := reader.ReadFile("pkg/db/SQL/table.sql")

	_, err := db.MyDB.Exec(string(initTable))
	if err != nil {
		panic(err)
	}

	defer db.CloseDB(db.MyDB)
}

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
