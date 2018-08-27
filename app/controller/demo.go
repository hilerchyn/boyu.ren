package controller

import "net/http"

type Demo struct {
}

func (c *Demo) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello"))

}
