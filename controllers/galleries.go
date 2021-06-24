package controllers

import (
	"github.com/yakushou730/go-web-development/models"
	"github.com/yakushou730/go-web-development/views"
)

type Galleries struct {
	New *views.View
	gs  models.GalleryService
}

func NewGalleries(gs models.GalleryService) *Galleries {
	return &Galleries{
		New: views.NewView("bootstrap", "galleries/new"),
		gs:  gs,
	}
}
