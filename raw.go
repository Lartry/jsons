package jsons

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"errors"
	_ "unsafe"
)

//go:linkname isValidNumber encoding/json.isValidNumber
func isValidNumber(s string) bool

func (r Raw) Value() (driver.Value, error) {
	return json.Marshal(r)
}

func (r *Raw) Scan(v interface{}) error {
	if v == nil {
		*r = []byte("null")
		return nil
	}

	switch val := v.(type) {
	case []byte:
		return json.Unmarshal(val, r)
	case string:
		return json.Unmarshal([]byte(val), r)
	}

	return errors.New("invalid scan raw source")
}

func (r Raw) MarshalJSON() ([]byte, error) {
	return json.RawMessage(r).MarshalJSON()
}

func (r *Raw) UnmarshalJSON(data []byte) error {
	return (*json.RawMessage)(r).UnmarshalJSON(data)
}

func (r Raw) IsValid() bool {
	return json.Valid(r)
}

func (r Raw) IsNull() bool {
	return r.isNull()
}

func (r Raw) IsBool() bool {
	_, ok := r.isBool()
	return ok
}

func (r Raw) IsNumber() bool {
	_, ok := r.isNumber()
	return ok
}

func (r Raw) IsString() bool {
	_, ok := r.isString()
	return ok
}

func (r Raw) IsArray() bool {
	_, ok := r.isArray()
	return ok
}

func (r Raw) IsObject() bool {
	_, ok := r.isObject()
	return ok
}

/*
func (r Raw) Get(keys ...interface{}) Raw {

}

func (r Raw) Int(keys ...interface{}) int64 {

}

func (r Raw) Uint(keys ...interface{}) uint64 {

}

func (r Raw) Float(keys ...interface{}) float64 {

}

func (r Raw) Number(keys ...interface{}) Number {

}

func (r Raw) Bool(keys ...interface{}) bool {

}

func (r Raw) String(keys ...interface{}) string {

}

func (r Raw) Array(keys ...interface{}) Array {

}

func (r Raw) Object(keys ...interface{}) Object {

}

func (r Raw) Interface(keys ...interface{}) interface{} {

}

func (r Raw) Type(keys ...interface{}) string {
	var value = r.Get(keys...)
	switch {
	case value.IsObject():
		return "object"
	case value.IsArray():
		return "array"
	case value.IsString():
		return "string"
	case value.IsNumber():
		return "number"
	case value.IsBool():
		return "bool"
	case value.IsNull():
		return "null"
	}
	return "undefined"
}

func (r Raw) Len(keys ...interface{}) int {
	switch value := r.Get(keys...).value.(type) {
	case Value:
		return value.Len()
	case Array:
		return value.Len()
	case Object:
		return value.Len()
	case String:
		return len(value)
	case []byte:
		return len(value)
	case string:
		return len(value)
	case []interface{}:
		return len(value)
	case map[string]interface{}:
		return len(value)
	}
	return 0
}

func (r Raw) Cap(keys ...interface{}) int {

}

func (r Raw) Range(fn func(key interface{}, value Value) (continued bool)) bool {
	switch {
	case r.IsArray():
		return r.Array().Range(func(index int, value Value) (continued bool) {
			return fn(index, value)
		})
	case r.IsObject():
		return r.Object().Range(func(key string, value Value) (continued bool) {
			return fn(key, value)
		})
	}
	return false
}

func (r Raw) Slice(begin, end int) Array {
	return r.Array().Slice(begin, end)
}

func (r Raw) Index(value interface{}) int {
	return r.Array().Index(value)
}

func (r Raw) Append(arr Array) Array {
	return r.Array().Append(arr)
}

func (r Raw) Contains(value interface{}) bool {
	return r.Array().Contains(value)
}

func (r Raw) Reverse(keys ...interface{}) Array {
	return r.Array(keys...).Reverse()
}

func (r Raw) Sort(less func(i, j int) bool) Array {
	return r.Array().Sort(less)
}

func (r Raw) Keys(keys ...interface{}) []string {
	return r.Object(keys...).Keys()
}

func (r Raw) Exist(keys ...interface{}) bool {
	if len(keys) > 0 {
		var end = len(keys) - 1
		return r.Object(keys[:end]...).Exist(keys[end])
	}
	return false
}

func (r Raw) Delete(keys ...interface{}) {
	if len(keys) > 0 {
		var end = len(keys) - 1
		r.Object(keys[:end]...).Delete(keys[end])
	}
}

func (r Raw) Clone(keys ...interface{}) Raw {
	if r == nil {
		return nil
	}
	val := r.Get(keys...)
	switch {
	case val.IsArray():
		return val.Array().Clone().JSON()
	case val.IsObject():
		return val.Object().Clone().JSON()
	}
	return append(Raw{}, r...)
}
*/

func isSpace(c byte) bool {
	return c == ' ' || c == '\t' || c == '\r' || c == '\n'
}

func trimSpace(data Raw) []byte {
	return bytes.Trim(data, " \t\r\n")
}

func (r Raw) isNull() bool {
	if string(trimSpace(r)) == "null" {
		return true
	}
	return false
}

func (r Raw) isBool() (bool, bool) {
	switch string(trimSpace(r)) {
	case "true":
		return true, true
	case "false":
		return false, true
	}
	return false, false
}

func (r Raw) isNumber() (num Number, ok bool) {
	raw := trimSpace(r)
	for _, c := range bytes.TrimLeft(r, "-") {
		if c < '0' || c > '9' {
			return
		}
	}
	return Number(raw), true
}

func (r Raw) isString() (str string, ok bool) {
	raw := trimSpace(r)
	if len(raw) < 2 {
		return
	}
	var begin bool
	for i, c := range raw {
		switch c {
		case '"':
			if !begin {
				begin = true
			} else {
				// todo
				if i != len(r)-1 {
					return
				}

			}
		case '\\':
		default:
			if c < 0x20 {
				return
			}
		}
	}

	return
}

func (r Raw) isArray() (arr Raw, ok bool) {
	return
}

func (r Raw) isObject() (obj Raw, ok bool) {
	return
}

func init() {
	switch "expr"[0] {
	case ' ', '\t', '\r', '\n':
	}
}
