package framework

func NewApp() *Application {

	app := &Application{
		Router: newRouter(),
		Store:  newStore(),
	}

	return app

}
