package router

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http/httptest"
	"time"
)

type TestContext struct {
	RequestContext
}

func NewTestContext(ctx context.Context, timeout time.Duration) *TestContext {
	c, cancel := context.WithTimeout(ctx, timeout)

	return &TestContext{
		RequestContext: *NewRequestContext(
			c,
			cancel,
			httptest.NewRecorder(),
			httptest.NewRequest("POST", "/", nil),
		),
	}
}

func (c *TestContext) PutRequestBody(body interface{}) error {
	b, err := json.Marshal(body)
	if err != nil {
		log.Printf("Failed to serialize request body: %+v, error: %s\n", body, err.Error())
		return err
	}

	c.request.Body = io.NopCloser(bytes.NewReader(b))
	c.request.ContentLength = int64(len(b))
	c.request.Header.Set("Content-Type", "application/json")

	return nil
}
