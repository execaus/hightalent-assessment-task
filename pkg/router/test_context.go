package router

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

type TestContext struct {
	BaseContext
}

func NewTestContext(ctx context.Context) *TestContext {
	return &TestContext{
		BaseContext: newBaseContext(ctx),
	}
}

func (c *TestContext) PutRequestBody(body interface{}) error {
	bytes, err := json.Marshal(body)
	if err != nil {
		log.Printf("Ошибка сериализации тела запроса: %+v, ошибка: %s\n", body, err.Error())
		return err
	}

	c.Request.Body = bytes

	return nil
}

func (c *TestContext) SendCreated(data interface{}) {
	bytes, _ := json.Marshal(data)
	c.Response.Body = bytes
	c.Response.StatusCode = http.StatusCreated
}
