package framework

import "fmt"

type IGroup interface {
	Group(string, ...ControllerHandler) IGroup
	Use(...ControllerHandler)
	Get(string, ...ControllerHandler)
	Put(string, ...ControllerHandler)
	Post(string, ...ControllerHandler)
	Delete(string, ...ControllerHandler)
}

type Group struct {
	c           *Core
	prefix      string
	middlewares []ControllerHandler
}

func NewGroup(c *Core, prefix string, middlewares ...ControllerHandler) *Group {
	return &Group{
		c:           c,
		prefix:      prefix,
		middlewares: middlewares,
	}
}

func (g *Group) Get(uri string, handlers ...ControllerHandler) {
	uri = g.prefix + uri
	allHandlers := append(g.middlewares, handlers...)
	fmt.Println(len(allHandlers))
	g.c.Get(uri, allHandlers...)
}

func (g *Group) Put(uri string, handlers ...ControllerHandler) {
	uri = g.prefix + uri
	allHandlers := append(g.middlewares, handlers...)
	g.c.Put(uri, allHandlers...)
}

func (g *Group) Post(uri string, handlers ...ControllerHandler) {
	allHandlers := append(g.middlewares, handlers...)
	uri = g.prefix + uri
	g.c.Post(uri, allHandlers...)

}

func (g *Group) Delete(uri string, handlers ...ControllerHandler) {
	uri = g.prefix + uri
	allHandlers := append(g.middlewares, handlers...)
	g.c.Delete(uri, allHandlers...)
}

func (g *Group) Group(uri string, middlewares ...ControllerHandler) IGroup {
	parentMiddlewares := append(g.middlewares, middlewares...)
	return NewGroup(g.c, g.prefix+uri, parentMiddlewares...)
}

func (g *Group) Use(middlewares ...ControllerHandler) {
	g.middlewares = append(g.middlewares, middlewares...)
}
