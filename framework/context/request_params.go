package context

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hilerchyn/boyu.ren/framework/memstore"
)

// RequestParams is a key string - value string storage which
// context's request dynamic path params are being kept.
// Empty if the route is static.
type RequestParams struct {
	store memstore.Store
}

// Set adds a key-value pair to the path parameters values
// it's being called internally so it shouldn't be used as a local storage by the user, use `ctx.Values()` instead.
func (r *RequestParams) Set(key, value string) {
	r.store.Set(key, value)
}

// Visit accepts a visitor which will be filled
// by the key-value params.
func (r *RequestParams) Visit(visitor func(key string, value string)) {
	r.store.Visit(func(k string, v interface{}) {
		visitor(k, v.(string)) // always string here.
	})
}

var emptyEntry memstore.Entry

// GetEntryAt returns the internal Entry of the memstore based on its index,
// the stored index by the router.
// If not found then it returns a zero Entry and false.
func (r RequestParams) GetEntryAt(index int) (memstore.Entry, bool) {
	if len(r.store) > index {
		return r.store[index], true
	}
	return emptyEntry, false
}

// GetEntry returns the internal Entry of the memstore based on its "key".
// If not found then it returns a zero Entry and false.
func (r RequestParams) GetEntry(key string) (memstore.Entry, bool) {
	// we don't return the pointer here, we don't want to give the end-developer
	// the strength to change the entry that way.
	if e := r.store.GetEntry(key); e != nil {
		return *e, true
	}
	return emptyEntry, false
}

// Get returns a path parameter's value based on its route's dynamic path key.
func (r RequestParams) Get(key string) string {
	return r.store.GetString(key)
}

// GetTrim returns a path parameter's value without trailing spaces based on its route's dynamic path key.
func (r RequestParams) GetTrim(key string) string {
	return strings.TrimSpace(r.Get(key))
}

// GetEscape returns a path parameter's double-url-query-escaped value based on its route's dynamic path key.
func (r RequestParams) GetEscape(key string) string {
	return DecodeQuery(DecodeQuery(r.Get(key)))
}

// GetDecoded returns a path parameter's double-url-query-escaped value based on its route's dynamic path key.
// same as `GetEscape`.
func (r RequestParams) GetDecoded(key string) string {
	return r.GetEscape(key)
}

// GetInt returns the path parameter's value as int, based on its key.
func (r RequestParams) GetInt(key string) (int, error) {
	return r.store.GetInt(key)
}

// GetInt64 returns the path paramete's value as int64, based on its key.
func (r RequestParams) GetInt64(key string) (int64, error) {
	return r.store.GetInt64(key)
}

// GetFloat64 returns a path parameter's value based as float64 on its route's dynamic path key.
func (r RequestParams) GetFloat64(key string) (float64, error) {
	return r.store.GetFloat64(key)
}

// GetBool returns the path parameter's value as bool, based on its key.
// a string which is "1" or "t" or "T" or "TRUE" or "true" or "True"
// or "0" or "f" or "F" or "FALSE" or "false" or "False".
// Any other value returns an error.
func (r RequestParams) GetBool(key string) (bool, error) {
	return r.store.GetBool(key)
}

// GetIntUnslashed same as Get but it removes the first slash if found.
// Usage: Get an id from a wildcard path.
//
// Returns -1 with an error if the parameter couldn't be found.
func (r RequestParams) GetIntUnslashed(key string) (int, error) {
	v := r.Get(key)
	if v != "" {
		if len(v) > 1 {
			if v[0] == '/' {
				v = v[1:]
			}
		}
		return strconv.Atoi(v)

	}

	return -1, fmt.Errorf("unable to find int for '%s'", key)
}

// Len returns the full length of the parameters.
func (r RequestParams) Len() int {
	return r.store.Len()
}
