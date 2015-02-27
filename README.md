# jj

Read arbitrary and deeply nested JSON structures in Go with as little boilerplate code as possible.

[![Build Status](https://travis-ci.org/levinalex/jj.svg?branch=master)](https://travis-ci.org/levinalex/jj)
[![GoDoc](https://godoc.org/github.com/levinalex/jj?status.svg)](https://godoc.org/github.com/levinalex/jj)


```go
var data jj.Value
str := `{ "foo": "bar", "bar": { "sub": "val", "int": 4 }, "baz": ["a", 9] }`
err := json.Unmarshal([]byte(str), &data)

fmt.Println(data.At("bar", "sub").String()) // panics if the key does not exist
// => "val"

fmt.Println(data.At("x").StringOrDefault("default")) // always succeeds
// => "default"

fmt.Println(data.At("baz", 1).Number())
// => 9

```

