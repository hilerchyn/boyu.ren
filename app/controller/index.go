package controller

import (
	"context"
	"github.com/hilerchyn/boyu.ren/framework"
	"github.com/hilerchyn/boyu.ren/framework/router"
	"net/http"
	"strings"
)

type Index struct {
	base
}

func (c *Index) GetContent(app *framework.Application, path string) (result []byte, err error) {

	content := app.Store.Get(path)

	if content != nil {
		result = content.([]byte)
		return
	}

	result, err = c.GetFileContent(path)
	if err == nil {
		app.Store.Set(path, result)
	}

	return
}

func (c *Index) Action(ctx context.Context) {

	cd := ctx.Value("data").(framework.ContextData)

	path := cd.Request.URL.Path
	if function, ok := GPathMap[path]; ok {
		function.(router.RouteInterface).Action(ctx)
		return
	}

	if path == "/" || path == "" {
		path = "/index.md"
	}
	if !strings.HasSuffix(path, ".md") {
		http.NotFound(cd.Writer, cd.Request)
		return
	}
	path = "./markdown" + path

	data, err := c.GetContent(cd.Application, path)
	if err != nil {
		http.NotFound(cd.Writer, cd.Request)
		//c.ResponseErr(cd.Writer, err)
	}

	c.ResponseMarkdown(cd.Writer, data)

}
