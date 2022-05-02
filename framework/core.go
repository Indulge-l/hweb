package framework

import (
	"log"
	"net/http"
	"strings"
)

type Core struct {
	router map[string]*Tree
}

func NewCore() *Core {
	router := map[string]*Tree{}
	router[http.MethodGet] = NewTree()
	router[http.MethodPut] = NewTree()
	router[http.MethodPost] = NewTree()
	router[http.MethodDelete] = NewTree()
	return &Core{router: router}
}

func (c *Core) Get(url string, handler ControllerHandler) {
	if err := c.router[http.MethodGet].AddRouter(url, handler); err != nil {
		log.Fatalf("add router error:(%+v)", err)
	}
}

func (c *Core) Put(url string, handler ControllerHandler) {
	if err := c.router[http.MethodPut].AddRouter(url, handler); err != nil {
		log.Fatalf("add router error:(%+v)", err)
	}
}

func (c *Core) Post(url string, handler ControllerHandler) {
	if err := c.router[http.MethodPost].AddRouter(url, handler); err != nil {
		log.Fatalf("add router error:(%+v)", err)
	}
}

func (c *Core) Delete(url string, handler ControllerHandler) {
	if err := c.router[http.MethodDelete].AddRouter(url, handler); err != nil {
		log.Fatalf("add router error:(%+v)", err)
	}
}

func (c *Core) FindRouteByRequest(req *http.Request) ControllerHandler {
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

	router := c.FindRouteByRequest(request)
	if router == nil {
		ctx.Json(http.StatusNotFound, "not found")
		return
	}
	if err := router(ctx); err != nil {
		ctx.Json(http.StatusInternalServerError, "inner error")
		return
	}
}

func (c *Core) Group(prefix string) IGroup {
	return NewGroup(c, prefix)
}
