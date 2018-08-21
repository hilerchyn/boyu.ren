package context

import (
	"github.com/hilerchyn/boyu.ren/framework/memstore"
	"net/http"
)

//  +------------------------------------------------------------+
//  | Context Implementation                                     |
//  +------------------------------------------------------------+

type context struct {
	// the unique id, it's zero until `String` function is called,
	// it's here to cache the random, unique context's id, although `String`
	// returns more than this.
	id uint64
	// the http.ResponseWriter wrapped by custom writer.
	writer ResponseWriter
	// the original http.Request
	request *http.Request
	// the current route's name registered to this request path.
	currentRouteName string

	// the local key-value storage
	params RequestParams  // url named parameters.
	values memstore.Store // generic storage, middleware communication.

	// the underline application app.
	app Application
	// the route's handlers
	handlers Handlers
	// the current position of the handler's chain
	currentHandlerIndex int
}

// NewContext returns the default, internal, context implementation.
// You may use this function to embed the default context implementation
// to a custom one.
//
// This context is received by the context pool.
func NewContext(app Application) Context {
	return &context{app: app}
}

var _ Context = (*context)(nil)
