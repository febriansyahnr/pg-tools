package http_snapCore

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"

	model_snapCore "github.com/febrianpaper/pg-tools/internal/model/snapCore"
)

const (
	CreateVAUrl = "/api/v1.0/virtual-account/create"
)

func (s *SnapCoreAdapter) CreateVA(ctx context.Context, request *model_snapCore.RequestCreateVA) error {
	url := s.config.SnapCoreURL + CreateVAUrl
	header := map[string]string{
		"X-Internal-Service-Key": s.secret.InternalServiceKey,
	}
	payloadBody, err := json.Marshal(request)
	if err != nil {
		slog.Error("CreateVA", "error", err)
		return errors.New("failed marshal create va payload body")
	}

	respBytes, status, err := s.httpClient.POST(ctx, url, payloadBody, header)
	if err != nil {
		slog.Error("CreateVA", "error", err)
		return errors.New("error when post data to snap core")
	}

	type Response struct {
		Code    string `json:"code,omitempty"`
		Message string `json:"message,omitempty"`
	}
	var response Response
	if err := json.Unmarshal(respBytes, &response); err != nil {
		return err
	}
	if status > 299 {
		return fmt.Errorf("error snap core create va: %s", response.Message)
	}

	return nil
}
