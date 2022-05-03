package framework

import (
	"log"
	"net/http"
	"strings"
)

type Core struct {
	router      map[string]*Tree
	middlewares []ControllerHandler
}

func NewCore() *Core {
	router := map[string]*Tree{}
	router[http.MethodGet] = NewTree()
	router[http.MethodPut] = NewTree()
	router[http.MethodPost] = NewTree()
	router[http.MethodDelete] = NewTree()
	return &Core{router: router}
}

func (c *Core) Get(url string, handlers ...ControllerHandler) {
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router[http.MethodGet].AddRouter(url, allHandlers); err != nil {
		log.Fatalf("add router error:(%+v)", err)
	}
}

func (c *Core) Put(url string, handlers ...ControllerHandler) {
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router[http.MethodPut].AddRouter(url, allHandlers); err != nil {
		log.Fatalf("add router error:(%+v)", err)
	}
}

func (c *Core) Post(url string, handlers ...ControllerHandler) {
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router[http.MethodPost].AddRouter(url, allHandlers); err != nil {
		log.Fatalf("add router error:(%+v)", err)
	}
}

func (c *Core) Delete(url string, handlers ...ControllerHandler) {
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router[http.MethodDelete].AddRouter(url, allHandlers); err != nil {
		log.Fatalf("add router error:(%+v)", err)
	}
}

func (c *Core) FindRouteByRequest(req *http.Request) []ControllerHandler {
	uri := req.URL.Path
	upperMethod := strings.ToUpper(req.Method)
	if methodHandlers, ok := c.router[upperMethod]; ok {
		return methodHandlers.FindHandler(uri)
	}
	return nil
}

func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	log.Println("core.serveHTTP")
	ctx := NewContext(request, response)

	handlers := c.FindRouteByRequest(request)
	if len(handlers) == 0 {
		ctx.Json(http.StatusNotFound, "not found")
		return
	}

	ctx.SetHandlers(handlers)

	if err := ctx.Next(); err != nil {
		ctx.Json(http.StatusInternalServerError, "inner error")
		return
	}
}

func (c *Core) Group(prefix string, middlewares ...ControllerHandler) IGroup {
	parentMiddlewares := append(c.middlewares, middlewares...)
	return NewGroup(c, prefix, parentMiddlewares...)
}

func (c *Core) Use(middlewares ...ControllerHandler) {
	c.middlewares = middlewares
}
