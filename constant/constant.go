package constant

type constantKey string

const (
	// ctxTraceIdKey is the context key for trace id
	CtxTraceIdKey    constantKey = "trace_id"
	HeaderCtxKey     constantKey = "snap-header"
	IntlCtxHeaderKey constantKey = "intl-header"

	EnvironmentDevelopment = "development"
	EnvironmentLocal       = "local"
	EnvironmentStaging     = "staging"
	EnvironmentProduction  = "production"
)

const (
	DefaultCurrencyIDR string = "IDR"
)
