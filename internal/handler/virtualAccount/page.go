package handler_virtualAccount

import (
	"net/http"

	"github.com/febrianpaper/pg-tools/internal/handler"
	"github.com/febrianpaper/pg-tools/view/ui"
	view_virtualAccount "github.com/febrianpaper/pg-tools/view/virtualAccount"
)

func (h *VirtualAccountHandler) LogPage(w http.ResponseWriter, r *http.Request) error {
	logVARequest, err := h.trxHistory.GetAllLogTrxVA(r.Context())
	if err != nil {
		return ErrorResponse(w, r, "log-va", err)
	}

	var logRequest []ui.TrxLog
	for _, log := range logVARequest {
		logRequest = append(logRequest, ui.TrxLog{
			ID:        log.ID,
			Type:      log.Type,
			Subtype:   log.Subtype,
			Number:    log.Number,
			CreatedAt: log.CreatedAt,
		})
	}

	return handler.Render(r, w, view_virtualAccount.LogVAPage(logRequest))
}

func (h *VirtualAccountHandler) LogDetailPage(w http.ResponseWriter, r *http.Request) error {
	id := r.URL.Query().Get("id")
	logVARequest, err := h.trxHistory.GetAllLogTrxVADetail(r.Context(), id)
	if err != nil {
		return ErrorResponse(w, r, "log-va", err)
	}

	return handler.Render(r, w, view_virtualAccount.LogDetailPage(*logVARequest))
}
