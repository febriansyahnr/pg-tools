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

const PaymentVAUrl = "/api/v1.0/transfer-va/payment"

// paymentVA implements port.SnapCorePort.
func (s *SnapCoreAdapter) PaymentVA(ctx context.Context, token string, req *model_snapCore.VAPaymentRequest) error {
	url := s.config.SnapCoreURL + PaymentVAUrl

	now := time.Now()
	timeStamp := now.Format(constant.SnapDateFormatLayout)

	signatureParam := serviceSignatureParams{
		TimeStamp:   timeStamp,
		AccessToken: token,
		HttpMethod:  "POST",
		EndpoinUrl:  PaymentVAUrl,
	}
	signature, err := s.GetServiceSignature(ctx, signatureParam, req)
	if err != nil {
		return fmt.Errorf("error when calling snapcore generate service signature: %w", err)
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
	slog.Info("[payment VA]", "url", url, "headers", headers, "body", req)

	respBytes, status, err := s.httpClient.POST(ctx, url, req, headers)
	if err != nil {
		return fmt.Errorf("error when calling snapcore payment va: %w", err)
	}
	if status != http.StatusOK {
		return fmt.Errorf("error when calling snapcore payment va status not ok")
	}
	var resp model_snapCore.SnapVAResponse
	if err := json.Unmarshal(respBytes, &resp); err != nil {
		return fmt.Errorf("error when unmarshaling response: %w", err)
	}
	slog.Info("[payment VA]", "response", resp)
	trxLog := model_trxHistory.NewTrxLogVA("payment", model_trxHistory.NewTrxLogParams{
		Number:     req.VirtualAccountNo,
		Additional: headers,
		Request:    req.Json(),
		Response:   string(respBytes),
		Status:     1,
	})

	s.logRepo.Create(ctx, trxLog)

	if resp.ResponseCode[:3] != "200" {
		return fmt.Errorf("%s", resp.ResponseMessage)
	}

	return nil
}
