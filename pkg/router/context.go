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
	Request struct {
		Body []byte
	}
	DynamicValues dynamicPathValues
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
	return json.Unmarshal(c.Request.Body, recipient)
}

func (c *BaseContext) Abort(err error) {
	c.isAbort = true
}
