package model_trxHistory

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

const (
	TrxLogStatusFailed  = 0
	TrxLogStatusSuccess = 1
)

type TrxLog struct {
	ID         string    `json:"id" db:"id"`
	Type       string    `json:"trxType" db:"trx_type"`
	Subtype    string    `json:"subtype" db:"sub_type"`
	Number     string    `json:"trxNo" db:"trx_no"`
	Additional string    `json:"additional" db:"additional"`
	Request    string    `json:"request" db:"request"`
	Response   string    `json:"response" db:"response"`
	Status     int       `json:"status" db:"status"`
	CreatedAt  time.Time `json:"createdAt" db:"created_at"`
}

type NewTrxLogParams struct {
	Number     string            `json:"trxNo" db:"trx_no"`
	Additional map[string]string `json:"additional" db:"additional"`
	Request    string            `json:"request" db:"request"`
	Response   string            `json:"response" db:"response"`
	Status     int               `json:"status" db:"status"`
}

func NewTrxLogVA(subtype string, params NewTrxLogParams) TrxLog {
	additional := ""

	if additionalByte, err := json.Marshal(params.Additional); err != nil {
		additional = ""
	} else {
		additional = string(additionalByte)
	}

	return TrxLog{
		ID:         uuid.NewString(),
		Type:       "va",
		Subtype:    subtype,
		Number:     params.Number,
		Additional: additional,
		Request:    params.Request,
		Response:   params.Response,
		Status:     params.Status,
		CreatedAt:  time.Now(),
	}
}
