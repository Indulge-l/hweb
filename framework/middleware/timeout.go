package middleware

import (
	"context"
	"fmt"
	"hweb/framework"
	"log"
	"time"
)

func Timeout(d time.Duration) framework.ControllerHandler {
	return func(ctx *framework.Context) error {
		duration, cancel := context.WithTimeout(ctx.BaseContext(), d)
		defer cancel()

		finish := make(chan struct{}, 1)
		panicChan := make(chan interface{}, 1)
		go func() {

			defer func() {
				if p := recover(); p != nil {
					panicChan <- p
				}
			}()
			ctx.Next()
			finish <- struct{}{}
		}()
		select {
		case p := <-panicChan:
			ctx.WriterMux().Lock()
			defer ctx.WriterMux().Unlock()
			log.Println(p)
			ctx.Json(500, "panic")
		case <-finish:
			fmt.Println("finish")
		case <-duration.Done():
			ctx.WriterMux().Lock()
			defer ctx.WriterMux().Unlock()
			ctx.Json(500, "time out")
			ctx.SetHasTimeout()
		}
		return nil
	}
}
