package controller

import (
	"github.com/hilerchyn/boyu.ren/framework/middleware"
	"net/http"
)

type Demo struct {
	base
}

func (c *Demo) Action(m middleware.MiddlewareArr) http.Handler {
	return c
}

func (c *Demo) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello"))

}
