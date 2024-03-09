package service

import (
	service_trxHistory "github.com/febrianpaper/pg-tools/internal/service/trxHistory"
	service_virtualAccount "github.com/febrianpaper/pg-tools/internal/service/virtualAccount"
)

type Services struct {
	VirtualAccountService service_virtualAccount.IVirtualAccountService
	TrxHistoryService     service_trxHistory.ITrxHistory
}
