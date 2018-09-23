package values

import (
	"strconv"
)

type Value string

func (v Value) IsEmpty() bool {
	return v == ""
}

func (v Value) Default(value string) Value {
	if v == "" {
		return Value(value)
	}
	return v
}

func (v Value) String() string {
	return string(v)
}

func (v Value) Int() int {
	num, _ := strconv.Atoi(string(v))

	return num
}

func (v Value) Uint() uint {
	num, _ := strconv.ParseUint(string(v), 10, 0)

	return uint(num)
}

func (v Value) Int8() int8 {
	num, _ := strconv.ParseInt(string(v), 10, 8)

	return int8(num)
}

func (v Value) Uint8() uint8 {
	num, _ := strconv.ParseUint(string(v), 10, 8)

	return uint8(num)
}

func (v Value) Int16() int16 {
	num, _ := strconv.ParseInt(string(v), 10, 16)

	return int16(num)
}

func (v Value) Uint16() uint16 {
	num, _ := strconv.ParseUint(string(v), 10, 16)

	return uint16(num)
}

func (v Value) Int32() int32 {
	num, _ := strconv.ParseInt(string(v), 10, 32)

	return int32(num)
}

func (v Value) Uint32() uint32 {
	num, _ := strconv.ParseUint(string(v), 10, 32)

	return uint32(num)
}

func (v Value) Int64() int64 {
	num, _ := strconv.ParseInt(string(v), 10, 64)

	return num
}

func (v Value) Uint64() uint64 {
	num, _ := strconv.ParseUint(string(v), 10, 64)

	return num
}

func (v Value) Bool() bool {
	b, _ := strconv.ParseBool(string(v))

	return b
}

func (v Value) Float32() float32 {
	num, _ := strconv.ParseFloat(string(v), 32)

	return float32(num)
}

func (v Value) Float64() float64 {
	num, _ := strconv.ParseFloat(string(v), 64)

	return num
}
