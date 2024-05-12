package handler_virtualAccount

import (
	"errors"
	"net/http"

	"github.com/febrianpaper/pg-tools/internal/handler"
	queries_virtualAccount "github.com/febrianpaper/pg-tools/internal/service/virtualAccount/queries"
	"github.com/febrianpaper/pg-tools/view/ui"
	"github.com/shopspring/decimal"
)

func (h *VirtualAccountHandler) Inquiry(w http.ResponseWriter, r *http.Request) error {
	vaNumber := r.FormValue("va_number")

	if len(vaNumber) > 16 {
		return ErrorResponse(w, r, vaNumber, errors.New("virtual account number length must not greather then 16 digits"))
	}

	token, err := h.vaService.GetToken(r.Context())
	if err != nil {
		return ErrorResponse(w, r, vaNumber, err)
	}

	inquiryData, queueID, err := h.vaService.InquiryVA(r.Context(), queries_virtualAccount.InquiryVAQuery{
		Number: vaNumber,
		Token:  token,
	})

	if err != nil {
		return ErrorResponse(w, r, vaNumber, err)
	}

	totalAmount, err := decimal.NewFromString(inquiryData.TotalAmount.Value)
	if err != nil {
		return ErrorResponse(w, r, vaNumber, err)
	}
	vaType := "open"
	if vaNumber[:4] == "7663" {
		vaType = "closed"
	}

	vaViewData := ui.VAItemData{
		ID:     queueID,
		Number: vaNumber,
		Name:   inquiryData.VirtualAccountName,
		Amount: totalAmount.InexactFloat64(),
		Type:   vaType,
	}

	return handler.Render(r, w, ui.VAItem(vaViewData))
}
