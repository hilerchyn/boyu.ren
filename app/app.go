package app

import "github.com/hilerchyn/boyu.ren/framework"

type BoYu struct {
	Banner string
}

func NewAapp() *BoYu {
	app := new(BoYu)
	app.Banner = framework.Banner()

	return app
}

func (by *BoYu) Start() {
	println(by.Banner)
}
