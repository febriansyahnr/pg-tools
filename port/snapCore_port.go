package port

import (
	"context"

	model_snapCore "github.com/febrianpaper/pg-tools/internal/model/snapCore"
)

type SnapCorePort interface {
	GetToken(ctx context.Context) (string, error)
	InquiryVA(ctx context.Context, token string, req *model_snapCore.InquiryRequestData) (*model_snapCore.InquiryData, error)
	PaymentVA(ctx context.Context, token string, req *model_snapCore.VAPaymentRequest) error
}

type SnapCoreIntlPort interface {
	CreateVA(ctx context.Context, request *model_snapCore.RequestCreateVA) error
}
