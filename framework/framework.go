package framework

func NewApp() *Application {

	app := &Application{
		Router: newRouter(),
	}

	return app

}
