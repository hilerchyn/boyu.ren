package controller

import (
	"context"
	"github.com/hilerchyn/boyu.ren/framework"
)

func init() {
	GPathMap["/demo"] = &Demo{}
}

type Demo struct {
	base
}

func (c *Demo) Action(ctx context.Context) {

	data := ctx.Value("data").(framework.ContextData)

	data.Writer.Header().Set("content-type", "text/html; charset=utf-8")
	templ.ExecuteTemplate(data.Writer, "layout.tmpl", nil)

}
