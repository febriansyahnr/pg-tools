package http_snapCore

import (
	"github.com/febrianpaper/pg-tools/config"
	"github.com/febrianpaper/pg-tools/pkg/httpRequestExt"
	"github.com/febrianpaper/pg-tools/port"
)

type SnapCoreAdapter struct {
	config     *config.Config
	secret     *config.Secret
	httpClient httpRequestExt.IHTTPRequest
	logRepo    port.TrxLogPort
}

func New(client httpRequestExt.IHTTPRequest, config *config.Config, secret *config.Secret, logRepo port.TrxLogPort) *SnapCoreAdapter {
	return &SnapCoreAdapter{
		httpClient: client,
		config:     config,
		secret:     secret,
		logRepo:    logRepo,
	}
}

var (
	_ port.SnapCorePort     = (*SnapCoreAdapter)(nil)
	_ port.SnapCoreIntlPort = (*SnapCoreAdapter)(nil)
)
