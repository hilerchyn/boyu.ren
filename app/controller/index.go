package controller

import (
	"context"
	"fmt"
	"github.com/hilerchyn/boyu.ren/framework"
	"strings"
)

type Index struct {
	base
}

func (c *Index) GetContent(app *framework.Application, path string) (result []byte, err error) {

	if !strings.HasSuffix(path, ".md") {
		path = "/index.md"
	}
	path = "./markdown" + path

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

	fmt.Println(cd.Request.URL.Path)

	data, err := c.GetContent(cd.Application, cd.Request.URL.Path)
	if err != nil {
		c.ResponseErr(cd.Writer, err)
	}

	c.ResponseMarkdown(cd.Writer, data)

}
