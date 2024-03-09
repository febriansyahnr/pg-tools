package response

import (
	"encoding/json"
	"net/http"
)

func SendResponseOK(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	resp := Response{
		Code:    GetHttpCodeService(HttpStatusOK),
		Message: http.StatusText(http.StatusOK),
		Data:    data,
	}

	return json.NewEncoder(w).Encode(resp)
}

func SendResponseError(w http.ResponseWriter, errType string, errMessage error) error {
	code, statusCode := HttpStatusErrorCode(errType)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	resp := Response{
		Code:    GetHttpCodeService(code),
		Error:   errMessage.Error(),
		Message: http.StatusText(statusCode),
	}

	return json.NewEncoder(w).Encode(resp)
}

func SendResponseCreated(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	resp := Response{
		Code:    GetHttpCodeService(HttpStatusCreated),
		Data:    data,
		Message: "Created",
	}

	err := json.NewEncoder(w).Encode(resp)
	return err
}

func SendVAResponseOK(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	return json.NewEncoder(w).Encode(data)
}
