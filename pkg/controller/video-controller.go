package controller

import (
	"api/entity"
	"api/pkg/service"
	"api/pkg/utils/convert"
	"github.com/gin-gonic/gin"
	"net/http"
)

type VideoController interface {
	Create(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetById(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Update(ctx *gin.Context)
}

type videoController struct {
	service service.VideoService
}

func NewController(service service.VideoService) VideoController {
	return &videoController{
		service: service,
	}
}

func (controller *videoController) Create(ctx *gin.Context) {
	var newVideo entity.Video
	err := ctx.ShouldBindJSON(&newVideo)
	CatchError(ctx, err)
	video := controller.service.Create(newVideo)
	ctx.JSON(http.StatusCreated, gin.H{"Created video": video})
}

func (controller *videoController) GetAll(ctx *gin.Context) {
	videos, err := controller.service.GetAll()
	CatchError(ctx, err)
	ctx.JSON(http.StatusOK, gin.H{"Videos": videos})
}

func (controller *videoController) GetById(ctx *gin.Context) {
	id := convert.MakeInt(ctx.Param("id"))
	video, err := controller.service.GetById(id)
	CatchError(ctx, err)
	ctx.JSON(http.StatusOK, gin.H{"searched video": video})
}

func (controller *videoController) Delete(ctx *gin.Context) {
	id := convert.MakeInt(ctx.Param("id"))
	video, err := controller.service.DeleteById(id)
	CatchError(ctx, err)
	ctx.JSON(http.StatusOK, gin.H{"deleted video": video})
}

func (controller *videoController) Update(ctx *gin.Context) {
	var updatedVideo entity.Video
	err := ctx.ShouldBindJSON(&updatedVideo)
	CatchError(ctx, err)
	id := convert.MakeInt(ctx.Param("id"))
	video, err := controller.service.UpdateById(id, updatedVideo)
	CatchError(ctx, err)
	ctx.JSON(http.StatusOK, gin.H{"updated video": video})
}

func CatchError(ctx *gin.Context, err error) {
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}
