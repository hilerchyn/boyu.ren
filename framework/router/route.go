package router

type Route struct {
	Method     string
	Path       string
	Middleware []func()
	Handler    RouteInterface
}
