package queries_trxHistory

import (
	"context"

	model_trxHistory "github.com/febrianpaper/pg-tools/internal/model/trxHistory"
	"github.com/febrianpaper/pg-tools/port"
)

type (
	GetAllLogTrxVAHandler struct {
		trxLog port.TrxLogPort
	}
)

func NewGetAllLogTrxVAHandler(trxLog port.TrxLogPort) GetAllLogTrxVAHandler {
	return GetAllLogTrxVAHandler{
		trxLog: trxLog,
	}
}

func (h *GetAllLogTrxVAHandler) GetAllLogTrxVA(ctx context.Context) ([]model_trxHistory.TrxLog, error) {
	return h.trxLog.GetAll(ctx)
}
