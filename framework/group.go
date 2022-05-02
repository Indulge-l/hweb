package framework

type IGroup interface {
	//Group(string) IGroup
	Get(string, ...ControllerHandler)
	Put(string, ...ControllerHandler)
	Post(string, ...ControllerHandler)
	Delete(string, ...ControllerHandler)
}

type Group struct {
	c      *Core
	prefix string
}

func NewGroup(c *Core, prefix string) *Group {
	return &Group{
		c:      c,
		prefix: prefix,
	}
}

func (g *Group) Get(uri string, handlers ...ControllerHandler) {
	uri = g.prefix + uri
	g.c.Get(uri, handlers[0])
}

func (g *Group) Put(uri string, handlers ...ControllerHandler) {
	uri = g.prefix + uri
	g.c.Put(uri, handlers[0])
}

func (g *Group) Post(uri string, handlers ...ControllerHandler) {
	uri = g.prefix + uri
	g.c.Post(uri, handlers[0])

}

func (g *Group) Delete(uri string, handlers ...ControllerHandler) {
	uri = g.prefix + uri
	g.c.Delete(uri, handlers[0])
}
