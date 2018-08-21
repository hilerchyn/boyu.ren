package context

import "net/http"

// ResponseWriter interface is used by the context to serve an HTTP handler to
// construct an HTTP response.
//
// Note: Only this ResponseWriter is an interface in order to be able
// for developers to change the response writer of the Context via `context.ResetResponseWriter`.
// The rest of the response writers implementations (ResponseRecorder & GzipResponseWriter) are coupled to the internal
// ResponseWriter implementation(*responseWriter).
//
// A ResponseWriter may not be used after the Handler
// has returned.
type ResponseWriter interface {
	http.ResponseWriter
	http.Flusher
	http.Hijacker
	http.CloseNotifier
	http.Pusher

	// Naive returns the simple, underline and original http.ResponseWriter
	// that backends this response writer.
	Naive() http.ResponseWriter

	// BeginResponse receives an http.ResponseWriter
	// and initialize or reset the response writer's field's values.
	BeginResponse(http.ResponseWriter)

	// EndResponse is the last function which is called right before the server sent the final response.
	//
	// Here is the place which we can make the last checks or do a cleanup.
	EndResponse()

	// Writef formats according to a format specifier and writes to the response.
	//
	// Returns the number of bytes written and any write error encountered.
	Writef(format string, a ...interface{}) (n int, err error)

	// WriteString writes a simple string to the response.
	//
	// Returns the number of bytes written and any write error encountered.
	WriteString(s string) (n int, err error)

	// StatusCode returns the status code header value.
	StatusCode() int

	// Written should returns the total length of bytes that were being written to the client.
	// In addition iris provides some variables to help low-level actions:
	// NoWritten, means that nothing were written yet and the response writer is still live.
	// StatusCodeWritten, means that status code were written but no other bytes are written to the client, response writer may closed.
	// > 0 means that the reply was written and it's the total number of bytes were written.
	Written() int

	// SetWritten sets manually a value for written, it can be
	// NoWritten(-1) or StatusCodeWritten(0), > 0 means body length which is useless here.
	SetWritten(int)

	// SetBeforeFlush registers the unique callback which called exactly before the response is flushed to the client.
	SetBeforeFlush(cb func())
	// GetBeforeFlush returns (not execute) the before flush callback, or nil if not setted by SetBeforeFlush.
	GetBeforeFlush() func()
	// FlushResponse should be called only once before EndResponse.
	// it tries to send the status code if not sent already
	// and calls the  before flush callback, if any.
	//
	// FlushResponse can be called before EndResponse, but it should
	// be the last call of this response writer.
	FlushResponse()

	// clone returns a clone of this response writer
	// it copies the header, status code, headers and the beforeFlush finally  returns a new ResponseRecorder.
	Clone() ResponseWriter

	// WiteTo writes a response writer (temp: status code, headers and body) to another response writer
	WriteTo(ResponseWriter)

	// Flusher indicates if `Flush` is supported by the client.
	//
	// The default HTTP/1.x and HTTP/2 ResponseWriter implementations
	// support Flusher, but ResponseWriter wrappers may not. Handlers
	// should always test for this ability at runtime.
	//
	// Note that even for ResponseWriters that support Flush,
	// if the client is connected through an HTTP proxy,
	// the buffered data may not reach the client until the response
	// completes.
	Flusher() (http.Flusher, bool)

	// CloseNotifier indicates if the protocol supports the underline connection closure notification.
	CloseNotifier() (http.CloseNotifier, bool)
}

// ResponseWriter is the basic response writer,
// it writes directly to the underline http.ResponseWriter
type responseWriter struct {
	http.ResponseWriter
	statusCode int // the saved status code which will be used from the cache service
	// statusCodeSent bool // reply header has been (logically) written | no needed any more as we have a variable to catch total len of written bytes
	written int // the total size of bytes were written
	// yes only one callback, we need simplicity here because on FireStatusCode the beforeFlush events should NOT be cleared
	// but the response is cleared.
	// Sometimes is useful to keep the event,
	// so we keep one func only and let the user decide when he/she wants to override it with an empty func before the FireStatusCode (context's behavior)
	beforeFlush func()
}
