package mysql_trxHistory

import (
	"context"
	"errors"
	"log/slog"

	model_trxHistory "github.com/febrianpaper/pg-tools/internal/model/trxHistory"
	"github.com/febrianpaper/pg-tools/pkg/mySqlExt"
	"github.com/febrianpaper/pg-tools/port"
)

type TrxQueueRepo struct {
	db mySqlExt.IMySqlExt
}

// FindByID implements port.TrxQueuePort.
func (t *TrxQueueRepo) FindByID(ctx context.Context, id string) (*model_trxHistory.TrxQueue, error) {
	var queue model_trxHistory.TrxQueue
	query := `
		SELECT id, trx_type, trx_no, amount, account_name, refnum, detail, in_queue, created_at
		FROM trx_queue
		WHERE id = ?
	`
	if err := t.db.GetContext(ctx, &queue, query, id); err != nil {
		return nil, err
	}

	return &queue, nil
}

// Create implements port.TrxQueuePort.
func (t *TrxQueueRepo) Create(ctx context.Context, queue model_trxHistory.TrxQueue) error {
	query := `
	INSERT INTO trx_queue
	(id, trx_type, trx_no, amount, account_name, refnum, detail, in_queue, created_at)
	VALUES(:id, :trx_type,:trx_no, :amount, :account_name,:refnum, :detail, 1,:created_at)
	`

	affected, err := t.db.NamedExecContext(ctx, query, queue)
	if err != nil {
		return err
	}
	if !affected {
		return errors.New("trx_queue insert not affected")
	}
	return nil
}

// GetAll implements port.TrxQueuePort.
func (t *TrxQueueRepo) GetAll(ctx context.Context) ([]model_trxHistory.TrxQueue, error) {
	query := `
	SELECT id, trx_type, trx_no, amount, account_name, refnum, detail, in_queue, created_at
	FROM trx_queue
	WHERE in_queue = 1
	`
	var queueList []model_trxHistory.TrxQueue
	if err := t.db.SelectContext(ctx, &queueList, query); err != nil {
		slog.Error("GetAll", "error", err.Error())
		return queueList, err
	}

	return queueList, nil
}

// Reset implements port.TrxQueuePort.
func (t *TrxQueueRepo) Reset(ctx context.Context) error {
	query := `UPDATE trx_queue SET in_queue = 0 WHERE in_queue = 1`
	_, err := t.db.ExecContext(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func (t *TrxQueueRepo) Acknowledge(ctx context.Context, id string) error {
	query := `UPDATE trx_queue SET in_queue = 0 WHERE id = ?`
	_, err := t.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func NewTrxQueueRepo(db mySqlExt.IMySqlExt) *TrxQueueRepo {
	return &TrxQueueRepo{db: db}
}

var _ port.TrxQueuePort = (*TrxQueueRepo)(nil)
