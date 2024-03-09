package model_snapCore

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

type SnapVAResponse struct {
	ResponseCode       string         `json:"responseCode"`
	ResponseMessage    string         `json:"responseMessage"`
	VirtualAccountData any            `json:"virtualAccountData,omitempty"`
	AdditionalInfo     map[string]any `json:"additionalInfo,omitempty"`
}

type Description struct {
	English   string `json:"english"`
	Indonesia string `json:"indonesia"`
}

type Amount struct {
	Value    string `json:"value"`
	Currency string `json:"currency"`
}

func (a *Amount) String() string {
	return fmt.Sprintf(
		`{"value":"%v","currency":"%s"}`,
		a.Value,
		a.Currency,
	)
}

func NewAmountFromSqlNullString(amount sql.NullString) Amount {
	var result Amount
	result.Currency = "IDR"
	result.Value = "0.00"
	if !amount.Valid {
		return result
	}
	if err := json.Unmarshal([]byte(amount.String), &result); err != nil {
		return result
	}
	return result
}

type BillDetail struct {
	BillerReferenceId string                 `json:"billerReferenceId"`
	BillCode          string                 `json:"billCode"`
	BillNo            string                 `json:"billNo"`
	BillName          string                 `json:"billName"`
	BillShortName     string                 `json:"billShortName"`
	BillDescription   Description            `json:"billDescription"`
	BillSubCompany    string                 `json:"billSubCompany"`
	BillAmount        Amount                 `json:"billAmount"`
	AdditionalInfo    map[string]interface{} `json:"additionalInfo"`
	Status            string                 `json:"status"`
	Reason            Description            `json:"reason"`
}

func NewBillDetailsFromSqlNullString(billDetails sql.NullString) []BillDetail {
	var result []BillDetail
	if !billDetails.Valid {
		return result
	}
	if err := json.Unmarshal([]byte(billDetails.String), &result); err != nil {
		return result
	}
	return result
}

type VACommonField struct {
	PartnerServiceId    string        `json:"partnerServiceId" validate:"required"`
	CustomerNo          string        `json:"customerNo" validate:"required"`
	VirtualAccountNo    string        `json:"virtualAccountNo" validate:"required,max=28"`
	VirtualAccountName  string        `json:"virtualAccountName"`
	VirtualAccountEmail string        `json:"virtualAccountEmail"`
	VirtualAccountPhone string        `json:"virtualAccountPhone"`
	TrxId               string        `json:"trxId"`
	BillDetails         []BillDetail  `json:"billDetails"`
	FreeTexts           []Description `json:"freeTexts"`
}
