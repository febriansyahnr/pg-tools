package handler_virtualAccount

import (
	"net/http"

	"github.com/febrianpaper/pg-tools/internal/handler"
	"github.com/febrianpaper/pg-tools/view/ui"
)

func ErrorResponse(w http.ResponseWriter, r *http.Request, id string, err error) error {
	w.Header().Set("HX-Reswap", "innerHTML transition:true")
	w.Header().Set("HX-Retarget", "#error-wrapper")
	return handler.Render(r, w, ui.ErrorBox(id, err.Error()))
}
