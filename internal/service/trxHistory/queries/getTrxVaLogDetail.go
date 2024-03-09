package queries_trxHistory

import (
	"context"

	model_trxHistory "github.com/febrianpaper/pg-tools/internal/model/trxHistory"
	"github.com/febrianpaper/pg-tools/port"
)

type (
	GetAllLogTrxVADetailHandler struct {
		trxLog port.TrxLogPort
	}
)

func NewGetAllLogTrxVADetailHandler(trxLog port.TrxLogPort) GetAllLogTrxVADetailHandler {
	return GetAllLogTrxVADetailHandler{
		trxLog: trxLog,
	}
}

func (h *GetAllLogTrxVADetailHandler) GetAllLogTrxVADetail(ctx context.Context, id string) (*model_trxHistory.TrxLog, error) {
	return h.trxLog.GetByID(ctx, id)
}
