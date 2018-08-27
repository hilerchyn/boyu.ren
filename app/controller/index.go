package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Index struct {
	base
}

func (c *Index) GetContent() (result []byte, err error) {
	file, err := os.OpenFile("./markdown/index.md", os.O_RDONLY, os.ModePerm)
	if err != nil {
		return
	}
	defer file.Close()

	result, err = ioutil.ReadAll(file)

	fmt.Println(string(result))

	return

}

func (c *Index) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	data, err := c.GetContent()
	if err != nil {
		c.ResponseErr(w, err)
	}

	c.ResponseMarkdown(w, data)

}
