package framework

import (
	"context"
	"net/http"
	"time"
)

type Context struct {
	request  *http.Request
	response http.ResponseWriter
}

func NewContext() *Context {
	return &Context{}
}

func (ctx *Context) BaseContext() context.Context {
	return ctx.request.Context()
}

func (ctx *Context) Deadline() (deadline time.Time, ok bool) {
	return time.Now(), false
}

func (ctx *Context) Err() error {
	return nil
}

func (ctx *Context) Done() <-chan struct{} {
	return ctx.BaseContext().Done()
}
