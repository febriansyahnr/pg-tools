package commands_virtualAccount

import (
	"context"
	"errors"
	"log/slog"

	model_snapCore "github.com/febrianpaper/pg-tools/internal/model/snapCore"
	"github.com/febrianpaper/pg-tools/port"
	"github.com/shopspring/decimal"
)

type (
	PaymentVAHandler struct {
		snapCore port.SnapCorePort
		trxQueue port.TrxQueuePort
	}
	PaymentVACommands struct {
		ID     string
		Token  string
		Amount decimal.Decimal
	}
)

func NewPaymentVAHandler(snapCore port.SnapCorePort, trxQueue port.TrxQueuePort) PaymentVAHandler {
	return PaymentVAHandler{
		snapCore: snapCore,
		trxQueue: trxQueue,
	}
}

func (h *PaymentVAHandler) PaymentVA(ctx context.Context, command PaymentVACommands) error {
	queue, err := h.trxQueue.FindByID(ctx, command.ID)
	if err != nil {
		return err
	}

	if (queue.Number[:4] == "7663") && !queue.Amount.Equal(command.Amount) {
		slog.Info("invalid amount", "expected", command.Amount, "actual", queue.Amount)
		return errors.New("invalid amount")
	}

	inqData, err := model_snapCore.NewInquiryDataFromString(queue.Detail)
	if err != nil {
		return err
	}
	payRequest := model_snapCore.NewVAPaymentRequestFromInquiryData(inqData)
	payRequest.PaidAmount = model_snapCore.Amount{
		Value:    command.Amount.String(),
		Currency: "IDR",
	}

	if err := h.snapCore.PaymentVA(ctx, command.Token, &payRequest); err != nil {
		return err
	}

	if err := h.trxQueue.Acknowledge(ctx, command.ID); err != nil {
		return err
	}

	return nil
}
