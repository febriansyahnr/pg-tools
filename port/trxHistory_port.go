package port

import (
	"context"

	model_trxHistory "github.com/febrianpaper/pg-tools/internal/model/trxHistory"
)

type TrxQueuePort interface {
	Create(ctx context.Context, queue model_trxHistory.TrxQueue) error
	GetAll(ctx context.Context) ([]model_trxHistory.TrxQueue, error)
	Reset(ctx context.Context) error
	Acknowledge(ctx context.Context, id string) error
	FindByID(ctx context.Context, id string) (*model_trxHistory.TrxQueue, error)
}

type TrxLogPort interface {
	Create(ctx context.Context, log model_trxHistory.TrxLog) error
	GetAll(ctx context.Context) ([]model_trxHistory.TrxLog, error)
	GetByID(ctx context.Context, id string) (*model_trxHistory.TrxLog, error)
}
