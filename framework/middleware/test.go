package middleware

import (
	"fmt"
	"hweb/framework"
)

func Test1() framework.ControllerHandler {
	return func(ctx *framework.Context) error {
		fmt.Println("middleware pre test1")
		ctx.Next() // 调用Next往下调用，会自增contxt.index
		fmt.Println("middleware post test1")
		return nil
	}
}

func Test2() framework.ControllerHandler {
	return func(ctx *framework.Context) error {
		fmt.Println("middleware pre test2")
		ctx.Next() // 调用Next往下调用，会自增contxt.index
		fmt.Println("middleware post test2")
		return nil
	}
}
