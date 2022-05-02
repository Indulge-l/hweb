package main

import "hweb/framework"

func registerRouter(core *framework.Core) {
	//core.Get("foo", FooControllerHandler)

	core.Post("/user/login", UserLoginController)

	subjectApi := core.Group("/subject")
	{
		subjectApi.Post("/add", SubjectAddController)
		subjectApi.Delete("/:id", SubjectDeleteController)
		subjectApi.Put("/:id", SubjectUpdateController)
		subjectApi.Get("/:id", SubjectGetController)
		subjectApi.Get("/list/all", SubjectListController)
	}
}
