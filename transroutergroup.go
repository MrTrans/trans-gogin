// Copyright 2020 trans owner

package transgin

import "net/http"

// Context is an implementation of the context.Context sub-package.
// context.Context is very extensible and developers can override
// its methods if that is actually needed.
type Context interface {
	String() string
	Request() *http.Request
}

type handleFunc func(*Context)

// TIRouter 路由
type TIRouter interface {
	GET(string, ...handleFunc)
	POST(string, ...handleFunc)
}
