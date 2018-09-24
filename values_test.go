package values

import (
	"testing"
	"encoding/json"
	"sync"
)

func Test_FromMap(t *testing.T) {
	m := map[string]Value{"foo": "bar", "bar": "foo"}
	values := FromMap(m)

	if v, _ := values.Get("foo"); v != "bar" {
		t.Errorf("expected bar, but got %s", v)
	}
	if v, _ := values.Get("bar"); v != "foo" {
		t.Errorf("expected foo, but got %s", v)
	}
}

func TestValues_Get(t *testing.T) {
	m := sync.Map{}
	m.Store("foo", Value("bar"))
	values := Values{
		values: m,
	}

	if v, found := values.Get("foo"); !found && v != "bar" {
		t.Errorf("expected bar, but got %s", v)
	}
}

func TestValues_Set(t *testing.T) {
	values := Values{}
	values.Set("foo", "bar")

	if v, _ := values.Get("foo"); v != "bar" {
		t.Errorf("expected bar, but got %s", v)
	}
}

func TestValues_Sets(t *testing.T) {
	m := map[string]Value{"foo": "bar", "bar": "foo"}
	values := Values{}

	values.Sets(m)
	if v, _ := values.Get("foo"); v != "bar" {
		t.Errorf("expected bar, but got %s", v)
	}
	if v, _ := values.Get("bar"); v != "foo" {
		t.Errorf("expected foo, but got %s", v)
	}

	m2 := map[string]Value{"earth": "moon"}
	values.Sets(m2)
	if v, _ := values.Get("foo"); v != "bar" {
		t.Errorf("expected bar, but got %s", v)
	}
	if v, _ := values.Get("bar"); v != "foo" {
		t.Errorf("expected foo, but got %s", v)
	}
	if v, _ := values.Get("earth"); v != "moon" {
		t.Errorf("expected moon, but got %s", v)
	}
}

func TestValues_Delete(t *testing.T) {
	m := sync.Map{}
	m.Store("foo", Value("bar"))
	values := Values{
		values: m,
	}

	values.Delete("foo")
	if _, exist := values.Get("foo"); exist {
		t.Errorf("unexpected exists")
	}
}

func TestValues_ToMap(t *testing.T) {
	m := sync.Map{}
	m.Store("foo", Value("bar"))
	values := Values{
		values: m,
	}

	vm := values.ToMap()
	if vm["foo"] != "bar"{
		t.Errorf("expected bar, but got %s", vm["foo"])
	}
}

func TestValues_Range(t *testing.T) {
	m := sync.Map{}
	m.Store("foo", Value("bar"))
	m.Store("earth", Value("moon"))
	values := Values{
		values: m,
	}

	values.Range(func(key string, value Value) bool {
		if key != "foo" && key != "earth" {
			t.Errorf("upecxtected key: %s", key)
		}

		if key == "foo" && value != "bar" {
			t.Errorf("upecxtected key value: %s, %s", key, value)
		}
		if key == "earth" && value != "moon" {
			t.Errorf("upecxtected key value: %s, %s", key, value)
		}

		return true
	})
}

func TestValues_MarshalJSON(t *testing.T) {
	values := Values{}
	values.Set("foo", "bar")
	values.Set("bar", "foo")

	bytes, err := json.Marshal(values)
	if err != nil {
		t.Errorf("marshal error: %s", err.Error())
	}

	if string(bytes) != `{"foo":"bar","bar":"foo"}` && string(bytes) != `{"bar":"foo","foo":"bar"}` {
		t.Errorf("not as expected, got: %s", bytes)
	}
}

func TestValues_UnmarshalJSON(t *testing.T) {
	bytes := `{"foo":"bar","int":-2018,"float":3.1415926,"bool":true,"null":null}`
	values := Values{}

	if err := json.Unmarshal([]byte(bytes), &values); err != nil {
		t.Errorf("unmarshal error: %s", err.Error())
	}
	if v, _ := values.Get("foo"); v != "bar" {
		t.Errorf("expected bar, but got %s", v)
	}
	if v, _ := values.Get("int"); v.Int() != -2018 {
		t.Errorf("expected -2018, but got %d", v.Int())
	}
	if v, _ := values.Get("float"); v.Float64() != 3.1415926 {
		t.Errorf("expected 3.1415926, but got %f", v.Float64())
	}
	if v, _ := values.Get("bool"); v.Bool() != true {
		t.Errorf("expected true, but got %v", v.Bool())
	}
	if v, _ := values.Get("null"); v.String() != "" {
		t.Errorf("expected empty, but got %v", v)
	}
}
