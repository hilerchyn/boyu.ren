package router

import (
	"github.com/hilerchyn/boyu.ren/framework/middleware"
	"net/http"
)

type RouteInterface interface {
	Action(middleware.MiddlewareArr) http.Handler
}
