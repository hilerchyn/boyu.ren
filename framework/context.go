package framework

import (
	"github.com/hilerchyn/boyu.ren/framework/router"
	"net/http"
)

type ContextData struct {
	Writer      http.ResponseWriter
	Request     *http.Request
	Route       router.Route
	Application *Application
}
