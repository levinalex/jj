# jj

tiny library that makes dealing with JSON structure assertions in Go a little bit easier

[![Build Status](https://travis-ci.org/levinalex/jj.svg?branch=master)](https://travis-ci.org/levinalex/jj)


```go
var data jj.Value
str := `{ "foo": "bar", "baz": 3, "bar": { "sub": "val", "int": 4 }}`
err := json.Unmarshal([]byte(str), &data)

assert.Equal(t, "val", data.At("bar", "sub").String())
```

