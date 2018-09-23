package values

import (
	"testing"
)

func TestValue_IsEmpty(t *testing.T) {
	if !Value("").IsEmpty() {
		t.Error("not empty for empty string value")
	}

	if Value("0").IsEmpty() {
		t.Error("empty for Value(0)")
	}
}

func TestValue_Default(t *testing.T) {
	if Value("").Default("what").String() != "what" {
		t.Error("wrong default value")
	}
}

func TestValue_String(t *testing.T) {
	if Value("hello").String() != "hello" {
		t.Error("wrong string")
	}
}

func TestValue_Int(t *testing.T) {
	if Value("-2018").Int() != -2018 {
		t.Error("wrong int")
	}
}

func TestValue_Uint(t *testing.T) {
	if Value("1988").Uint() != 1988 {
		t.Error("wrong int")
	}
}

func TestValue_Int8(t *testing.T) {
	if Value("-100").Int8() != -100 {
		t.Error("wrong int64")
	}
}

func TestValue_Uint8(t *testing.T) {
	if Value("100").Uint8() != 100 {
		t.Error("wrong int64")
	}
}

func TestValue_Int16(t *testing.T) {
	if Value("-2018").Int16() != -2018 {
		t.Error("wrong int64")
	}
}

func TestValue_Uint16(t *testing.T) {
	if Value("1988").Uint16() != 1988 {
		t.Error("wrong int64")
	}
}

func TestValue_Int32(t *testing.T) {
	if Value("-2018").Int32() != -2018 {
		t.Error("wrong int64")
	}
}

func TestValue_Uint32(t *testing.T) {
	if Value("1988").Uint32() != 1988 {
		t.Error("wrong int64")
	}
}

func TestValue_Int64(t *testing.T) {
	if Value("-2018").Int64() != -2018 {
		t.Error("wrong int64")
	}
}

func TestValue_Uint64(t *testing.T) {
	if Value("1988").Uint64() != 1988 {
		t.Error("wrong int64")
	}
}

func TestValue_Bool(t *testing.T) {
	if Value("").Bool() {
		t.Error("true for empty value")
	}

	if !Value("1").Bool() {
		t.Error("wrong bool")
	}
	if !Value("t").Bool() {
		t.Error("wrong bool")
	}
	if !Value("T").Bool() {
		t.Error("wrong bool")
	}
	if !Value("true").Bool() {
		t.Error("wrong bool")
	}
	if !Value("TRUE").Bool() {
		t.Error("wrong bool")
	}
	if !Value("True").Bool() {
		t.Error("wrong bool")
	}
}

func TestValue_Float32(t *testing.T) {
	if Value("3.1415926").Float32() != 3.1415926 {
		t.Error("wrong float")
	}
}

func TestValue_Float64(t *testing.T) {
	if Value("-3.1415926").Float64() != -3.1415926 {
		t.Error("wrong float")
	}
}
