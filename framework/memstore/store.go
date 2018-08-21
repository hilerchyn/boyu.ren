package memstore

import "strings"

type (
	Store []Entry
)

// Save same as `Set`
// However, if "immutable" is true then saves it as immutable (same as `SetImmutable`).
//
//
// Returns the entry and true if it was just inserted, meaning that
// it will return the entry and a false boolean if the entry exists and it has been updated.
func (r *Store) Save(key string, value interface{}, immutable bool) (Entry, bool) {
	args := *r
	n := len(args)

	// replace if we can, else just return
	for i := 0; i < n; i++ {
		kv := &args[i]
		if kv.Key == key {
			if immutable && kv.immutable {
				// if called by `SetImmutable`
				// then allow the update, maybe it's a slice that user wants to update by SetImmutable method,
				// we should allow this
				kv.ValueRaw = value
				kv.immutable = immutable
			} else if kv.immutable == false {
				// if it was not immutable then user can alt it via `Set` and `SetImmutable`
				kv.ValueRaw = value
				kv.immutable = immutable
			}
			// else it was immutable and called by `Set` then disallow the update
			return *kv, false
		}
	}

	// expand slice to add it
	c := cap(args)
	if c > n {
		args = args[:n+1]
		kv := &args[n]
		kv.Key = key
		kv.ValueRaw = value
		kv.immutable = immutable
		*r = args
		return *kv, true
	}

	// add
	kv := Entry{
		Key:       key,
		ValueRaw:  value,
		immutable: immutable,
	}
	*r = append(args, kv)
	return kv, true
}

// Set saves a value to the key-value storage.
// Returns the entry and true if it was just inserted, meaning that
// it will return the entry and a false boolean if the entry exists and it has been updated.
//
// See `SetImmutable` and `Get`.
func (r *Store) Set(key string, value interface{}) (Entry, bool) {
	return r.Save(key, value, false)
}

// SetImmutable saves a value to the key-value storage.
// Unlike `Set`, the output value cannot be changed by the caller later on (when .Get OR .Set)
//
// An Immutable entry should be only changed with a `SetImmutable`, simple `Set` will not work
// if the entry was immutable, for your own safety.
//
// Returns the entry and true if it was just inserted, meaning that
// it will return the entry and a false boolean if the entry exists and it has been updated.
//
// Use it consistently, it's far slower than `Set`.
// Read more about muttable and immutable go types: https://stackoverflow.com/a/8021081
func (r *Store) SetImmutable(key string, value interface{}) (Entry, bool) {
	return r.Save(key, value, true)
}

// GetEntry returns a pointer to the "Entry" found with the given "key"
// if nothing found then it returns nil, so be careful with that,
// it's not supposed to be used by end-developers.
func (r *Store) GetEntry(key string) *Entry {
	args := *r
	n := len(args)
	for i := 0; i < n; i++ {
		kv := &args[i]
		if kv.Key == key {
			return kv
		}
	}

	return nil
}

// GetDefault returns the entry's value based on its key.
// If not found returns "def".
// This function checks for immutability as well, the rest don't.
func (r *Store) GetDefault(key string, def interface{}) interface{} {
	v := r.GetEntry(key)
	if v == nil || v.ValueRaw == nil {
		return def
	}
	vv := v.Value()
	if vv == nil {
		return def
	}
	return vv
}

// Get returns the entry's value based on its key.
// If not found returns nil.
func (r *Store) Get(key string) interface{} {
	return r.GetDefault(key, nil)
}

// Visit accepts a visitor which will be filled
// by the key-value objects.
func (r *Store) Visit(visitor func(key string, value interface{})) {
	args := *r
	for i, n := 0, len(args); i < n; i++ {
		kv := args[i]
		visitor(kv.Key, kv.Value())
	}
}

