package commands_virtualAccount

import (
	"context"

	"github.com/febrianpaper/pg-tools/port"
)

type (
	ResetQueuesHandler struct {
		trxQueue port.TrxQueuePort
	}
)

func NewResetQueuesHandler(trxQueue port.TrxQueuePort) ResetQueuesHandler {
	return ResetQueuesHandler{
		trxQueue: trxQueue,
	}
}

func (h *ResetQueuesHandler) ResetQueues(ctx context.Context) error {
	return h.trxQueue.Reset(ctx)
}
