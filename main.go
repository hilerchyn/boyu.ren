package main

import (
	"fmt"
	"github.com/microcosm-cc/bluemonday"
	"gopkg.in/russross/blackfriday.v2"
)

func main() {

	unsafe := blackfriday.Run([]byte("*HELLO WORLD*"))
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)

	fmt.Println(string(html))

}
