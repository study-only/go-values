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

func (v Value) Int() (int, error) {
	return strconv.Atoi(string(v))
}

func (v Value) Uint() (uint, error) {
	num, err := strconv.ParseUint(string(v), 10, 0)

	return uint(num), err
}

func (v Value) Int8() (int8, error) {
	num, err := strconv.ParseInt(string(v), 10, 8)

	return int8(num), err
}

func (v Value) Uint8() (uint8, error) {
	num, err := strconv.ParseUint(string(v), 10, 8)

	return uint8(num), err
}

func (v Value) Int16() (int16, error) {
	num, err := strconv.ParseInt(string(v), 10, 16)

	return int16(num), err
}

func (v Value) Uint16() (uint16, error) {
	num, err := strconv.ParseUint(string(v), 10, 16)

	return uint16(num), err
}

func (v Value) Int32() (int32, error) {
	num, err := strconv.ParseInt(string(v), 10, 32)

	return int32(num), err
}

func (v Value) Uint32() (uint32, error) {
	num, err := strconv.ParseUint(string(v), 10, 32)

	return uint32(num), err
}

func (v Value) Int64() (int64, error) {
	return strconv.ParseInt(string(v), 10, 64)
}

func (v Value) Uint64() (uint64, error) {
	return strconv.ParseUint(string(v), 10, 64)
}

func (v Value) Bool() (bool, error) {
	return strconv.ParseBool(string(v))
}

func (v Value) Float32() (float32, error) {
	num, err := strconv.ParseFloat(string(v), 32)

	return float32(num), err
}

func (v Value) Float64() (float64, error) {
	return strconv.ParseFloat(string(v), 64)
}
