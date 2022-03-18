/**
 * @file filter.go
 * @author zhaodezhen
 * @brief
 * @version 0.1
 * @date 2022-03-18
 * @copyright Copyright (c) 2021 The zweb-go Authors. All rights reserved.
**/

package zweb

import (
	"log"
	"time"
)

//AOP -  责任链

type FilterBuilder func(next Filter) Filter

type Filter func(c *Context)

var _ FilterBuilder = MetricsBuilder

func MetricsBuilder(next Filter) Filter {
	return func(c *Context) {
		start := time.Now().Nanosecond()
		next(c)
		end := time.Now().Nanosecond()
		log.Printf("执行了: %d 纳秒", end-start)
	}
}

func Auth(next Filter) Filter {
	return func(c *Context) {
		if c.R.URL.Path == "/info" {
			c.ErrorJson("鉴权失败咯·····")
			return
		}
		next(c)
	}
}
