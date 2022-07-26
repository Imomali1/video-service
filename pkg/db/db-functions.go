package db

import (
	"api/entity"
	"fmt"
	"time"
)

func Create(video *entity.Video) {
	var command = `INSERT INTO videos(title, description, url, last_update) VALUES($1, $2, $3, $4);`
	_, err := MyDB.Exec(command, video.Title, video.Description, video.Url, time.Now().Format("01-02-2006"))
	check(err)
}

func Find() []entity.Video {
	var command = `SELECT url, title, description FROM videos;`
	var videos []entity.Video
	rows, err := MyDB.Query(command)
	check(err)

	for rows.Next() {
		var video entity.Video
		err = rows.Scan(&video.Url, &video.Title, &video.Description)
		check(err)
		videos = append(videos, video)
	}

	return videos
}

func FindById(id int) (entity.Video, error) {
	var command = fmt.Sprintf(`SELECT url, title, description FROM videos WHERE id=%d;`, id)
	var video entity.Video

	err := MyDB.QueryRow(command).Scan(&video.Url, &video.Title, &video.Description)

	return video, err
}

func Update(id int, updatedVideo entity.Video) (entity.Video, error) {
	_, err := FindById(id)
	if err != nil {
		return entity.Video{}, err
	}

	var command = fmt.Sprintf(`
	UPDATE videos 
	SET title = '%s', description = '%s', url = '%s' 
	WHERE id = %d RETURNING *;`, updatedVideo.Title, updatedVideo.Description, updatedVideo.Url, id)

	_, err = MyDB.Exec(command)

	return updatedVideo, err
}

func Delete(id int) (entity.Video, error) {
	video, err := FindById(id)
	if err != nil {
		return video, err
	}

	var command = fmt.Sprintf(`
	DELETE FROM videos
	WHERE id = %d;`, id)

	_, err = MyDB.Exec(command)

	return video, err
}
