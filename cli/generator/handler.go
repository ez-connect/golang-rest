package generator

import (
	"fmt"
	"strings"
)

var find = `func (h *Handler) Find%s(c echo.Context) error {
	f := filter.Find(c, &Model{})
	o := filter.Option(c)
	docs := []Model{}
	return h.Find(c, f, o, nil, &docs)
}
`

var insert = `func (h *Handler) Insert%s(c echo.Context) error {
	doc := Model{
		CreatedAt: core.Now(),
		UpdatedAt: core.Now(),
	}
	return h.Insert(c, &doc)
}
`

var findOne = `func (h *Handler) FindOne%s(c echo.Context) error {
	f := filter.FindOne(c, &Model{})
	if f == nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	doc := Model{}
	return h.FindOne(c, f, nil, &doc)
}
`

var update = `func (h *Handler) Update%s(c echo.Context) error {
	f := filter.FindOne(c, &Model{})
	if f == nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	doc := Model{
		UpdatedAt: core.Now(),
	}
	return h.UpdateOne(c, f, &doc)
}
`

var delete = `func (h *Handler) Delete%s(c echo.Context) error {
	f := filter.FindOne(c, &Model{})
	if f == nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	return h.DeleteOne(c, f)
}
`

func GenerateBaseHandler() string {
	buf := []string{}
	buf = append(buf, "package _base\n")

	buf = append(buf, "import (")
	buf = append(buf, "\t\"github.com/ez-connect/go-rest/core\"")
	buf = append(buf, ")\n")

	buf = append(buf, "type HandlerBase struct {")
	buf = append(buf, "\trest.HandlerBase")
	buf = append(buf, "}\n")

	return strings.Join(buf, "\n")
}

func GenerateHandler(packageName string) string {
	buf := []string{}
	buf = append(buf, fmt.Sprintf("package %s\n", packageName))

	buf = append(buf, "import (")
	buf = append(buf, "\t\"net/http\"\n")

	buf = append(buf, "\t\"github.com/ez-connect/go-rest/core\"")
	buf = append(buf, "\t\"github.com/ez-connect/go-rest/rest/filter\"")
	buf = append(buf, "\t\"github.com/ez-connect/go-rest/rest\"")
	buf = append(buf, "\t\"github.com/labstack/echo/v4\"\n")
	// buf = append(buf, fmt.Sprintf("\t\"app/services/%s\"", packageName))
	buf = append(buf, ")\n")

	buf = append(buf, "type Handler struct {")
	// buf = append(buf, fmt.Sprintf("\t%s.Handler", packageName))
	buf = append(buf, "\trest.HandlerBase")
	buf = append(buf, "\tRepo Repository")
	buf = append(buf, "}\n")

	buf = append(buf, "///////////////////////////////////////////////////////////////////\n")
	buf = append(buf, fmt.Sprintf(find, strings.Title(packageName)))
	buf = append(buf, fmt.Sprintf(insert, strings.Title(packageName)))
	buf = append(buf, fmt.Sprintf(findOne, strings.Title(packageName)))
	buf = append(buf, fmt.Sprintf(update, strings.Title(packageName)))
	buf = append(buf, fmt.Sprintf(delete, strings.Title(packageName)))

	return strings.Join(buf, "\n")
}

func GenerateHandlerExt(packageName string) string {
	buf := []string{}
	buf = append(buf, fmt.Sprintf("package %s\n", packageName))

	buf = append(buf, "import (")
	buf = append(buf, "\t\"app/services/_base\"")
	buf = append(buf, fmt.Sprintf("\t\"app/generated/%s\"", packageName))
	buf = append(buf, ")\n")

	buf = append(buf, "type Handler struct {")
	buf = append(buf, fmt.Sprintf("\t%s.Handler", packageName))
	buf = append(buf, "\t_base.HandlerBase")
	buf = append(buf, "\tRepo Repository")
	buf = append(buf, "}\n")

	return strings.Join(buf, "\n")
}
