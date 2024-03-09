package handler

import (
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
)

func MakeHandler(h func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			slog.Error("Error handling request", "error", err, "path", r.URL.Path)
		}
	}
}

func Render(r *http.Request, w http.ResponseWriter, component templ.Component) error {
	return component.Render(r.Context(), w)
}
