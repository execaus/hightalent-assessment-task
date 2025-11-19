package router

import (
	"context"
	"encoding/json"
)

type Context interface {
	context.Context
	IsAbort() bool
	BindJSON(interface{}) error
	Abort(error)
	SendCreated(interface{})
}

type BaseContext struct {
	context.Context
	isAbort  bool
	Response struct {
		StatusCode int
		Body       []byte
	}
}

func newBaseContext(ctx context.Context) BaseContext {
	return BaseContext{
		Context: ctx,
	}
}

func (c *BaseContext) IsAbort() bool {
	return c.isAbort
}

func (c *BaseContext) BindJSON(recipient interface{}) error {
	// TODO
	return json.Unmarshal([]byte{}, recipient)
}

func (c *BaseContext) Abort(err error) {
	c.isAbort = true
}

func (c *BaseContext) SendCreated(data interface{}) {

}
