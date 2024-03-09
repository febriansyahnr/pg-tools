package response

type Response struct {
	Code    string      `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Error   string      `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type SnapResponse struct {
	ResponseCode       string         `json:"responseCode"`
	ResponseMessage    string         `json:"responseMessage"`
	VirtualAccountData any            `json:"virtualAccountData,omitempty"`
	AdditionalInfo     map[string]any `json:"additionalInfo,omitempty"`
}
