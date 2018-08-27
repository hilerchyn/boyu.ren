package framework

import (
	"github.com/containous/traefik/log"
	"github.com/hilerchyn/boyu.ren/framework/memstore"
	"github.com/hilerchyn/boyu.ren/framework/router"
	"net/http"
)

type Router struct {
	routes []*router.Route
	store  []*memstore.Store
}

func (r *Router) Register(route *router.Route) {

	if _, ok := route.Handler.(http.Handler); !ok {
		log.Fatal("incorrect handler")
	}

	r.routes = append(r.routes, route)
}

func (r *Router) Exec() {
	for _, v := range r.routes {

		// use middleware
		println("use middleware")

		http.Handle(v.Path, v.Handler.Action(v.Middleware))
	}
}

func newRouter() *Router {
	return &Router{
		routes: []*router.Route{},
	}
}
