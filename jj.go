package jj

import (
	"encoding/json"
	"fmt"
	"sort"
)

type Value struct {
	data interface{}
}

type Type int

const (
	Null   = iota
	Number = iota
	String = iota
	Bool   = iota
	Object = iota
	List   = iota
)

func (v *Value) AtOrError(keys ...interface{}) (*Value, error) {
	obj := v
	for _, key := range keys {
		switch key := key.(type) {
		case string:
			h, err := obj.MapOrError()
			if err != nil {
				return nil, err
			}
			obj = h[key]
		case int:
			lst, err := obj.ListOrError()
			if err != nil {
				return nil, err
			}
			obj = lst[key]
		default:
			panic(fmt.Sprintf("jj: key %#v is neither int nor string", key))
		}
	}
	return obj, nil
}

func (v *Value) At(keys ...interface{}) *Value {
	val, _ := v.AtOrError(keys...)
	return val
}

func (v *Value) Type() Type {
	if v == nil {
		return Null
	}
	switch v.data.(type) {
	case string:
		return String
	case float64:
		return Number
	case bool:
		return Bool
	case nil:
		return Null
	case map[string]interface{}:
		return Object
	case []interface{}:
		return List
	default:
		panic(fmt.Errorf("jj: unknown type %#v", v.data))
	}
}

func (v *Value) IsObject() bool {
	return v.Type() == Object
}

func (v *Value) IsNumber() bool {
	return v.Type() == Number
}

func (v *Value) IsString() bool {
	return v.Type() == String
}

func (v *Value) IsNull() bool {
	return v.Type() == Null
}

func (v *Value) IsList() bool {
	return v.Type() == List
}

func (v *Value) IsBool() bool {
	return v.Type() == Bool
}

func (v *Value) StringOrError() (string, error) {
	if v == nil {
		return "", fmt.Errorf("jj: no such key")
	}
	val, ok := v.data.(string)
	if ok {
		return val, nil
	} else {
		return "", fmt.Errorf("jj: object is not a string")
	}

}
func (v *Value) String() string {
	val, err := v.StringOrError()
	if err != nil {
		panic(err)
	}
	return val
}

func (v *Value) StringOrDefault(d string) string {
	val, err := v.StringOrError()
	if err != nil {
		return d
	}
	return val
}

func (v *Value) NumberOrError() (int64, error) {
	if v == nil {
		return 0, fmt.Errorf("jj: no such key")
	}
	d, ok := v.data.(float64)
	if ok {
		return int64(d), nil
	} else {
		return 0, fmt.Errorf("jj: object is not a number")
	}
}

func (v *Value) Number() int64 {
	val, err := v.NumberOrError()
	if err != nil {
		panic(err)
	}
	return val
}

func (v *Value) NumberOrDefault(d int64) int64 {
	val, err := v.NumberOrError()
	if err != nil {
		return d
	}
	return val
}

func (v *Value) MapOrError() (map[string]*Value, error) {
	if v == nil {
		return nil, fmt.Errorf("jj: no such key")
	}
	m, ok := v.data.(map[string]interface{})
	res := make(map[string]*Value, len(m))
	if !ok {
		return res, fmt.Errorf("jj: object is not a map")
	}

	for key, val := range m {
		res[key] = &Value{val}
	}
	return res, nil
}

func (v *Value) Map() map[string]*Value {
	val, err := v.MapOrError()
	if err != nil {
		panic(err)
	}
	return val
}

func (v *Value) ListOrError() ([]*Value, error) {
	if v == nil {
		return nil, fmt.Errorf("jj: no such key")
	}
	m, ok := v.data.([]interface{})
	res := make([]*Value, len(m))
	if !ok {
		return res, fmt.Errorf("jj: object is not a list")
	}
	for i, val := range m {
		res[i] = &Value{val}
	}
	return res, nil
}

func (v *Value) List() []*Value {
	val, err := v.ListOrError()
	if err != nil {
		panic(err)
	}
	return val
}

func (v *Value) KeysSorted() []string {
	m := v.Map()
	res := make([]string, len(m))
	i := 0
	for k, _ := range m {
		res[i] = k
		i++
	}
	sort.Strings(res)
	return res
}

func (v *Value) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &v.data)
}
