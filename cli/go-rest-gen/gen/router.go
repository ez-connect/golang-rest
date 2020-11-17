package gen

import (
	"fmt"
	"strings"
)

var initRoute = `
func (r *Router) Init(e *echo.Echo, db db.DatabaseBase) {
	r.Router.Init(e, &r.Handler.Handler)
	r.Handler.Init(db, %s.CollectionName, &r.Repo.Repository)
}
`

func GenerateRoutes(packageName string, config Config) string {
	buf := []string{}
	buf = append(buf, fmt.Sprintf("package %s\n", packageName))

	buf = append(buf, "import (")
	buf = append(buf, "\t\"github.com/labstack/echo/v4\"\n")
	buf = append(buf, "\t\"github.com/ez-connect/go-rest/rest\"")

	// buf = append(buf, fmt.Sprintf("\t\"app/services/%s\"", packageName))
	for _, v := range config.Import.Router {
		buf = append(buf, fmt.Sprintf("\t\"%s\"", v))
	}

	buf = append(buf, ")\n")

	buf = append(buf, "type Router struct {")
	// buf = append(buf, fmt.Sprintf("\t%s.Router", packageName))
	buf = append(buf, "\trest.RouterBase")
	buf = append(buf, "}\n")

	buf = append(buf, "func (r *Router) Init(e *echo.Echo, h *Handler) {")

	for i, v := range config.Routes {
		buf = append(buf, fmt.Sprintf("\tg%v := e.Group(\"%s\")", i, v.Path))
		if v.MiddlewareFunc != "" {
			buf = append(buf, fmt.Sprintf("\tg%v.Use(rest.JWTWithAuthHandler(%s))", i, v.MiddlewareFunc))
		}
		for _, r := range v.Children {
			buf = append(buf,
				fmt.Sprintf("\tg%v.%s(\"%s\", h.%s)", i, r.Method, r.Path, r.Handler),
			)
		}
	}

	buf = append(buf, "}\n")

	return strings.Join(buf, "\n")
}

func GenerateRoutesService(packageName string) string {
	buf := []string{}
	buf = append(buf, fmt.Sprintf("package %s\n", packageName))

	buf = append(buf, "import (")
	buf = append(buf, fmt.Sprintf("\t\"app/generated/%s\"\n", packageName))
	buf = append(buf, "\t\"github.com/ez-connect/go-rest/db\"")
	buf = append(buf, "\t\"github.com/labstack/echo/v4\"")
	buf = append(buf, ")\n")

	buf = append(buf, "type Router struct {")
	buf = append(buf, fmt.Sprintf("\t%s.Router", packageName))
	buf = append(buf, "\tHandler")
	buf = append(buf, "}")

	buf = append(buf, fmt.Sprintf(initRoute, packageName))

	return strings.Join(buf, "\n")
}
