package model_snapCore

type RequestCreateVA struct {
	CustomerNo     string                 `json:"customerNo"`
	MerchantID     string                 `json:"mid"`
	Number         string                 `json:"vaNumber"`
	AccountName    string                 `json:"accountName" validate:"required"`
	AccountEmail   string                 `json:"accountEmail"`
	AccountPhone   string                 `json:"accountPhone"`
	SubCompany     string                 `json:"subCompany"`
	TotalAmount    Amount                 `json:"totalAmount" validate:"required"`
	FeeAmount      Amount                 `json:"feeAmount"`
	Acquirer       string                 `json:"acquirer" validate:"required"`
	BillDetails    []BillDetail           `json:"billDetails"`
	FreeTexts      []Description          `json:"freeTexts"`
	IsClosedAmount bool                   `json:"isClosedAmount"`
	IsSingleUse    bool                   `json:"isSingleUse"`
	AdditionalInfo map[string]interface{} `json:"additionalInfo"`
}
