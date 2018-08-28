package framework

import (
	"context"
	"github.com/hilerchyn/boyu.ren/framework/router"
	"net/http"
)

type Router struct {
	routes map[string]*router.Route
}

func (r *Router) Register(route *router.Route) {

	r.routes[route.Path] = route
}

func (r *Router) Exec(app *Application) {
	for _, v := range r.routes {

		function := func(w http.ResponseWriter, req *http.Request) {
			data := ContextData{
				Writer:      w,
				Request:     req,
				Route:       *v,
				Application: app,
			}

			ctx := context.WithValue(context.Background(), "data", data)
			v.Handler.Action(ctx)
		}

		http.HandleFunc(v.Path, function)
	}

}

func newRouter() *Router {
	return &Router{
		routes: map[string]*router.Route{},
	}
}
