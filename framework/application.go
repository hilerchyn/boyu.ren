package framework

type Application struct {
	Router *Router
}

func (a *Application) Run() {

	a.Router.Exec()

}
