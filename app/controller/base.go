package controller

import (
	"fmt"
	"github.com/microcosm-cc/bluemonday"
	"gopkg.in/russross/blackfriday.v2"
	"net/http"
)

type base struct {
}

func (b *base) ResponseErr(w http.ResponseWriter, err error) {

	w.Write([]byte(err.Error()))
}

func (b *base) ResponseMarkdown(w http.ResponseWriter, data []byte) {

	unsafe := blackfriday.Run(data)
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)

	fmt.Println(string(html))
	w.Write(html)
}
