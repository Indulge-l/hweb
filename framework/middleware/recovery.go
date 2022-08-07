package middleware

import (
	"hweb/framework"
	"net/http"
)

// Recovery 中间件，防止整个进程panic
func Recovery() framework.ControllerHandler {
	return func(c *framework.Context) error {
		defer func() {
			if err := recover(); err != nil {
				c.Json(http.StatusInternalServerError, err)
			}
		}()
		c.Next()
		return nil
	}
}
