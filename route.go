package main

import "hweb/framework"

func registerRouter(core *framework.Core) {
	core.Get("foo", FooControllerHandler)
}
