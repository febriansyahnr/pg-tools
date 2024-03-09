package http_snapCore

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
)

const (
	AuthSignatureUrl    = "/api/v1.0/utilities/signature-auth"
	ServiceSignatureUrl = "/api/v1.0/utilities/signature-service"
)

func (s *SnapCoreAdapter) GetAuthSignature(ctx context.Context, body json.RawMessage, timeStamp string) (string, error) {
	url := s.config.SnapCoreURL + AuthSignatureUrl

	headers := map[string]string{
		"X-TIMESTAMP":  timeStamp,
		"X-CLIENT-KEY": s.secret.SnapCoreKey,
	}

	type Response struct {
		Signature string `json:"signature"`
	}
	var resp Response

	fmt.Println("====================")
	slog.Info("[Get Auth Signature]", "url", url, "headers", headers, "body", body)
	respByte, _, err := s.httpClient.POST(ctx, url, body, headers)
	if err != nil {
		return "", errors.New("need to re-request")
	}

	if err := json.Unmarshal(respByte, &resp); err != nil {
		return "", fmt.Errorf("error when unmarshaling response: %w", err)
	}
	slog.Info("[Get Auth Signature]", "response", resp)

	return resp.Signature, nil
}

type serviceSignatureParams struct {
	TimeStamp   string `json:"timeStamp"`
	AccessToken string `json:"accessToken"`
	HttpMethod  string `json:"httpMethod"`
	EndpoinUrl  string `json:"endpointUrl"`
}

func (s *SnapCoreAdapter) GetServiceSignature(ctx context.Context, param serviceSignatureParams, body any) (string, error) {
	url := s.config.SnapCoreURL + ServiceSignatureUrl

	headers := map[string]string{
		"X-TIMESTAMP":     param.TimeStamp,
		"X-CLIENT-SECRET": s.secret.SnapCoreSecret,
		"AccessToken":     "Bearer " + param.AccessToken,
		"HttpMethod":      param.HttpMethod,
		"EndpoinUrl":      param.EndpoinUrl,
	}
	fmt.Println("====================")
	slog.Info("[Get Service Signature]", "url", url, "headers", headers, "body", body)

	respByte, _, err := s.httpClient.POST(ctx, url, body, headers)
	if err != nil {
		return "", errors.New("need to re-request")
	}

	type Response struct {
		Signature string `json:"signature"`
	}
	var resp Response

	if err := json.Unmarshal(respByte, &resp); err != nil {
		return "", fmt.Errorf("error when unmarshaling response: %w", err)
	}

	slog.Info("[Get Service Signature]", "response", resp)

	return resp.Signature, nil
}
