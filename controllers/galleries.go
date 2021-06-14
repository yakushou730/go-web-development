package controllers

import (
	"net/http"

	"github.com/yakushou730/go-web-development/views"
)

func NewGalleries() *Galleries {
	return &Galleries{
		NewView: views.NewView("bootstrap", "galleries/new"),
	}
}

type Galleries struct {
	NewView *views.View
}

func (g *Galleries) New(w http.ResponseWriter, r *http.Request) {
	if err := g.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}
