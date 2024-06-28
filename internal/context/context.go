package context

import (
	"context"
	"net/http"

	"go.uber.org/zap"
)

type Context struct {
	context.Context
	*zap.Logger
	*http.Client
}

func NewContext(ctx context.Context, log *zap.Logger, cli *http.Client) *Context {
	return &Context{
		Context: ctx,
		Logger:  log,
		Client:  cli,
	}
}
