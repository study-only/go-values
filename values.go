package values

import (
	"sync"
	"encoding/json"
	"strconv"
)

type Values struct {
	values *sync.Map
}

func FromMap(data map[string]string) Values {
	values := Values{}
	values.Sets(data)
	return values
}

func (vs Values) Defined(key string) bool {
	if vs.values == nil {
		return false
	}

	if _, exist := vs.values.Load(key); exist {
		return true
	} else {
		return false
	}
}

func (vs Values) HasMatch(key string, value Value) bool {
	if vs.values == nil {
		return false
	}

	found := false
	vs.values.Range(func(rawKey, rawValue interface{}) bool {
		if k, ok := rawKey.(string); ok && k == key {
			if v, ok := rawValue.(string); ok && v == value.String() {
				found = true
				return false
			} else {
				return false
			}
		}

		return true
	})

	return found
}

func (vs Values) Get(key string) Value {
	if vs.values == nil {
		return Value("")
	}

	raw, exist := vs.values.Load(key)
	if !exist {
		return Value("")
	}

	if v, ok := raw.(string); ok {
		return Value(v)
	} else {
		return Value("")
	}
}

func (vs *Values) Set(key, value string) {
	if vs.values == nil {
		vs.values = new(sync.Map)
	}
	vs.values.Store(key, value)
}

func (vs *Values) Sets(extras map[string]string) {
	if vs.values == nil {
		vs.values = new(sync.Map)
	}
	for k, v := range extras {
		vs.values.Store(k, v)
	}
}

func (vs Values) ToMap() map[string]string {
	data := make(map[string]string)
	if vs.values == nil {
		return data
	}

	vs.values.Range(func(rawKey, rawValue interface{}) bool {
		if k, ok := rawKey.(string); ok {
			if v, ok := rawValue.(string); ok {
				data[k] = v
			}
		}

		return true
	})

	return data
}

func (vs *Values) UnmarshalJSON(data []byte) error {
	tmp := make(map[string]interface{})
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}

	for k, v := range tmp {
		vs.Set(k, jsonString(v))
	}

	return nil
}

func (vs Values) MarshalJSON() ([]byte, error) {
	tmp := make(map[string]string)
	if vs.values == nil {
		return []byte("{}"), nil
	}

	vs.values.Range(func(rawKey, rawValue interface{}) bool {
		if k, ok := rawKey.(string); ok {
			if v, ok := rawValue.(string); ok {
				tmp[k] = v
			}
		}

		return true
	})

	return json.Marshal(tmp)
}

func jsonString(v interface{}) string {
	switch v := v.(type) {
	case bool:
		return strconv.FormatBool(v)
	case int:
		return strconv.Itoa(v)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case string:
		return v
	default:
		return ""
	}
}
