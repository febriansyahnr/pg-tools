package httputil

import (
	"bytes"
	"context"

	"encoding/json"
	"io"
	"net/http"
)

func RequestHitAPI(
	ctx context.Context,
	method string,
	uri string,
	data interface{},
	header map[string]string,
) (
	res []byte,
	code int,
	err error,
) {

	httpClient := &http.Client{}
	// httpClient.Transport = newrelic.NewRoundTripper(httpClient.Transport)
	request, err := assertTypeRequest(data, method, uri)
	if err != nil {
		return res, code, err
	}

	for k, v := range header {
		request.Header.Add(k, v)
	}

	request.Header.Set("Content-type", "application/json")

	response, err := httpClient.Do(request)
	if err != nil {
		return res, code, err
	}

	defer response.Body.Close()

	code = response.StatusCode

	res, err = io.ReadAll(response.Body)
	if err != nil {
		return res, code, err
	}

	if isHttpError := code != http.StatusOK && code != http.StatusCreated; isHttpError {
		var errRes map[string]interface{}
		err := json.Unmarshal(res, &errRes)
		return res, code, err
	}

	return res, code, err
}

func assertTypeRequest(data interface{}, method string, uri string) (request *http.Request, err error) {
	if data == nil {
		request, err = http.NewRequest(method, uri, nil)
		return
	}

	paramReq, _ := json.Marshal(data)
	request, err = http.NewRequest(method, uri, bytes.NewBuffer(paramReq))

	return
}

func NewHttpRequest(
	ctx context.Context,
	method string,
	url string,
	header map[string]string,
	bodyReq interface{},
) (*http.Request, error) {
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	// request = newrelic.RequestWithTransactionContext(request, newRelicExt.GetTxnFromCtx(ctx))

	// Set or modify the request body
	if bodyReq != nil {
		// For handling form data
		if formData, ok := bodyReq.(*bytes.Buffer); ok {
			request.Body = io.NopCloser(formData)
		} else {
			payloadBody, _ := json.Marshal(bodyReq)
			request.Body = io.NopCloser(bytes.NewBuffer(payloadBody))
		}
	}

	for k, v := range header {
		request.Header.Add(k, v)
	}

	return request, nil
}
