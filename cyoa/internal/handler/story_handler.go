package handler

import (
	"github.com/aguidis/cyoa/internal/layout"
	"github.com/aguidis/cyoa/internal/model"
	"html/template"
	"log"
	"net/http"
	"strings"
)

type StoryHandlerOption func(h *handler)

func WithTemplate(t *template.Template) StoryHandlerOption {
	return func(h *handler) {
		h.t = t
	}
}

func WithPathFunc(fn func(r *http.Request) string) StoryHandlerOption {
	return func(h *handler) {
		h.pathFn = fn
	}
}

func NewHandler(s model.Story, opts ...StoryHandlerOption) http.Handler {
	h := handler{s, layout.Tpl, defaultPathFn}
	for _, opt := range opts {
		opt(&h)
	}

	return h
}

type handler struct {
	s      model.Story
	t      *template.Template
	pathFn func(r *http.Request) string
}

func defaultPathFn(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)

	if path == "" || path == "/" {
		path = "/intro"
	}

	return path[1:]
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defaultPath := h.pathFn(r)

	if chapter, ok := h.s[defaultPath]; ok {
		err := h.t.Execute(w, chapter)
		if err != nil {
			log.Printf("%v", err)
			http.Error(w, "Something went wrong...", http.StatusInternalServerError)
		}
		return
	}

	http.Error(w, "Chapter no found.", http.StatusNotFound)
}

// PathFn Updated chapter parsing function. Technically you don't
// *have* to get the story from the path (it could be a
// header or anything else) but I'm not going to rename this
// since "path" is what we used in the videos.
func PathFn(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	if path == "/story" || path == "/story/" {
		path = "/story/intro"
	}
	return path[len("/story/"):]
}
