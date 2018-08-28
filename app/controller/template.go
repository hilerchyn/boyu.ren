package controller

import "html/template"

var templ = &template.Template{}

func init() {

	// parse template
	templ = template.Must(template.New("boyu.ren").ParseGlob("templates/*.tmpl"))
}
