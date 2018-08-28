package app

import (
	"github.com/hilerchyn/boyu.ren/framework"
	"html/template"
)

type BoYu struct {
	Banner   string
	App      *framework.Application
	Template *template.Template
}

func NewAapp() *BoYu {
	app := new(BoYu)
	app.Banner = framework.Banner()
	app.App = framework.NewApp()

	return app
}

func (by *BoYu) Start() {

	println(by.Banner)

	// register route
	by.RegisterRoute()
	by.App.Run()

}
