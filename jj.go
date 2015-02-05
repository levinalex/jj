package jj

import (
	"encoding/json"
	"fmt"
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

func getValueAtPath(o interface{}, keys ...interface{}) (interface{}, error) {
	obj := o

	for _, key := range keys {

		switch key := key.(type) {
		case string:
			h, ok := obj.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("jsvalue: cannot treat %v as a map", obj)
			}
			obj = h[key]
		case int:
			lst, ok := obj.([]interface{})
			if !ok {
				return nil, fmt.Errorf("jsvalue: cannot treat %v as a list", obj)
			}
			obj = lst[key]
		default:
			panic(fmt.Sprintf("jsvalue: key %#v is neither int nor string", key))
		}
	}
	return obj, nil
}

func (v *Value) At(keys ...interface{}) *Value {
	val, _ := getValueAtPath(v.data, keys...)
	if val == nil {
		return nil
	} else {
		return &Value{val}
	}
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
	case map[string]interface{}:
		return Object
	case []interface{}:
		return List
	default:
		panic("unknown type")
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

func (v *Value) Map() map[string]*Value {
	m, ok := v.data.(map[string]interface{})
	if !ok {
		panic("jj: object is not a map")
	}

	res := make(map[string]*Value, len(m))
	for key, val := range m {
		res[key] = &Value{val}
	}
	return res
}

func (v *Value) List() []*Value {
	m, ok := v.data.([]interface{})
	if !ok {
		panic("jj: object is not a list")
	}
	res := make([]*Value, len(m))
	for i, val := range m {
		res[i] = &Value{val}
	}
	return res
}

func (v *Value) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &v.data)
}