// GetStringDefault returns the entry's value as string, based on its key.
// If not found returns "def".
func (r *Store) GetStringDefault(key string, def string) string {
	v := r.GetEntry(key)
	if v == nil {
		return def
	}

	return v.StringDefault(def)
}

// GetString returns the entry's value as string, based on its key.
func (r *Store) GetString(key string) string {
	return r.GetStringDefault(key, "")
}

// GetStringTrim returns the entry's string value without trailing spaces.
func (r *Store) GetStringTrim(name string) string {
	return strings.TrimSpace(r.GetString(name))
}

// GetInt returns the entry's value as int, based on its key.
// If not found returns -1 and a non-nil error.
func (r *Store) GetInt(key string) (int, error) {
	v := r.GetEntry(key)
	if v == nil {
		return 0, errFindParse.Format("int", key)
	}
	return v.IntDefault(-1)
}

// GetIntDefault returns the entry's value as int, based on its key.
// If not found returns "def".
func (r *Store) GetIntDefault(key string, def int) int {
	if v, err := r.GetInt(key); err == nil {
		return v
	}

	return def
}

// GetInt64 returns the entry's value as int64, based on its key.
// If not found returns -1 and a non-nil error.
func (r *Store) GetInt64(key string) (int64, error) {
	v := r.GetEntry(key)
	if v == nil {
		return -1, errFindParse.Format("int64", key)
	}
	return v.Int64Default(-1)
}

// GetInt64Default returns the entry's value as int64, based on its key.
// If not found returns "def".
func (r *Store) GetInt64Default(key string, def int64) int64 {
	if v, err := r.GetInt64(key); err == nil {
		return v
	}

	return def
}

// GetFloat64 returns the entry's value as float64, based on its key.
// If not found returns -1 and a non nil error.
func (r *Store) GetFloat64(key string) (float64, error) {
	v := r.GetEntry(key)
	if v == nil {
		return -1, errFindParse.Format("float64", key)
	}
	return v.Float64Default(-1)
}

// GetFloat64Default returns the entry's value as float64, based on its key.
// If not found returns "def".
func (r *Store) GetFloat64Default(key string, def float64) float64 {
	if v, err := r.GetFloat64(key); err == nil {
		return v
	}

	return def
}

// GetBool returns the user's value as bool, based on its key.
// a string which is "1" or "t" or "T" or "TRUE" or "true" or "True"
// or "0" or "f" or "F" or "FALSE" or "false" or "False".
// Any other value returns an error.
//
// If not found returns false and a non-nil error.
func (r *Store) GetBool(key string) (bool, error) {
	v := r.GetEntry(key)
	if v == nil {
		return false, errFindParse.Format("bool", key)
	}

	return v.BoolDefault(false)
}

// GetBoolDefault returns the user's value as bool, based on its key.
// a string which is "1" or "t" or "T" or "TRUE" or "true" or "True"
// or "0" or "f" or "F" or "FALSE" or "false" or "False".
//
// If not found returns "def".
func (r *Store) GetBoolDefault(key string, def bool) bool {
	if v, err := r.GetBool(key); err == nil {
		return v
	}

	return def
}

// Remove deletes an entry linked to that "key",
// returns true if an entry is actually removed.
func (r *Store) Remove(key string) bool {
	args := *r
	n := len(args)
	for i := 0; i < n; i++ {
		kv := &args[i]
		if kv.Key == key {
			// we found the index,
			// let's remove the item by appending to the temp and
			// after set the pointer of the slice to this temp args
			args = append(args[:i], args[i+1:]...)
			*r = args
			return true
		}
	}
	return false
}

// Reset clears all the request entries.
func (r *Store) Reset() {
	*r = (*r)[0:0]
}

// Len returns the full length of the entries.
func (r *Store) Len() int {
	args := *r
	return len(args)
}

// Serialize returns the byte representation of the current Store.
func (r Store) Serialize() []byte { // note: no pointer here, ignore linters if shows up.
	b, _ := GobSerialize(r)
	return b
}
