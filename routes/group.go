package routes

type RouterGroup struct {
	prefix      string       // 前缀
	parent      *RouterGroup // 父级
	engine      *Engine
	middlewares []HandlerFunc
}

func (this *RouterGroup) Group(prefix string) *RouterGroup {
	engine := this.engine

	routerGroup := &RouterGroup{
		prefix: this.prefix + prefix,
		parent: this,
		engine: engine,
	}

	engine.groups = append(engine.groups, routerGroup)

	return routerGroup
}

// 追加中间件
func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}

func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	group.engine.addRouter(method, pattern, handler)
}

// GET defines the method to add GET request
func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST request
func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}
