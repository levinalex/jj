package jj

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeserialize(t *testing.T) {
	var data Value
	var err error
	var str string

	str = `{ "foo": "bar", "baz": 3, "bar": { "sub": "val", "int": 4 }}`
	err = json.Unmarshal([]byte(str), &data)
	assert.Nil(t, err)

	assert.Equal(t, "bar", data.At("foo").String())
	assert.Equal(t, "val", data.At("bar", "sub").String())
	assert.Equal(t, "val", data.At("bar").At("sub").String())

	assert.EqualValues(t, 4, data.At("bar", "int").Int())
	assert.EqualValues(t, 4, data.At("bar", "int").Value())

	assert.Equal(t, []string{"bar", "baz", "foo"}, data.KeysSorted())

	str = `{ "foo": [1,2,"bar",{ "baz": "fred" }]}`
	err = json.Unmarshal([]byte(str), &data)
	assert.Nil(t, err)

	assert.Equal(t, "bar", data.At("foo", 2).String())
	assert.Equal(t, "fred", data.At("foo", 3, "baz").String())
	assert.Equal(t, "fred", data.At("foo", 3, "baz").Value())

	assert.EqualValues(t, 1, data.At("foo", 0).Number())
	assert.EqualValues(t, 2, data.At("foo", 1).Number())

	assert.Equal(t, "", data.At("missing").StringOrDefault(""))

	assert.Panics(t, func() {
		data.AtOrError(nil)
	}, "calling AtOrError with something that is neither string not int should panic")

	assert.Nil(t, data.At(4))
	assert.Nil(t, data.At("foo", "bar"))
	assert.Nil(t, data.At("doesnotexist").At("bar"))
	assert.Nil(t, data.At(1000))
	assert.Nil(t, data.At(1000).At(1000, 1000))
	assert.Nil(t, data.At("foo", "bar", "baz", "fred"))
}

func TestSerialize(t *testing.T) {
	var data Value
	var err error
	var str string
	var text []byte

	str = `{"bar":{"int":4,"sub":"val"},"baz":3,"foo":"bar"}`
	err = json.Unmarshal([]byte(str), &data)
	assert.Nil(t, err)

	text, err = json.Marshal(data)
	assert.Nil(t, err)

	assert.Equal(t, str, string(text))
}

func TestObjectsAndLists(t *testing.T) {
	var err error
	var data *Value
	var str string

	str = `{ "foo": "bar", "baz": 4 }`
	err = json.Unmarshal([]byte(str), &data)
	assert.Nil(t, err)

	assert.Equal(t, "bar", data.Map()["foo"].String())
	assert.Equal(t, int64(4), data.Map()["baz"].Number())

	str = `["foo", "bar", 1, 2]`
	err = json.Unmarshal([]byte(str), &data)
	assert.Nil(t, err)

	assert.Equal(t, "bar", data.List()[1].String())
	assert.Equal(t, int64(2), data.List()[3].Number())
	assert.Equal(t, 2, data.List()[3].Int())
}

func TestTypes(t *testing.T) {
	var err error
	var data Value

	err = json.Unmarshal([]byte(`{}`), &data)
	assert.Nil(t, err)

	assert.False(t, data.IsNumber())
	assert.True(t, data.IsObject())
	assert.False(t, data.IsString())
	assert.False(t, data.IsNull())
	assert.False(t, data.IsList())

	err = json.Unmarshal([]byte(`12`), &data)
	assert.Nil(t, err)

	assert.True(t, data.IsNumber())
	assert.False(t, data.IsObject())
	assert.False(t, data.IsString())
	assert.False(t, data.IsNull())
	assert.False(t, data.IsList())
	assert.Equal(t, int64(12), data.Number())
	assert.Equal(t, int64(12), data.NumberOrDefault(102))
	assert.Equal(t, int64(102), data.At("doesnotexist").NumberOrDefault(102))

	err = json.Unmarshal([]byte(`"foo"`), &data)
	assert.Nil(t, err)

	assert.False(t, data.IsNumber())
	assert.False(t, data.IsObject())
	assert.True(t, data.IsString())
	assert.False(t, data.IsNull())
	assert.False(t, data.IsList())
	assert.Equal(t, "foo", data.String())
	assert.Equal(t, "foo", data.StringOrDefault("bar"))
	assert.Equal(t, int64(102), data.NumberOrDefault(102))

	err = json.Unmarshal([]byte(`null`), &data)
	assert.Nil(t, err)

	assert.False(t, data.IsNumber())
	assert.False(t, data.IsObject())
	assert.False(t, data.IsString())
	assert.True(t, data.IsNull())
	assert.False(t, data.IsList())

	err = json.Unmarshal([]byte(`[1,2,3]`), &data)
	assert.Nil(t, err)

	assert.False(t, data.IsNumber())
	assert.False(t, data.IsObject())
	assert.False(t, data.IsString())
	assert.False(t, data.IsNull())
	assert.True(t, data.IsList())

	err = json.Unmarshal([]byte(`true`), &data)
	assert.Nil(t, err)

	assert.False(t, data.IsNumber())
	assert.False(t, data.IsObject())
	assert.False(t, data.IsString())
	assert.False(t, data.IsNull())
	assert.False(t, data.IsList())
	assert.True(t, data.IsBool())

	var dataNil *Value
	assert.False(t, dataNil.IsNumber())
	assert.False(t, dataNil.IsObject())
	assert.False(t, dataNil.IsString())
	assert.True(t, dataNil.IsNull())
	assert.False(t, dataNil.IsList())
}
