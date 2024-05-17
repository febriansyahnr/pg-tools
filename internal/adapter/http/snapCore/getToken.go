package http_snapCore

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/febrianpaper/pg-tools/constant"
)

const (
	// B2BUrl = "/snap-core/api/v1.0/access-token/b2b"
	B2BUrl = "/api/v1.0/access-token/b2b"
)

// GetToken implements port.SnapCorePort.
func (s *SnapCoreAdapter) GetToken(ctx context.Context) (string, error) {
	now := time.Now()
	timeStamp := now.Format(constant.SnapDateFormatLayout)
	body := json.RawMessage(`{"grantType":"client_credential"}`)

	signature, err := s.GetAuthSignature(ctx, body, timeStamp)
	if err != nil {
		return "", err
	}

	url := s.config.SnapCoreURL + s.config.SnapCoreBase + B2BUrl
	header := map[string]string{
		"X-TIMESTAMP":  timeStamp,
		"X-CLIENT-KEY": s.secret.SnapCoreKey,
		"X-SIGNATURE":  signature,
	}
	fmt.Println("====================")
	slog.Info("[Get Token]", "url", url, "headers", header, "body", body)

	respByte, status, err := s.httpClient.POST(ctx, url, body, header)
	if err != nil {
		return "", errors.New("error when calling snapcore token b2b")
	}
	if status != http.StatusOK {
		return "", errors.New("error when calling snapcore token b2b status not ok")
	}

	type Response struct {
		ResponseCode    string `json:"responseCode"`
		ResponseMessage string `json:"responseMessage"`
		AccessToken     string `json:"accessToken"`
		TokenType       string `json:"tokenType"`
		ExpiresIn       string `json:"expiresIn"`
	}
	var resp Response
	if err := json.Unmarshal(respByte, &resp); err != nil {
		return "", fmt.Errorf("error when unmarshaling response: %w", err)
	}

	slog.Info("[Get Token]", "resp", resp)

	if resp.ResponseCode[:3] != "200" {
		return "", fmt.Errorf("error when calling snapcore token b2b status not ok: %s", resp.ResponseMessage)
	}

	return resp.AccessToken, nil
}
