package router

import (
	"context"
)

type RouteInterface interface {
	Action(context.Context)
}
