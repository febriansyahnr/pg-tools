package model_trxHistory

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type TrxQueue struct {
	ID          string          `json:"id" db:"id"`
	Type        string          `json:"trxType" db:"trx_type"`
	Number      string          `json:"trxNo" db:"trx_no"`
	Amount      decimal.Decimal `json:"amount" db:"amount"`
	AccountName string          `json:"accountName" db:"account_name"`
	Refnum      string          `json:"refnum" db:"refnum"`
	Detail      string          `json:"detail" db:"detail"`
	InQueue     bool            `json:"inQueue" db:"in_queue"`
	CreatedAt   time.Time       `json:"createdAt" db:"created_at"`
}

type NewTrxQueueParams struct {
	Number      string          `json:"trxNo"`
	AccountName string          `json:"accountName"`
	Amount      decimal.Decimal `json:"amount"`
	Refnum      string          `json:"refnum"`
	Detail      string          `json:"detail"`
}

func NewTrxQueueVA(params NewTrxQueueParams) TrxQueue {
	return TrxQueue{
		ID:          uuid.NewString(),
		Type:        "va",
		Number:      params.Number,
		AccountName: params.AccountName,
		Amount:      params.Amount,
		Refnum:      params.Refnum,
		Detail:      params.Detail,
		InQueue:     true,
		CreatedAt:   time.Now(),
	}
}
