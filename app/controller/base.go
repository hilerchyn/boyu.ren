package controller

import (
	"github.com/microcosm-cc/bluemonday"
	"gopkg.in/russross/blackfriday.v2"
	"io/ioutil"
	"net/http"
	"os"
)

type base struct {
}

func (b *base) ResponseErr(w http.ResponseWriter, err error) {

	w.Write([]byte(err.Error()))
}

func (b *base) ResponseMarkdown(w http.ResponseWriter, data []byte) {

	unsafe := blackfriday.Run(data)
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
	w.Write(html)
}

func (b *base) ParseToMarkdown(data []byte) []byte {
	unsafe := blackfriday.Run(data)
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)

	return html
}

func (b *base) GetFileContent(path string) (result []byte, err error) {

	file, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return
	}
	defer file.Close()

	result, err = ioutil.ReadAll(file)
	return
}
