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
	data.Writer.Write([]byte("from demo"))

}
