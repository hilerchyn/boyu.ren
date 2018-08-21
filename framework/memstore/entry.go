package memstore

import (
	"github.com/hilerchyn/boyu.ren/framework/common"
	"reflect"
	"strconv"
)

type (
	Entry struct {
		Key       string
		ValueRaw  interface{}
		immutable bool // if true then it can't be changed by its caller
	}
)

var errFindParse = common.NewErrors("unable to find the %s with key: %s")

func (e Entry) GetByKindOrNil(k reflect.Kind) interface{} {
	switch k {
	case reflect.String:
		v := e.StringDefault("__$nf")
		if v == "__$nf" {
			return nil
		}
		return v
	case reflect.Int:
		v, err := e.IntDefault(-1)
		if err != nil || v == -1 {
			return nil
		}
		return v
	case reflect.Int64:
		v, err := e.Int64Default(-1)
		if err != nil || v == -1 {
			return nil
		}
		return v
	case reflect.Bool:
		v, err := e.BoolDefault(false)
		if err != nil {
			return nil
		}
		return v
	default:
		return e.ValueRaw
	}
}

// StringDefault returns the entry's value as string.
// If not found returns "def".
func (e Entry) StringDefault(def string) string {
	v := e.ValueRaw

	if vString, ok := v.(string); ok {
		return vString
	}

	return def
}

// IntDefault returns the entry's value as int.
// If not found returns "def" and a non-nil error.
func (e Entry) IntDefault(def int) (int, error) {
	v := e.ValueRaw
	if v == nil {
		return def, errFindParse.Format("int", e.Key)
	}
	if vint, ok := v.(int); ok {
		return vint, nil
	} else if vstring, sok := v.(string); sok && vstring != "" {
		vint, err := strconv.Atoi(vstring)
		if err != nil {
			return def, err
		}

		return vint, nil
	}

	return def, errFindParse.Format("int", e.Key)
}

// Int64Default returns the entry's value as int64.
// If not found returns "def" and a non-nil error.
func (e Entry) Int64Default(def int64) (int64, error) {
	v := e.ValueRaw
	if v == nil {
		return def, errFindParse.Format("int64", e.Key)
	}

	if vint64, ok := v.(int64); ok {
		return vint64, nil
	}

	if vint, ok := v.(int); ok {
		return int64(vint), nil
	}

	if vstring, sok := v.(string); sok {
		return strconv.ParseInt(vstring, 10, 64)
	}

	return def, errFindParse.Format("int64", e.Key)
}

// Float64Default returns the entry's value as float64.
// If not found returns "def" and a non-nil error.
func (e Entry) Float64Default(def float64) (float64, error) {
	v := e.ValueRaw

	if v == nil {
		return def, errFindParse.Format("float64", e.Key)
	}

	if vfloat32, ok := v.(float32); ok {
		return float64(vfloat32), nil
	}

	if vfloat64, ok := v.(float64); ok {
		return vfloat64, nil
	}

	if vint, ok := v.(int); ok {
		return float64(vint), nil
	}

	if vstring, sok := v.(string); sok {
		vfloat64, err := strconv.ParseFloat(vstring, 64)
		if err != nil {
			return def, err
		}

		return vfloat64, nil
	}

	return def, errFindParse.Format("float64", e.Key)
}

// Float32Default returns the entry's value as float32.
// If not found returns "def" and a non-nil error.
func (e Entry) Float32Default(key string, def float32) (float32, error) {
	v := e.ValueRaw

	if v == nil {
		return def, errFindParse.Format("float32", e.Key)
	}

	if vfloat32, ok := v.(float32); ok {
		return vfloat32, nil
	}

	if vfloat64, ok := v.(float64); ok {
		return float32(vfloat64), nil
	}

	if vint, ok := v.(int); ok {
		return float32(vint), nil
	}

	if vstring, sok := v.(string); sok {
		vfloat32, err := strconv.ParseFloat(vstring, 32)
		if err != nil {
			return def, err
		}

		return float32(vfloat32), nil
	}

	return def, errFindParse.Format("float32", e.Key)
}

// BoolDefault returns the user's value as bool.
// a string which is "1" or "t" or "T" or "TRUE" or "true" or "True"
// or "0" or "f" or "F" or "FALSE" or "false" or "False".
// Any other value returns an error.
//
// If not found returns "def" and a non-nil error.
func (e Entry) BoolDefault(def bool) (bool, error) {
	v := e.ValueRaw
	if v == nil {
		return def, errFindParse.Format("bool", e.Key)
	}

	if vBoolean, ok := v.(bool); ok {
		return vBoolean, nil
	}

	if vString, ok := v.(string); ok {
		b, err := strconv.ParseBool(vString)
		if err != nil {
			return def, err
		}
		return b, nil
	}

	if vInt, ok := v.(int); ok {
		if vInt == 1 {
			return true, nil
		}
		return false, nil
	}

	return def, errFindParse.Format("bool", e.Key)
}

// Value returns the value of the entry,
// respects the immutable.
func (e Entry) Value() interface{} {
	if e.immutable {
		// take its value, no pointer even if setted with a reference.
		vv := reflect.Indirect(reflect.ValueOf(e.ValueRaw))

		// return copy of that slice
		if vv.Type().Kind() == reflect.Slice {
			newSlice := reflect.MakeSlice(vv.Type(), vv.Len(), vv.Cap())
			reflect.Copy(newSlice, vv)
			return newSlice.Interface()
		}
		// return a copy of that map
		if vv.Type().Kind() == reflect.Map {
			newMap := reflect.MakeMap(vv.Type())
			for _, k := range vv.MapKeys() {
				newMap.SetMapIndex(k, vv.MapIndex(k))
			}
			return newMap.Interface()
		}
		// if was *value it will return value{}.
		return vv.Interface()
	}
	return e.ValueRaw
}
