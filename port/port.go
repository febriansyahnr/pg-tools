package port

import "github.com/febrianpaper/pg-tools/pkg/redisExt"

type Adapters struct {
	SnapCore SnapCorePort
	Cache    redisExt.IRedisExt
	TrxQueue TrxQueuePort
	TrxLog   TrxLogPort
}
