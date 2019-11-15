package stag

import (
	"reflect"
	"strconv"
)

const (
	BOOL    = 0
	INT     = 1
	INT8    = 8
	INT16   = 16
	INT32   = 32
	INT64   = 64
	UINT    = 2
	UINT8   = 9
	UINT16  = 17
	UINT32  = 33
	UINT64  = 65
	FLOAT32 = 34
	FLOAT64 = 66
)

func Get(s interface{}, field, tag string) (value string, ok bool) {
	value, ok = get(s, field, tag)
	return
}

func GetInt(s interface{}, field, tag string) (value int, ok bool) {
	n, ok := getNum(s, field, tag, INT)
	value = n.(int)
	return
}

func GetInt8(s interface{}, field, tag string) (value int8, ok bool) {
	n, ok := getNum(s, field, tag, INT8)
	value = n.(int8)
	return
}

func GetInt16(s interface{}, field, tag string) (value int16, ok bool) {
	n, ok := getNum(s, field, tag, INT16)
	value = n.(int16)
	return
}

func GetInt32(s interface{}, field, tag string) (value int32, ok bool) {
	n, ok := getNum(s, field, tag, INT32)
	value = n.(int32)
	return
}

func GetInt64(s interface{}, field, tag string) (value int64, ok bool) {
	n, ok := getNum(s, field, tag, INT64)
	value = n.(int64)
	return
}

func GetUint8(s interface{}, field, tag string) (value uint8, ok bool) {
	n, ok := getNum(s, field, tag, UINT8)
	value = n.(uint8)
	return
}

func GetUint16(s interface{}, field, tag string) (value uint16, ok bool) {
	n, ok := getNum(s, field, tag, UINT16)
	value = n.(uint16)
	return
}

func GetUint32(s interface{}, field, tag string) (value uint32, ok bool) {
	n, ok := getNum(s, field, tag, UINT32)
	value = n.(uint32)
	return
}

func GetUint64(s interface{}, field, tag string) (value uint64, ok bool) {
	n, ok := getNum(s, field, tag, UINT64)
	value = n.(uint64)
	return
}

func GetFloat32(s interface{}, field, tag string) (value float32, ok bool) {
	n, ok := getNum(s, field, tag, FLOAT32)
	value = n.(float32)
	return
}

func GetFloat64(s interface{}, field, tag string) (value float64, ok bool) {
	n, ok := getNum(s, field, tag, FLOAT64)
	value = n.(float64)
	return
}

// GetBool returns the boolean value associated with key in the tag string if the key
// is present in the tag then it return the value otherwise the ok returns false.
func GetBool(s interface{}, field, tag string) (bool, bool) {
	n, ok := getNum(s, field, tag, BOOL)
	return n.(bool), ok
}

func getNum(s interface{}, field, tag string, t int) (interface{}, bool) {
	var (
		n   interface{}
		err error
	)

	v, ok := get(s, field, tag)
	if !ok {
		return 0, ok
	}

	switch t {
	case INT:
		n, err = strconv.Atoi(v)
		if err != nil {
			return 0, !ok
		}

	case UINT:
		// TODO
	case INT8, INT16, INT32, INT64:
		n, err = strconv.ParseInt(v, 10, t)
		if err != nil {
			return 0, !ok
		}
	case UINT8, UINT16, UINT32, UINT64:
		n, err = strconv.ParseUint(v, 10, t-1)
		if err != nil {
			return 0, !ok
		}
	case FLOAT32, FLOAT64:
		n, err = strconv.ParseFloat(v, t-2)
		if err != nil {
			return 0, !ok
		}
	case BOOL:
		n, err = strconv.ParseBool(v)
		if err != nil {
			return 0, !ok
		}
	default:
		return 0, false
	}

	return n, ok
}

func get(s interface{}, field, tag string) (string, bool) {
	t := reflect.TypeOf(s)
	sf, ok := t.FieldByName(field)
	if !ok {
		return "", ok
	}

	return sf.Tag.Lookup(tag)
}
