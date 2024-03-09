package queries_virtualAccount

import (
	"context"

	model_trxHistory "github.com/febrianpaper/pg-tools/internal/model/trxHistory"
	"github.com/febrianpaper/pg-tools/port"
)

type (
	GetAllQueueHandler struct {
		trxQueue port.TrxQueuePort
	}
)

func NewGetAllQueueHandler(trxQueue port.TrxQueuePort) GetAllQueueHandler {
	return GetAllQueueHandler{
		trxQueue: trxQueue,
	}
}

func (h *GetAllQueueHandler) GetAllQueue(ctx context.Context) ([]model_trxHistory.TrxQueue, error) {
	return h.trxQueue.GetAll(ctx)
}
