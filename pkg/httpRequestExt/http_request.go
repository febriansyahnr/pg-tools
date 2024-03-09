package httpRequestExt

import (
	"context"

	httputil "github.com/febrianpaper/pg-tools/pkg/util/http"
)

type IHTTPRequest interface {
	GET(ctx context.Context, uri string, header map[string]string) ([]byte, int, error)
	POST(ctx context.Context, uri string, data interface{}, header map[string]string) ([]byte, int, error)
}

type HTTPRequest struct{}

func New() IHTTPRequest {
	return &HTTPRequest{}
}

// POST implements IHTTPRequest.
func (*HTTPRequest) POST(ctx context.Context, uri string, data interface{}, header map[string]string) ([]byte, int, error) {
	return httputil.RequestHitAPI(ctx, "POST", uri, data, header)
}

// GET implements IHTTPRequest.
func (*HTTPRequest) GET(ctx context.Context, uri string, header map[string]string) ([]byte, int, error) {
	return httputil.RequestHitAPI(ctx, "GET", uri, nil, header)
}
