# Values

A wrapper around string & map[string]string to provide some strong typing and concurrency safety.

## Install
```bash
    go get github.com/liamylian/values
```

## Usage

### Value

```go
    import "github.com/liamylian/values"

    var v values.Value

    // default value
    v = values.Value("")
    d := v.Default("happy").String()

    // is empty
    v = values.Value("")
    isEmpty := v.IsEmpty()

    // value conversion
    v = values.Value("1")
    s := v.String()
    b, _ := v.Bool()
    f32, _ := v.Float32()
    f64, _ := v.Float64()
    i, _ := v.Int()
    u, _ := v.Uint()
    i8, _ := v.Int8()
    u8, _ := v.Uint8()
    i16, _ := v.Int16()
    u16, _ := v.Uint16()
    i32, _ := v.Int32()
    u32, _ := v.Uint32()
    i64, _ := v.Int64()
    u64, _ := v.Uint64()
```

### Values

```go
    import "github.com/liamylian/values"

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
