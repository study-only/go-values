package values

import (
	"reflect"
)

type Strict struct {
	val interface{}
}

func NewStrict(val interface{}) Strict {
	return Strict{val}
}

func (v Strict) IsNil() bool {
	return isNil(v.val)
}

func (v Strict) IfNil(val interface{}) Strict {
	if isNil(v.val) {
		return Strict{val}
	}

	return v
}

func (v Strict) Interface() interface{} {
	return v.val
}

func (v Strict) MustString() (str string, ok bool) {
	str, ok = v.val.(string)
	return
}

func (v Strict) MustInt() (i int, ok bool) {
	i, ok = v.val.(int)
	return
}

func (v Strict) MustUint() (u uint, ok bool) {
	u, ok = v.val.(uint)
	return
}

func (v Strict) MustInt8() (i int8, ok bool) {
	i, ok = v.val.(int8)
	return
}

func (v Strict) MustUint8() (u uint8, ok bool) {
	u, ok = v.val.(uint8)
	return
}

func (v Strict) MustInt16() (i int16, ok bool) {
	i, ok = v.val.(int16)
	return
}

func (v Strict) MustUint16() (u uint16, ok bool) {
	u, ok = v.val.(uint16)
	return
}

func (v Strict) MustInt32() (i int32, ok bool) {
	i, ok = v.val.(int32)
	return
}

func (v Strict) MustUint32() (u uint32, ok bool) {
	u, ok = v.val.(uint32)
	return
}

func (v Strict) MustInt64() (i int64, ok bool) {
	i, ok = v.val.(int64)
	return
}

func (v Strict) MustUint64() (u uint64, ok bool) {
	u, ok = v.val.(uint64)
	return
}

func (v Strict) MustBool() (b bool, ok bool) {
	b, ok = v.val.(bool)
	return
}

func (v Strict) MustFloat32() (f float32, ok bool) {
	f, ok = v.val.(float32)
	return
}

func (v Strict) MustFloat64() (f float64, ok bool) {
	f, ok = v.val.(float64)
	return
}

func isNil(i interface{}) bool {
	if i == nil {
		return true
	}

	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	return false
}
