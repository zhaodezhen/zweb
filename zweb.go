/**
 * @file zweb.go
 * @author zhaodezhen
 * @brief
 * @version 0.1
 * @date 2022-03-15
 * @copyright Copyright (c) 2021 The zweb-go Authors. All rights reserved.
**/

package zweb

import (
	"net/http"
)

type Server interface {
	Routable
	Start(addr string) error
}

type sdkHttpServer struct {
	//Name string
	//Mux  http.ServeMux
	handler Handler
	root    Filter
}

// Route 注册路由
func (s *sdkHttpServer) Route(method, pattern string, handleFunc func(ctx *Context)) {
	//http.HandleFunc(pattern,
	//	func(writer http.ResponseWriter, request *http.Request) {
	//		ctx := NewContext(writer, request)
	//		handleFunc(ctx)
	//	})

	//key := s.handler.Key(method, pattern)
	//s.handler.handlers[key] = handleFunc

	s.handler.Route(method, pattern, handleFunc)
}

// Start 启动一个 http server
func (s *sdkHttpServer) Start(addr string) error {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		c := NewContext(writer, request)
		s.root(c)
	})
	return http.ListenAndServe(addr, nil)
}

func NewHttpServer(builders ...FilterBuilder) Server {
	handler := NewHandlerBasedOnMap()
	var root Filter = func(c *Context) {
		handler.ServeHTTP(c)
	}

	for i := len(builders) - 1; i >= 0; i-- {
		b := builders[i]
		root = b(root)

	}

	return &sdkHttpServer{
		handler: handler,
		root:    root,
	}
}
