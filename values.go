package values

import (
	"sync"
	"encoding/json"
	"strconv"
)

type Values struct {
	values sync.Map
}

func FromMap(data map[string]Value) Values {
	values := Values{}
	values.Sets(data)
	return values
}

func (vs Values) Get(key string) (v Value, found bool) {
	raw, exist := vs.values.Load(key)
	if !exist {
		found = false
		return
	}

	v, _ = raw.(Value)
	found = true
	return
}

func (vs *Values) Set(key string, value Value) {
	vs.values.Store(key, value)
}

func (vs *Values) Sets(kvs map[string]Value) {
	for k, v := range kvs {
		vs.values.Store(k, v)
	}
}

func (vs *Values) Delete(key string) {
	vs.values.Delete(key)
}

func (vs Values) Range(f func(key string, value Value) (isContinue bool)) {
	vs.values.Range(func(k, v interface{}) bool {
		if kk, ok := k.(string); ok {
			if vv, ok := v.(Value); ok {
				return f(kk, vv)
			}
		}

		return true
	})
}

func (vs Values) ToMap() map[string]Value {
	data := make(map[string]Value)
	vs.values.Range(func(rawKey, rawValue interface{}) bool {
		if k, ok := rawKey.(string); ok {
			if v, ok := rawValue.(Value); ok {
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
		vs.Set(k, Value(jsonString(v)))
	}

	return nil
}

func (vs Values) MarshalJSON() ([]byte, error) {
	tmp := make(map[string]string)
	vs.values.Range(func(rawKey, rawValue interface{}) bool {
		if k, ok := rawKey.(string); ok {
			if v, ok := rawValue.(Value); ok {
				tmp[k] = v.String()
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
