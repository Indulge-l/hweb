package framework

import (
	"fmt"
	"log"
	"net/http"
)

type Core struct {
	router map[string]ControllerHandler
}

func NewCore() *Core {
	return &Core{router: make(map[string]ControllerHandler)}
}

func (c *Core) Get(url string, handler ControllerHandler) {
	c.router[url] = handler
}

func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	// TODO
	log.Println("core.serveHTTP")
	ctx := NewContext(request, response)

	router := c.router["foo"]
	fmt.Println(router)
	if router == nil {
		return
	}
	log.Println("core.router")
	router(ctx)
}
