package mysql_trxHistory

import (
	"context"

	model_trxHistory "github.com/febrianpaper/pg-tools/internal/model/trxHistory"
	"github.com/febrianpaper/pg-tools/pkg/mySqlExt"
	"github.com/febrianpaper/pg-tools/port"
)

type TrxLogRepo struct {
	db mySqlExt.IMySqlExt
}

// GetByID implements port.TrxLogPort.
func (t *TrxLogRepo) GetByID(ctx context.Context, id string) (*model_trxHistory.TrxLog, error) {
	query := `
	SELECT 
		id, trx_type, sub_type, trx_no, additional, request, response, status, created_at 
	FROM trx_log 
	WHERE id = ?
	LIMIT 1`
	var log model_trxHistory.TrxLog

	if err := t.db.GetContext(ctx, &log, query, id); err != nil {
		return nil, err
	}
	return &log, nil
}

// Create implements port.TrxLogPort.
func (t *TrxLogRepo) Create(ctx context.Context, log model_trxHistory.TrxLog) error {
	query := `
	INSERT INTO trx_log
	(id, trx_type, sub_type, trx_no, additional, request, response, status, created_at)
	VALUES(:id, :trx_type, :sub_type, :trx_no, :additional, :request, :response, :status, :created_at)
	`
	_, err := t.db.NamedExecContext(ctx, query, log)
	if err != nil {
		return err
	}
	return nil
}

// GetAll implements port.TrxLogPort.
func (t *TrxLogRepo) GetAll(ctx context.Context) ([]model_trxHistory.TrxLog, error) {
	query := `
	SELECT id, trx_type, sub_type, trx_no, additional, request, response, status, created_at
	FROM trx_log
	ORDER BY created_at DESC
	`
	var logList []model_trxHistory.TrxLog
	if err := t.db.SelectContext(ctx, &logList, query); err != nil {
		return logList, err
	}
	return logList, nil
}

func NewTrxLogRepo(db mySqlExt.IMySqlExt) *TrxLogRepo {
	return &TrxLogRepo{db: db}
}

var _ port.TrxLogPort = (*TrxLogRepo)(nil)
