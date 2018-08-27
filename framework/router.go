package framework

import (
	"github.com/hilerchyn/boyu.ren/framework/router"
	"net/http"
)

type Router struct {
	routes []*router.Route
}

func (r *Router) Register(route *router.Route) {
	r.routes = append(r.routes, route)
}

func (r *Router) Exec() {
	for _, v := range r.routes {
		println(v.Path)

		// use middleware
		println("use middleware")

		http.Handle(v.Path, v.Handler)
	}
}

func newRouter() *Router {
	return &Router{
		routes: []*router.Route{},
	}
}
