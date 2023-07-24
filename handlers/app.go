package handlers

import (
	"github.com/ftsog/goweb/database"
)

type ModelInterface interface {
	CreatePost(title, body string) error
	ReadPost(postID int64) (*database.Post, error)
	DeletePost(postID int64) error
}

type Application struct {
	Model ModelInterface
}
