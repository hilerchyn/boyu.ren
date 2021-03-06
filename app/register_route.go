package app

import (
	"github.com/hilerchyn/boyu.ren/app/controller"
	"github.com/hilerchyn/boyu.ren/framework/router"
	"net/http"
)

func (by *BoYu) RegisterRoute() {

	// /
	route := &router.Route{}
	route.Method = http.MethodGet
	route.Path = "/"
	route.Handler = &controller.Index{}

	by.App.Router.Register(route)

}
