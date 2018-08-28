package controller

import (
	"context"
	"net/http"
)

type Demo struct {
	base
}

func (c *Demo) Action(ctx context.Context) {
}

func (c *Demo) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello"))

}
