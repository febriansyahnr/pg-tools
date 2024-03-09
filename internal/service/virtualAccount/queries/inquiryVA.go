package queries_virtualAccount

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	model_snapCore "github.com/febrianpaper/pg-tools/internal/model/snapCore"
	model_trxHistory "github.com/febrianpaper/pg-tools/internal/model/trxHistory"
	"github.com/febrianpaper/pg-tools/port"
	"github.com/shopspring/decimal"
)

type (
	InquiryVAHandler struct {
		snapCore port.SnapCorePort
		trxQueue port.TrxQueuePort
	}
	InquiryVAQuery struct {
		Number string
		Token  string
	}
)

func NewInquiryVAHandler(snapCore port.SnapCorePort, trxQueue port.TrxQueuePort) InquiryVAHandler {
	return InquiryVAHandler{
		snapCore: snapCore,
		trxQueue: trxQueue,
	}
}

func (h *InquiryVAHandler) InquiryVA(ctx context.Context, query InquiryVAQuery) (*model_snapCore.InquiryData, string, error) {

	refNum := time.Now().Unix()
	reqData := model_snapCore.InquiryRequestData{
		PartnerServiceId: "7001",
		CustomerNo:       "7001085933245068",
		VirtualAccountNo: query.Number,
		ChannelCode:      751,
		InquiryRequestId: fmt.Sprintf("%d", refNum),
	}

	vaData, err := h.snapCore.InquiryVA(ctx, query.Token, &reqData)
	if err != nil {
		return nil, "", err
	}
	amount := decimal.Zero

	if amt, err := decimal.NewFromString(vaData.TotalAmount.Value); err == nil {
		amount = amt
	}

	queue := model_trxHistory.NewTrxQueueVA(model_trxHistory.NewTrxQueueParams{
		Number:      query.Number,
		Amount:      amount,
		AccountName: vaData.VirtualAccountName,
		Refnum:      fmt.Sprintf("%d", refNum),
		Detail:      vaData.Json(),
	})

	if err := h.trxQueue.Create(ctx, queue); err != nil {
		slog.Error("[trxQueueAdapter]", "error", err)
	}

	return vaData, queue.ID, nil
}
