/**
 * @file map_based_handle.go
 * @author zhaodezhen
 * @brief
 * @version 0.1
 * @date 2022-03-17
 * @copyright Copyright (c) 2021 The zweb-go Authors. All rights reserved.
**/

package zweb

import "net/http"

type Handler interface { //我实现了 ServeHTTP  ，然后实现 Routable
	//http.Handler
	ServeHTTP(c *Context)
	Routable
}

type Routable interface { // 我实现了  Route
	Route(method, pattern string, handleFunc func(ctx *Context))
}

type HandlerBasedOnMap struct { // 我实现了  Handler
	handlers map[string]func(ctx *Context)
}

func (h *HandlerBasedOnMap) Route(method, pattern string, handleFunc func(ctx *Context)) {
	key := h.Key(method, pattern)
	h.handlers[key] = handleFunc
}

// 我实现了 http.Handler
func (h *HandlerBasedOnMap) ServeHTTP(c *Context) {
	// 处理路由
	key := h.Key(c.R.Method, c.R.URL.Path)
	if handle, ok := h.handlers[key]; ok {
		handle(c)
	} else {
		c.W.WriteHeader(http.StatusNotFound)
		c.W.Write([]byte("Not Found"))
	}
}

func (h *HandlerBasedOnMap) Key(method, pattern string) string {
	return method + "#" + pattern
}

// 确保完全实现了 Handler
var _ Handler = &HandlerBasedOnMap{}

func NewHandlerBasedOnMap() Handler {
	return &HandlerBasedOnMap{
		handlers: make(map[string]func(ctx *Context)),
	}
}
