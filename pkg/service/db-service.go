package service

import (
	"api/entity"
	"api/pkg/db"
	"errors"
)

type VideoService interface {
	Create(entity.Video) entity.Video
	GetAll() ([]entity.Video, error)
	GetById(int) (entity.Video, error)
	DeleteById(int) (entity.Video, error)
	UpdateById(int, entity.Video) (entity.Video, error)
}

type videoService struct {
	videos []entity.Video
}

func NewService() VideoService {
	return &videoService{}
}

func (service *videoService) Create(video entity.Video) entity.Video {
	db.Create(&video)
	service.videos = append(service.videos, video)
	return video
}

func (service *videoService) GetAll() ([]entity.Video, error) {
	videos := db.Find()
	if videos == nil {
		err := errors.New("there are not videos saved")
		return nil, err
	}
	return videos, nil
}

func (service *videoService) GetById(id int) (entity.Video, error) {
	return db.FindById(id)
}

func (service *videoService) DeleteById(id int) (entity.Video, error) {
	return db.Delete(id)
}

func (service *videoService) UpdateById(id int, updatedVideo entity.Video) (entity.Video, error) {
	return db.Update(id, updatedVideo)
}
