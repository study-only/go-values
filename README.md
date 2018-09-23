# Values

Simple string value that can convert to bool, int, float, etc. easily.

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
    b := v.Bool()
    f32 := v.Float32()
    f64 := v.Float64()
    i := v.Int()
    u := v.Uint()
    i8 := v.Int8()
    u8 := v.Uint8()
    i16 := v.Int16()
    u16 := v.Uint16()
    i32 := v.Int32()
    u32 := v.Uint32()
    i64 := v.Int64()
    u64 := v.Uint64()
```

### Values

```go
    import "github.com/liamylian/values"

    // create from map
    m := map[string]string{"foo": "bar"}
    vs := values.FromMap(m)

    // convert to map
    vs.ToMap()

    // json
    bytes, _ :=json.Marshal(vs)
    json.Unmarshal([]byte(`{"foo":"bar","int":1,"bool":true}`), &vs)

    // check key exists
    vs.Defined("foo")

    // get value
    vs.Get("foo")

    // set value
    vs.Set("earth", "moon")

    // set values
    m = map[string]string{"bar": "foo", "moon":"earth"}
    vs.Sets(m)

    // check contains
    vs.Contains("foo", "bar")
```
