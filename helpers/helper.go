package helpers

import (
	"encoding/base64"
	"encoding/json"
	"html/template"
	"lara-blog/pkg"
	"lara-blog/routes"
	"log"
	"os"
	"path"
	"strings"
)

var (
	// 存储 mix-manifest.json 解析出来的 path map
	manifests = make(map[string]string)
)

// Static 生成项目静态文件地址
func Static(staticFilePath string) string {
	return "/static" + staticFilePath
}

// Mix 根据 laravel-mix 生成静态文件 path
func Mix(staticFilePath string) string {
	result := manifests[staticFilePath]

	if result == "" {
		filename := path.Join("./public", "mix-manifest.json")
		file, err := os.Open(filename)
		if err != nil {
			return "静态文件读取失败"
		}

		defer file.Close()

		dec := json.NewDecoder(file)

		if err := dec.Decode(&manifests); err != nil {
			return Static(staticFilePath)
		}

		result = manifests[staticFilePath]
	}

	if result == "" {
		return Static(staticFilePath)
	}

	return Static(result)
}

// 返回页面
func View(c *routes.Context, path string, data map[string]interface{}) {
	tmpl, err := template.New("app.html").Funcs(template.FuncMap{"Mix": Mix}).ParseFiles(
		"./views/layout/app.html",
		"./views/layout/_header.html",
		"./views/layout/_footer.html",
		path,
	)

	if err != nil {
		log.Fatal(err)
	}

	cookie, err := c.Req.Cookie(pkg.ValidateCookieName)

	if err == nil {
		messages, err := base64.StdEncoding.DecodeString(cookie.Value)

		if err == nil {
			messageStr := string(messages)
			messageArr := strings.Split(messageStr, "|")
			data["validateMessage"] = messageArr
			pkg.DelValidateMessages(c.Writer)
		}
	}

	tmpl.Execute(c.Writer, data)

}
