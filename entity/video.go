package entity

type Video struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Url         string `json:"url" binding:"required"`
}
