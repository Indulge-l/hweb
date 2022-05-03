package middleware

import (
	"hweb/framework"
	"log"
	"time"
)

// 获取请求请求处理花费时间
func Cost() framework.ControllerHandler {
	return func(c *framework.Context) error {
		start := time.Now()
		c.Next()
		end := time.Now()

		cost := end.Sub(start)
		log.Printf("api uri(%+v),cost(%+v)", c.GetRequest().RequestURI, cost.Seconds())
		return nil
	}
}
