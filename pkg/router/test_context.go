package router

import "context"

type TestContext struct {
	BaseContext
}

func NewTestContext(ctx context.Context) *TestContext {
	return &TestContext{
		BaseContext: newBaseContext(ctx),
	}
}
