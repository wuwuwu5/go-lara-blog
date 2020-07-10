package routes

import (
	"net/http"
	"strings"
)

type HandlerFunc func(c *Context)

type Engine struct {
	router  map[string]*node
	handler map[string]HandlerFunc
	*RouterGroup
	groups []*RouterGroup
}

func InitRouter() *Engine {
	engine := &Engine{
		router:  make(map[string]*node),
		handler: make(map[string]HandlerFunc),
	}

	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}

	return engine
}

// 解析URL
func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

// 添加路由
func (this *Engine) addRouter(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern

	parts := parsePattern(pattern)

	if _, ok := this.router[method]; !ok {
		this.router[method] = &node{}
	}

	this.router[method].insert(pattern, parts, 0)
	this.handler[key] = handler
}

// 获取路由
func (this *Engine) getRouter(method string, pattern string) (*node, map[string]string) {
	if root, ok := this.router[method]; ok {
		params := make(map[string]string)

		parts := parsePattern(pattern)

		result := root.search(parts, 0)

		if result != nil {
			resultParts := parsePattern(result.pattern)

			for index, part := range resultParts {
				if part[0] == ':' {
					params[part[1:]] = parts[index]
				}

				if part[0] == '*' && len(part) > 1 {
					params[part[1:]] = strings.Join(parts[index:], "/")
					break
				}
			}
		}

		return result, params
	}

	return nil, nil
}

// GET请求
func (this *Engine) Get(pattern string, handler HandlerFunc) {
	this.addRouter("GET", pattern, handler)
}

// 路由匹配
func (this *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	handlers := make([]HandlerFunc, 0)

	for _, group := range this.groups {
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			handlers = append(handlers, group.middlewares...)
		}
	}

	node, params := this.getRouter(req.Method, req.URL.Path)

	c := NewContext(w, req)
	c.Params = params

	if node != nil {
		key := req.Method + "-" + node.pattern
		if handler, ok := this.handler[key]; ok {
			handlers = append(handlers, handler)
		} else {
			handlers = append(handlers, func(c *Context) {
				c.String(404, "NOT FOUND")
			})
		}
	} else {
		handlers = append(handlers, func(c *Context) {
			c.String(404, "NOT FOUND")
		})
	}

	c.handlers = handlers

	c.Next()
}

// 启动
func (this *Engine) Run(addr string) error {
	this.Get("/static/*filepath", this.createStaticHandler(http.Dir("/Users/wujian/go/src/lara-blog/public")))
	return http.ListenAndServe(addr, this)
}

// create static handler
func (this *Engine) createStaticHandler(fs http.FileSystem) HandlerFunc {
	fileServer := http.StripPrefix("/static", http.FileServer(fs))
	return func(c *Context) {
		file := c.Param("filepath")
		if _, err := fs.Open(file); err != nil {
			c.Status(http.StatusNotFound)
			return
		}

		fileServer.ServeHTTP(c.Writer, c.Req)
	}
}
