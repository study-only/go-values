# Values

A wrapper around string & map[string]string to provide some strong typing and concurrency safety.

## Install
```bash
    go get github.com/study-only/values
```

## Usage

### Value

```go
    import "github.com/study-only/values"

    var v values.Value

    // default value
    v = values.NewValue("")
    d := v.Default("happy").String()

    // is empty
    v = values.Value("")
    isEmpty := v.IsEmpty()

    // value conversion
    v = values.Value("1")
    s := v.String()
    b, err := v.Bool()
    f32, err := v.Float32()
    f64, err := v.Float64()
    i, err := v.Int()
    u, err := v.Uint()
    i8, err := v.Int8()
    u8, err := v.Uint8()
    i16, err := v.Int16()
    u16, err := v.Uint16()
    i32, err := v.Int32()
    u32, err := v.Uint32()
    i64, err := v.Int64()
    u64, err := v.Uint64()
```

### Strict
```go
    import "github.com/study-only/values"

    var v Strict

    // is nil
    v = NewStrict(nil)
    b := v.IsNil()

    // default value
    v = NewStrict(nil)
    i := v.IfNil("abc").Interface()

    // retrieve value
    s, ok := v.MustString()
    b, ok := v.MustBool()
    f32, ok := v.MustFloat32()
    f64, ok := v.MustFloat64()
    i, ok := v.MustInt()
    u, ok := v.MustUint()
    i8, ok := v.MustInt8()
    ui8, ok := v.MustUint8()
    i16, ok := v.MustInt16()
    u16, ok := v.MustUint16()
    i32, ok := v.MustInt32()
    u32, ok := v.MustUint32()
    i64, ok := v.MustInt64()
    u64, ok := v.MustUint64()
```
### Values

```go
    import "github.com/study-only/values"

    // create from map
    m := map[string]Value{"foo": "bar"}
    vs := values.FromMap(m)

    // convert to map
    vs.ToMap()

    // json
    bytes, _ :=json.Marshal(vs)
    json.Unmarshal([]byte(`{"foo":"bar","int":1,"bool":true}`), &vs)

    // get value
    value, exists := vs.Get("foo")

    // set value
    vs.Set("earth", "moon")
    
    // delete value
    vs.Delete("foo")

    // set values
    m = map[string]Value{"bar": "foo", "moon":"earth"}
    vs.Sets(m)
    
    // range
    vs.Range(func(key string, value Value) bool {
        fmt.Println("%s: %s", key, value)
    	return true
    })
```
