package main

import (
	"hweb/framework"
	"hweb/framework/middleware"
)

func registerRouter(core *framework.Core) {
	//core.Get("foo", FooControllerHandler)

	core.Post("/user/login", UserLoginController)

	subjectApi := core.Group("/subject", middleware.Test1())
	{
		temp := subjectApi.Group("/ssstest", middleware.Test2())
		{
			temp.Get("/:id", middleware.Test1(), SubjectGetController)
		}
		subjectApi.Post("/add", SubjectAddController)
		subjectApi.Delete("/:id", SubjectDelController)
		subjectApi.Put("/:id", SubjectUpdateController)
		subjectApi.Get("/:id", SubjectGetController)
		subjectApi.Get("/list/all", SubjectListController)
	}
}
