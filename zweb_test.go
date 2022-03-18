/**
 * @file server_test
 * @author zhaodezhen
 * @brief
 * @version 0.1
 * @date 2022-03-16
 * @copyright Copyright (c) 2021 The zweb-go Authors. All rights reserved.
**/

package zweb

import (
	"log"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {

	server := NewHttpServer(MetricsBuilder, Auth)
	//server.Route("POST", "/", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprint(w, "hello")
	//})

	server.Route(http.MethodPost, "/signup", SignUp)
	if err := server.Start(":1998"); err != nil {
		log.Fatal(err.Error())
	}
}

type SignUpReq struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func SignUp(ctx *Context) {
	req := &SignUpReq{}
	if err := ctx.ReadJson(req); err != nil {
		ctx.ErrorJson(err.Error())
		return
	}

	ctx.SuccessJson(req)
	return
}
