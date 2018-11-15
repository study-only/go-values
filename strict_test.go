package values

import "testing"

func TestStrict_IsNil(t *testing.T) {
	if !NewStrict(nil).IsNil() {
		t.Error("not nil for nil")
	}

	var a *int
	if !NewStrict(a).IsNil() {
		t.Error("not nil for((*int)(nil)")
	}

	var b int
	if NewStrict(b).IsNil() {
		t.Error("nil for int(0)")
	}

	var c string
	if NewStrict(c).IsNil() {
		t.Error("nil for empty string")
	}
}

func TestStrict_Default(t *testing.T) {
	s := NewStrict(nil).IfNil("what")
	if str := s.Interface().(string); str != "what" {
		t.Error("wrong default value")
	}
}

func TestStrict_String(t *testing.T) {
	if s, ok := NewStrict("hello").MustString(); !ok || s != "hello" {
		t.Error("wrong string")
	}
}

func TestStrict_Int(t *testing.T) {
	if i, ok := NewStrict(int(-2018)).MustInt(); !ok || i != -2018 {
		t.Error("wrong int")
	}
}

func TestStrict_Uint(t *testing.T) {
	if i, ok := NewStrict(uint(1998)).MustUint(); !ok || i != 1998 {
		t.Error("wrong uint")
	}
}

func TestStrict_Int8(t *testing.T) {
	if i, ok := NewStrict(int8(-100)).MustInt8(); !ok || i != -100 {
		t.Error("wrong int8")
	}
}

func TestStrict_Uint8(t *testing.T) {
	if i, ok := NewStrict(uint8(100)).MustUint8(); !ok || i != 100 {
		t.Error("wrong uint8")
	}
}

func TestStrict_Int16(t *testing.T) {
	if i, ok := NewStrict(int16(-1998)).MustInt16(); !ok || i != -1998 {
		t.Error("wrong int16")
	}
}

func TestStrict_Uint16(t *testing.T) {
	if i, ok := NewStrict(uint16(1998)).MustUint16(); !ok || i != 1998 {
		t.Error("wrong uint16")
	}
}

func TestStrict_Int32(t *testing.T) {
	if i, ok := NewStrict(int32(-1998)).MustInt32(); !ok || i != -1998 {
		t.Error("wrong int32")
	}
}

func TestStrict_Uint32(t *testing.T) {
	if i, ok := NewStrict(uint32(1998)).MustUint32(); !ok || i != 1998 {
		t.Error("wrong uint32")
	}
}

func TestStrict_Int64(t *testing.T) {
	if i, ok := NewStrict(int64(-1998)).MustInt64(); !ok || i != -1998 {
		t.Error("wrong int64")
	}
}

func TestStrict_Uint64(t *testing.T) {
	if i, ok := NewStrict(uint64(1998)).MustUint64(); !ok || i != 1998 {
		t.Error("wrong uint64")
	}
}

func TestStrict_Bool(t *testing.T) {
	if b, ok := NewStrict(true).MustBool(); !ok || b != true {
		t.Error("wrong bool")
	}

}

func TestStrict_Float32(t *testing.T) {
	if f, ok := NewStrict(float32(3.14)).MustFloat32(); !ok || f != 3.14 {
		t.Error("wrong float32")
	}
}

func TestStrict_Float64(t *testing.T) {
	if f, ok := NewStrict(float64(3.14)).MustFloat64(); !ok || f != 3.14 {
		t.Error("wrong float64")
	}
}