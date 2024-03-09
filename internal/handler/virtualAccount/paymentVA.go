package handler_virtualAccount

import (
	"net/http"

	"github.com/febrianpaper/pg-tools/internal/handler"
	commands_virtualAccount "github.com/febrianpaper/pg-tools/internal/service/virtualAccount/commands"
	"github.com/febrianpaper/pg-tools/view/ui"
	view_virtualAccount "github.com/febrianpaper/pg-tools/view/virtualAccount"
	"github.com/shopspring/decimal"
)

func (h *VirtualAccountHandler) Payment(w http.ResponseWriter, r *http.Request) error {
	amount := r.FormValue("amount")
	id := r.FormValue("id")

	amt, err := decimal.NewFromString(amount)
	if err != nil {
		return ErrorResponse(w, r, "0", err)
	}

	token, err := h.vaService.GetToken(r.Context())
	if err != nil {
		return ErrorResponse(w, r, "0", err)
	}

	if err := h.vaService.PaymentVA(r.Context(), commands_virtualAccount.PaymentVACommands{
		ID:     id,
		Amount: amt,
		Token:  token,
	}); err != nil {
		return ErrorResponse(w, r, "0", err)
	}

	listVA := []ui.VAItemData{}
	queueVA, err := h.vaService.GetAllQueue(r.Context())
	if err == nil {
		for _, va := range queueVA {
			vaType := "close"
			if va.Number[:4] == "7664" {
				vaType = "open"
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

	return handler.Render(r, w, view_virtualAccount.ListVA(listVA))
}
