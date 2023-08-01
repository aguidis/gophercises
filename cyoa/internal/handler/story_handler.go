package handler

import (
	"github.com/aguidis/cyoa/internal/layout"
	"github.com/aguidis/cyoa/internal/model"
	"net/http"
)

func NewHandler(s model.Story) http.Handler {
	return handler{s}
}

type handler struct {
	s model.Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := layout.Tpl.Execute(w, h.s["intro"])
	if err != nil {
		panic(err)
	}
}
