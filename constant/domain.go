package constant

type tableType string

const (
	// CtxNewRelicTxnKey is the context key for newrelic app
	CtxNewRelicTxnKey constantKey = "newrelic_txn"
	// CtxSQLTableNameKey is the context key for sql table name
	CtxSQLTableNameKey tableType = "table_name"
	// CtxRabbitMQStartTime is the context key for rabbitmq start time
	CtxRabbitMQStartTime string = "start_time"
)
