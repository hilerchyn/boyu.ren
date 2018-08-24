package framework

import "github.com/hilerchyn/boyu.ren/framework/router"

type Router struct {
	routes []*router.Route
}

func (r *Router) Register(route *router.Route) {
	r.routes = append(r.routes, route)
}

func (r *Router) Exec() {
	for _, v := range r.routes {
		println(v.Path)
	}
}

func newRouter() *Router {
	return &Router{
		routes: []*router.Route{},
	}
}
