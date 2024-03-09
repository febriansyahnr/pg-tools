package service_trxHistory

import (
	"context"

	model_trxHistory "github.com/febrianpaper/pg-tools/internal/model/trxHistory"
	queries_trxHistory "github.com/febrianpaper/pg-tools/internal/service/trxHistory/queries"
	"github.com/febrianpaper/pg-tools/port"
)

type (
	ITrxHistory interface {
		Commands
		Queries
	}
	Commands interface{}
	Queries  interface {
		GetAllLogTrxVA(ctx context.Context) ([]model_trxHistory.TrxLog, error)
		GetAllLogTrxVADetail(ctx context.Context, id string) (*model_trxHistory.TrxLog, error)
	}
	appCommands struct{}
	appQueries  struct {
		queries_trxHistory.GetAllLogTrxVAHandler
		queries_trxHistory.GetAllLogTrxVADetailHandler
	}
	TrxHistory struct {
		trxHistoryAdapter port.TrxLogPort
		appCommands
		appQueries
	}
)

var _ ITrxHistory = (*TrxHistory)(nil)

func New(trxHistoryAdapter port.TrxLogPort) *TrxHistory {
	return &TrxHistory{
		trxHistoryAdapter: trxHistoryAdapter,
		appCommands:       appCommands{},
		appQueries: appQueries{
			GetAllLogTrxVAHandler:       queries_trxHistory.NewGetAllLogTrxVAHandler(trxHistoryAdapter),
			GetAllLogTrxVADetailHandler: queries_trxHistory.NewGetAllLogTrxVADetailHandler(trxHistoryAdapter),
		},
	}
}
