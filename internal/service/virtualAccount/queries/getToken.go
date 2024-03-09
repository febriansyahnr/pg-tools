package queries_virtualAccount

import (
	"context"
	"time"

	"github.com/febrianpaper/pg-tools/pkg/redisExt"
	"github.com/febrianpaper/pg-tools/port"
)

type (
	GetTokenHandler struct {
		snapCore port.SnapCorePort
		Cache    redisExt.IRedisExt
	}
)

func NewGetTokenHandler(snapCore port.SnapCorePort, cache redisExt.IRedisExt) GetTokenHandler {
	return GetTokenHandler{
		snapCore: snapCore,
		Cache:    cache,
	}
}

func (h *GetTokenHandler) GetToken(ctx context.Context) (string, error) {

	// get from cache
	cacheToken, err := h.Cache.Get(ctx, "b2btoken").Result()
	if err != nil || cacheToken == "" {
		token, err := h.snapCore.GetToken(ctx)
		if err != nil {
			return "", err
		}
		err = h.Cache.Set(ctx, "b2btoken", token, time.Minute*14).Err()
		if err != nil {
			return "", err
		}
	}

	return cacheToken, nil
}
