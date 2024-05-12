package handler_virtualAccount

import (
	"fmt"
	"net/http"
	"time"

	"github.com/febrianpaper/pg-tools/internal/handler"
	service_trxHistory "github.com/febrianpaper/pg-tools/internal/service/trxHistory"
	service_virtualAccount "github.com/febrianpaper/pg-tools/internal/service/virtualAccount"
	"github.com/febrianpaper/pg-tools/view/ui"
	view_virtualAccount "github.com/febrianpaper/pg-tools/view/virtualAccount"
)

type VirtualAccountHandler struct {
	vaService  service_virtualAccount.IVirtualAccountService
	trxHistory service_trxHistory.ITrxHistory
}

func New(vaService service_virtualAccount.IVirtualAccountService, trxHistory service_trxHistory.ITrxHistory) *VirtualAccountHandler {
	return &VirtualAccountHandler{
		vaService:  vaService,
		trxHistory: trxHistory,
	}
}

func (h *VirtualAccountHandler) Index(w http.ResponseWriter, r *http.Request) error {
	listVA := []ui.VAItemData{}
	queueVA, err := h.vaService.GetAllQueue(r.Context())
	if err == nil {
		for _, va := range queueVA {
			vaType := "open"
			if va.Number[:4] == "7663" {
				vaType = "close"
			}
			listVA = append(listVA, ui.VAItemData{
				ID:     va.ID,
				Number: va.Number,
				Name:   va.AccountName,
				Amount: va.Amount.InexactFloat64(),
				Type:   vaType,
			})
		}
	}

	return handler.Render(r, w, view_virtualAccount.Index(listVA))
}

func (h *VirtualAccountHandler) ListProcessedVA(w http.ResponseWriter, r *http.Request) error {
	if err := h.vaService.ResetQueues(r.Context()); err != nil {
		now := time.Now().Unix()
		return ErrorResponse(w, r, fmt.Sprintf("%d", now), err)
	}
	return handler.Render(r, w, view_virtualAccount.ListVA([]ui.VAItemData{}))
}
