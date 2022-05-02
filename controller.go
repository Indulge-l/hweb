package main

import (
	"context"
	"fmt"
	"hweb/framework"
	"log"
	"time"
)

func FooControllerHandler(ctx *framework.Context) error {
	fmt.Println("foo")
	duration, cancel := context.WithTimeout(ctx.BaseContext(), time.Duration(1*time.Second))
	defer cancel()

	finish := make(chan struct{}, 1)
	panicChan := make(chan interface{}, 1)
	go func() {
		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()
		time.Sleep(10 * time.Second)
		ctx.Json(200, "ok")
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
