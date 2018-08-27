package controller

import (
	"github.com/hilerchyn/boyu.ren/framework/middleware"
	"net/http"
)

type Index struct {
	base
}

func (c *Index) GetContent() (result []byte, err error) {
	result, err = c.GetFileContent("./markdown/index.md")
	return
}

func (c *Index) Action(m middleware.MiddlewareArr) http.Handler {

	if len(m) > 0 {

		for i := 0; i < len(m); i++ {
			function := m[i]

			function()
		}

	}

	return c
}

func (c *Index) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	data, err := c.GetContent()
	if err != nil {
		c.ResponseErr(w, err)
	}

	c.ResponseMarkdown(w, data)

}
