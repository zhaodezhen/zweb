/**
 * @file context.go
 * @author zhaodezhen
 * @brief
 * @version 0.1
 * @date 2022-03-16
 * @copyright Copyright (c) 2021 The zweb-go Authors. All rights reserved.
**/

package zweb

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Context struct {
	W http.ResponseWriter
	R *http.Request
}

func (c *Context) ReadJson(obj interface{}) error {
	body, err := io.ReadAll(c.R.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, obj)
}

func (c *Context) WriterJson(code int, resp interface{}) error {
	c.W.WriteHeader(code)
	response, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	_, err = c.W.Write(response)
	return err
}

func (c *Context) SuccessJson(resp interface{}) error {
	defer func() {
		//打印日志
		log.Println("URI:" + c.R.RequestURI)
		log.Println("METHOD:" + c.R.Method)
	}()
	return c.WriterJson(http.StatusOK, resp)
}

func (c *Context) ErrorJson(resp interface{}) error {
	return c.WriterJson(http.StatusInternalServerError, resp)
}

func (c *Context) BadRequestJson(resp interface{}) error {
	return c.WriterJson(http.StatusNotFound, resp)
}

func NewContext(writer http.ResponseWriter, request *http.Request) *Context {
	return &Context{
		W: writer,
		R: request,
	}
}
