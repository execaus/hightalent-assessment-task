package router

import "context"

type RequestContext struct {
	BaseContext
	cancel context.CancelFunc
}

func NewRequestContext(ctx context.Context) Context {
	ctx, cancel := context.WithCancel(ctx)

	return &RequestContext{
		BaseContext: newBaseContext(ctx),
		cancel:      cancel,
	}
}

func (c *RequestContext) SendCreated(data interface{}) {
	// TODO
}
