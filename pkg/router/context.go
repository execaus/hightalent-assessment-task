package router

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"hightalent-assessment-task/internal/service"
	"io"
	"log"
	"net/http"
	"reflect"

	"github.com/google/uuid"
)

type Context interface {
	context.Context
	IsAbort() bool
	BindJSON(recipient interface{}) error
	Abort(err error)
	SendCreated(data interface{})
	SendNotFound(data interface{})
	SendOK(data interface{})
	GetUserID() (uuid.UUID, error)
	SetUserID(id string) error
	GetHeader(key string) (string, error)
	GetIntDynamicValue(key string) (int, error)
}

type RequestContext struct {
	cancel     context.CancelFunc
	isFinished bool
	context.Context
	isAbort       bool
	DynamicValues dynamicPathValues
	ErrorToCode   map[reflect.Type]int
	request       *http.Request
	writer        http.ResponseWriter
	ResponseBody  []byte
	userID        *uuid.UUID
}

func (c *RequestContext) GetIntDynamicValue(key string) (int, error) {
	val, ok := c.DynamicValues[key]
	if !ok {
		return 0, fmt.Errorf("dynamic value for key %q not found", key)
	}
	var i int
	_, err := fmt.Sscanf(fmt.Sprintf("%v", val), "%d", &i)
	if err != nil {
		return 0, fmt.Errorf("dynamic value for key %q is not an int: %v", key, err)
	}
	return i, nil
}

func (c *RequestContext) GetHeader(key string) (string, error) {
	value := c.request.Header.Get(key)
	if value == "" {
		return "", fmt.Errorf("header %q not found", key)
	}
	return value, nil
}

func (c *RequestContext) GetUserID() (uuid.UUID, error) {
	if c.userID == nil {
		return uuid.Nil, errors.New("user ID not set")
	}

	return *c.userID, nil
}

func (c *RequestContext) SetUserID(id string) error {
	parsed, err := uuid.Parse(id)
	if err != nil {
		c.userID = nil
		return err
	}

	c.userID = &parsed

	return nil
}

func NewRequestContext(ctx context.Context, cancel context.CancelFunc, writer http.ResponseWriter, request *http.Request) *RequestContext {
	return &RequestContext{
		Context: ctx,
		ErrorToCode: map[reflect.Type]int{
			reflect.TypeOf(BadRequestError{}): http.StatusBadRequest,
		},
		cancel:  cancel,
		writer:  writer,
		request: request,
	}
}

func (c *RequestContext) IsAbort() bool {
	return c.isAbort
}

func (c *RequestContext) BindJSON(recipient interface{}) error {
	body, err := io.ReadAll(c.request.Body)
	if err != nil {
		return NewBadRequestError("failed to read request body: " + err.Error())
	}
	defer c.request.Body.Close()

	if err = json.Unmarshal(body, recipient); err != nil {
		return NewBadRequestError("invalid JSON: " + err.Error())
	}
	return nil
}

func (c *RequestContext) Abort(err error) {
	c.isAbort = true

	var httpError HTTPError
	var businessLoginError service.BusinessLoginError

	if errors.As(err, &httpError) {
		c.writer.WriteHeader(httpError.StatusCode())
		return
	}
	if errors.As(err, &businessLoginError) {
		c.writer.WriteHeader(http.StatusBadRequest)
		return
	}

	c.writer.WriteHeader(http.StatusInternalServerError)
}

func (c *RequestContext) SendCreated(data interface{}) {
	c.send(http.StatusCreated, data)
}

func (c *RequestContext) SendNotFound(data interface{}) {
	c.send(http.StatusNotFound, data)
}

func (c *RequestContext) SendOK(data interface{}) {
	c.send(http.StatusOK, data)
}

func (c *RequestContext) send(code int, data interface{}) {
	c.writer.Header().Set("Content-Type", "application/json")
	c.writer.WriteHeader(code)
	if err := json.NewEncoder(c.writer).Encode(data); err != nil {
		log.Printf("failed to encode response: %v", err)
	}
	c.ResponseBody, _ = json.Marshal(data)
	c.isFinished = true
}
