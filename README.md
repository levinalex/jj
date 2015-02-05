# jj

tiny library that makes dealing with JSON structure assertions in Go a little bit easier


```go
var data jj.Value
str := `{ "foo": "bar", "baz": 3, "bar": { "sub": "val", "int": 4 }}`
err := json.Unmarshal([]byte(str), &data)

assert.Equal(t, "val", data.At("bar", "sub").String())
```

