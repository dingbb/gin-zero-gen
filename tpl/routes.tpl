// Code generated by goctl. DO NOT EDIT.
package {{.pkgName}}

import (
    "{{.rootPkg}}/{{.handlePkg}}"
    {{if .hasMiddleware}}"{{.rootPkg}}/middleware"{{end}}

	"github.com/gin-gonic/gin"
)

func Register{{.funcName}}Route(e *gin.Engine) {
    g := e.Group("{{if .hasPrefix}}{{.prefix}}{{end}}"){{if .hasMiddleware}}
    g.Use({{range .middleware}}middleware.{{.}}, {{end}}){{end}}
    {{- range .routes}}
    g.{{.method}}("{{.path}}", {{$.handleBase}}.{{.handle}}Handle)
    {{- end}}
}