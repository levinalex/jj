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

	assert.Equal(t, 4, data.At("bar", "int").Number())

	str = `{ "foo": [1,2,"bar",{ "baz": "fred" }]}`
	err = json.Unmarshal([]byte(str), &data)
	assert.Nil(t, err)

	assert.Equal(t, "bar", data.At("foo", 2).String())
	assert.Equal(t, "fred", data.At("foo", 3, "baz").String())
	assert.Equal(t, 1, data.At("foo", 0).Number())
	assert.Equal(t, 2, data.At("foo", 1).Number())
}

func TestObjectsAndLists(t *testing.T) {
	var err error
	var data *Value
	var str string

	str = `{ "foo": "bar", "baz": 4 }`
	err = json.Unmarshal([]byte(str), &data)
	assert.Nil(t, err)

	assert.Equal(t, "bar", data.Map()["foo"].String())
	assert.Equal(t, 4, data.Map()["baz"].Number())

	str = `["foo", "bar", 1, 2]`
	err = json.Unmarshal([]byte(str), &data)
	assert.Nil(t, err)

	assert.Equal(t, "bar", data.List()[1].String())
	assert.Equal(t, 2, data.List()[3].Number())
}

func TestTypes(t *testing.T) {
	var err error
	var data *Value

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
	assert.Equal(t, 12, data.Number())

	err = json.Unmarshal([]byte(`"foo"`), &data)
	assert.Nil(t, err)

	assert.False(t, data.IsNumber())
	assert.False(t, data.IsObject())
	assert.True(t, data.IsString())
	assert.False(t, data.IsNull())
	assert.False(t, data.IsList())
	assert.Equal(t, "foo", data.String())

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

	var dataNil *Value
	assert.False(t, dataNil.IsNumber())
	assert.False(t, dataNil.IsObject())
	assert.False(t, dataNil.IsString())
	assert.True(t, dataNil.IsNull())
	assert.False(t, dataNil.IsList())
}
