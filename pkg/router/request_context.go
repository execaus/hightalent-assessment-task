package router

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

type RequestContext struct {
	BaseContext
	cancel     context.CancelFunc
	writer     http.ResponseWriter
	isFinished bool
}

func NewRequestContext(ctx context.Context, cancel context.CancelFunc, writer http.ResponseWriter) *RequestContext {
	return &RequestContext{
		BaseContext: newBaseContext(ctx),
		cancel:      cancel,
		writer:      writer,
	}
}

func (c *RequestContext) SendCreated(data interface{}) {
	c.send(http.StatusCreated, data)
}

func (c *RequestContext) SendNotFound(data interface{}) {
	c.send(http.StatusCreated, data)
}

func (c *RequestContext) send(code int, data interface{}) {
	c.writer.Header().Set("Content-Type", "application/json")
	c.writer.WriteHeader(code)
	if err := json.NewEncoder(c.writer).Encode(data); err != nil {
		log.Printf("failed to encode response: %v", err)
	}
	c.isFinished = true
}
