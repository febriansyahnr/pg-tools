package http_snapCore

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/febrianpaper/pg-tools/constant"
	model_snapCore "github.com/febrianpaper/pg-tools/internal/model/snapCore"
	model_trxHistory "github.com/febrianpaper/pg-tools/internal/model/trxHistory"
)

const (
	InquiryVAUrl = "/api/v1.0/transfer-va/inquiry"
)

// InquiryVA implements port.SnapCorePort.
func (s *SnapCoreAdapter) InquiryVA(ctx context.Context, token string, req *model_snapCore.InquiryRequestData) (*model_snapCore.InquiryData, error) {
	url := s.config.SnapCoreURL + InquiryVAUrl

	now := time.Now()
	timeStamp := now.Format(constant.SnapDateFormatLayout)

	signatureParam := serviceSignatureParams{
		TimeStamp:   timeStamp,
		AccessToken: token,
		HttpMethod:  "POST",
		EndpoinUrl:  InquiryVAUrl,
	}
	signature, err := s.GetServiceSignature(ctx, signatureParam, req)
	if err != nil {
		return nil, fmt.Errorf("error when calling snapcore generate service signature: %w", err)
	}

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"X-TIMESTAMP":   timeStamp,
		"X-SIGNATURE":   signature,
		"X-PARTNER-ID":  s.secret.SnapCoreKey,
		"X-EXTERNAL-ID": fmt.Sprintf("%d", time.Now().UnixMilli()),
		"CHANNEL-ID":    "751",
	}
	fmt.Println("====================")
	slog.Info("[Inquiry VA]", "url", url, "headers", headers, "body", req)

	respBytes, status, err := s.httpClient.POST(ctx, url, req, headers)
	if err != nil {
		slog.Info("Inquiry", "response", string(respBytes))
		return nil, fmt.Errorf("error when calling snapcore inquiry va: %w", err)
	}
	if status != http.StatusOK {
		slog.Info("Inquiry", "response", string(respBytes))
		return nil, fmt.Errorf("error when calling snapcore inquiry va status not ok")
	}
	var resp model_snapCore.SnapVAResponse
	var inqData model_snapCore.InquiryData
	resp.VirtualAccountData = &inqData
	if err := json.Unmarshal(respBytes, &resp); err != nil {
		return nil, fmt.Errorf("error when unmarshaling response: %w", err)
	}
	slog.Info("[Inquiry VA]", "response", resp)

	trxLog := model_trxHistory.NewTrxLogVA("inquiry", model_trxHistory.NewTrxLogParams{
		Number:     req.VirtualAccountNo,
		Additional: headers,
		Request:    req.Json(),
		Response:   string(respBytes),
		Status:     1,
	})
	s.logRepo.Create(ctx, trxLog)

	if resp.ResponseCode[:3] != "200" {
		return nil, fmt.Errorf("%s", resp.ResponseMessage)
	}

	return &inqData, nil
}
