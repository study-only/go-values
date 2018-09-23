package values

import (
	"testing"
	"encoding/json"
	"sync"
)

func Test_FromMap(t *testing.T) {
	m := map[string]string{"foo": "bar", "bar": "foo"}
	values := FromMap(m)

	if v := values.Get("foo"); v != "bar" {
		t.Errorf("expected bar but got, %s", v)
	}
	if v := values.Get("bar"); v != "foo" {
		t.Errorf("expected foo but got, %s", v)
	}
}

func TestValues_Defined(t *testing.T) {
	m := new(sync.Map)
	m.Store("foo", "bar")
	values := Values{
		values: m,
	}

	if !values.Defined("foo") {
		t.Errorf("got undefined for defiend")
	}

	if values.Defined("puppy") {
		t.Errorf("got defined for undefiend")
	}
}

func TestValues_Match(t *testing.T) {
	m := new(sync.Map)
	m.Store("foo", "bar")
	m.Store("earth", "moon")
	values := Values{
		values: m,
	}

	if !values.Contains("foo", "bar") {
		t.Errorf("unexpected unmatch")
	}
	if !values.Contains("earth", "moon") {
		t.Errorf("unexpected unmatch")
	}
	if values.Contains("foo", "moon") {
		t.Errorf("unexpected match")
	}
	if values.Contains("earth", "human") {
		t.Errorf("unexpected match")
	}
	if values.Contains("mars", "jupiter") {
		t.Errorf("unexpected match")
	}
}

func TestValues_Get(t *testing.T) {
	m := new(sync.Map)
	m.Store("foo", "bar")
	values := Values{
		values: m,
	}

	if v := values.Get("foo"); v != "bar" {
		t.Errorf("expected bar but got, %s", v)
	}
}

func TestValues_Set(t *testing.T) {
	values := Values{}
	values.Set("foo", "bar")

	if v := values.Get("foo"); v != "bar" {
		t.Errorf("expected bar but got, %s", v)
	}
}

func TestValues_Sets(t *testing.T) {
	m := map[string]string{"foo": "bar", "bar": "foo"}
	values := Values{}

	values.Sets(m)
	if v := values.Get("foo"); v != "bar" {
		t.Errorf("expected bar but got, %s", v)
	}
	if v := values.Get("bar"); v != "foo" {
		t.Errorf("expected foo but got, %s", v)
	}

	m2 := map[string]string{"earth": "moon"}
	values.Sets(m2)
	if v := values.Get("foo"); v != "bar" {
		t.Errorf("expected bar but got, %s", v)
	}
	if v := values.Get("bar"); v != "foo" {
		t.Errorf("expected foo but got, %s", v)
	}
	if v := values.Get("earth"); v != "moon" {
		t.Errorf("expected moon but got, %s", v)
	}
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
	if v := values.Get("foo"); v != "bar" {
		t.Errorf("expected bar, but got %s", v)
	}
	if v := values.Get("int").Int(); v != -2018 {
		t.Errorf("expected -2018, but got %d", v)
	}
	if v := values.Get("float").Float64(); v != 3.1415926 {
		t.Errorf("expected 3.1415926, but got %f", v)
	}
	if v := values.Get("bool").Bool(); v != true {
		t.Errorf("expected true, but got %v", v)
	}
	if v := values.Get("null").String(); v != "" {
		t.Errorf("expected empty, but got %v", v)
	}
}
