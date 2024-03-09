package model_snapCore

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/febrianpaper/pg-tools/constant"
)

type InquiryData struct {
	InquiryStatus         string         `json:"inquiryStatus"`
	InquiryReason         Description    `json:"inquiryReason"`
	PartnerServiceId      string         `json:"partnerServiceId"`
	CustomerNo            string         `json:"customerNo"`
	VirtualAccountNo      string         `json:"virtualAccountNo"`
	VirtualAccountName    string         `json:"virtualAccountName"`
	VirtualAccountEmail   string         `json:"virtualAccountEmail"`
	VirtualAccountPhone   string         `json:"virtualAccountPhone"`
	InquiryRequestId      string         `json:"inquiryRequestId"`
	TotalAmount           Amount         `json:"totalAmount"`
	SubCompany            string         `json:"subCompany"`
	BillDetails           []BillDetail   `json:"billDetails"`
	FreeTexts             []Description  `json:"freeTexts"`
	VirtualAccountTrxType string         `json:"virtualAccountTrxType"`
	FeeAmount             Amount         `json:"feeAmount"`
	AdditionalInfo        map[string]any `json:"additionalInfo"`
}

func (i *InquiryData) Json() string {
	b, err := json.Marshal(i)
	if err != nil {
		return fmt.Sprintf(`{"error": "%s"}`, err.Error())
	}
	return string(b)
}

func NewInquiryDataFromString(s string) (*InquiryData, error) {
	var i InquiryData
	err := json.Unmarshal([]byte(s), &i)
	if err != nil {
		return nil, err
	}
	return &i, nil
}

type InquiryRequestData struct {
	PartnerServiceId string `json:"partnerServiceId" validate:"required"`
	CustomerNo       string `json:"customerNo" validate:"required"`
	VirtualAccountNo string `json:"virtualAccountNo" validate:"required"`
	InquiryRequestId string `json:"inquiryRequestId" validate:"required"`
	ChannelCode      int    `json:"channelCode"`
}

func (i *InquiryRequestData) Json() string {
	b, err := json.Marshal(i)
	if err != nil {
		return fmt.Sprintf(`{"error": "%s"}`, err.Error())
	}
	return string(b)
}

type VAPaymentRequest struct {
	VACommonField
	PaymentRequestId        string         `json:"paymentRequestId"`
	ChannelCode             int            `json:"channelCode"`
	PaidAmount              Amount         `json:"paidAmount"`
	PaidBills               string         `json:"paidBills"`
	TotalAmount             Amount         `json:"totalAmount"`
	TrxDateTime             string         `json:"trxDateTime"`
	ReferenceNo             string         `json:"referenceNo"`
	JournalNum              string         `json:"journalNum"`
	PaymentType             int            `json:"paymentType"`
	FlagAdvise              string         `json:"flagAdvise"`
	SourceAcquirer          string         `json:"sourceAcquirer"`
	HashedSourceAccountNo   string         `json:"hashedSourceAccountNo"`
	CumulativePaymentAmount Amount         `json:"cumulativePaymentAmount"`
	SubCompany              string         `json:"subCompany"`
	AdditionalInfo          map[string]any `json:"additionalInfo"`
}

func (v *VAPaymentRequest) Json() string {
	b, err := json.Marshal(v)
	if err != nil {
		return fmt.Sprintf(`{"error": "%s"}`, err.Error())
	}
	return string(b)
}

func NewVAPaymentRequestFromInquiryData(inq *InquiryData) VAPaymentRequest {
	now := time.Now()
	return VAPaymentRequest{
		VACommonField: VACommonField{
			PartnerServiceId:    inq.PartnerServiceId,
			CustomerNo:          inq.CustomerNo,
			VirtualAccountNo:    inq.VirtualAccountNo,
			VirtualAccountName:  inq.VirtualAccountName,
			VirtualAccountEmail: inq.VirtualAccountEmail,
			VirtualAccountPhone: inq.VirtualAccountPhone,
			TrxId:               inq.InquiryRequestId,
			BillDetails:         inq.BillDetails,
			FreeTexts:           inq.FreeTexts,
		},
		PaymentRequestId: inq.InquiryRequestId,
		ChannelCode:      751,
		PaidAmount:       inq.TotalAmount,
		PaidBills:        "",
		TotalAmount:      inq.TotalAmount,
		TrxDateTime:      now.Format(constant.SnapDateFormatLayout),
		SubCompany:       inq.SubCompany,
		AdditionalInfo:   inq.AdditionalInfo,
	}
}
