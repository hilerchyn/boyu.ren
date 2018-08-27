package framework

import (
	"github.com/containous/traefik/log"
	"github.com/hilerchyn/boyu.ren/framework/memstore"
	"net/http"
)

type Application struct {
	Router *Router
	Store  *memstore.Store
}

func (a *Application) Run() {

	a.Router.Exec()

	err := http.ListenAndServe(":8443", nil) //http.ListenAndServeTLS(":8443", )
	log.Fatal(err)

}
