package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	// origin objects
	Writer http.ResponseWriter
	Req    *http.Request
	// request info
	Path   string
	Method string
	// response info
	StatusCode int
	Params     map[string]string
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Writer:     w,
		Req:        r,
		Path:       r.URL.Path,
		Method:     r.Method,
		StatusCode: 0,
	}
}

func (this *Context) PostForm(key string) string {
	return this.Req.FormValue(key)
}

func (this *Context) Query(key string) string {
	return this.Req.URL.Query().Get(key)
}

func (this *Context) Status(code int) {
	this.StatusCode = code
	this.Writer.WriteHeader(code)
}

func (this *Context) SetHeader(key, val string) {
	this.Writer.Header().Set(key, val)
}

func (this *Context) String(code int, format string, values ...interface{}) {
	str := fmt.Sprintf(format, values...)

	this.Status(code)
	this.SetHeader("Content-Type", "text/plain")
	this.Writer.Write([]byte(str))
}

func (this *Context) Json(code int, data interface{}) {
	this.Status(code)
	this.SetHeader("Content-Type", "application/json")
	encoder := json.NewEncoder(this.Writer)

	if err := encoder.Encode(data); err != nil {
		http.Error(this.Writer, err.Error(), 500)
	}
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}

func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}