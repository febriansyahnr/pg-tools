package service_virtualAccount

import (
	"context"

	model_snapCore "github.com/febrianpaper/pg-tools/internal/model/snapCore"
	model_trxHistory "github.com/febrianpaper/pg-tools/internal/model/trxHistory"
	commands_virtualAccount "github.com/febrianpaper/pg-tools/internal/service/virtualAccount/commands"
	queries_virtualAccount "github.com/febrianpaper/pg-tools/internal/service/virtualAccount/queries"
	"github.com/febrianpaper/pg-tools/pkg/redisExt"
	"github.com/febrianpaper/pg-tools/port"
)

type (
	IVirtualAccountService interface {
		Commands
		Queries
	}
	Commands interface {
		ResetQueues(ctx context.Context) error
		PaymentVA(ctx context.Context, command commands_virtualAccount.PaymentVACommands) error
	}
	Queries interface {
		GetToken(ctx context.Context) (string, error)
		// InquiryVA get va data from snapCore
		InquiryVA(ctx context.Context, query queries_virtualAccount.InquiryVAQuery) (*model_snapCore.InquiryData, string, error)
		GetAllQueue(ctx context.Context) ([]model_trxHistory.TrxQueue, error)
	}

	appCommands struct {
		commands_virtualAccount.ResetQueuesHandler
		commands_virtualAccount.PaymentVAHandler
	}
	appQueries struct {
		queries_virtualAccount.GetTokenHandler
		queries_virtualAccount.InquiryVAHandler
		queries_virtualAccount.GetAllQueueHandler
	}

	VirtualAccountService struct {
		SnapAuthAdapter port.SnapCorePort
		CacheAdpter     redisExt.IRedisExt
		TrxQueueAdapter port.TrxQueuePort
		appCommands
		appQueries
	}
)

var _ IVirtualAccountService = (*VirtualAccountService)(nil)

func New(snapAuthAdapter port.SnapCorePort, cacheAdpter redisExt.IRedisExt, TrxQueueAdapter port.TrxQueuePort) *VirtualAccountService {
	svc := &VirtualAccountService{
		SnapAuthAdapter: snapAuthAdapter,
		CacheAdpter:     cacheAdpter,
		TrxQueueAdapter: TrxQueueAdapter,
		appCommands: appCommands{
			ResetQueuesHandler: commands_virtualAccount.NewResetQueuesHandler(TrxQueueAdapter),
			PaymentVAHandler:   commands_virtualAccount.NewPaymentVAHandler(snapAuthAdapter, TrxQueueAdapter),
		},
	}

	svc.appQueries.GetTokenHandler = queries_virtualAccount.NewGetTokenHandler(snapAuthAdapter, cacheAdpter)
	svc.appQueries.InquiryVAHandler = queries_virtualAccount.NewInquiryVAHandler(snapAuthAdapter, TrxQueueAdapter)
	svc.appQueries.GetAllQueueHandler = queries_virtualAccount.NewGetAllQueueHandler(TrxQueueAdapter)

	return svc
}
