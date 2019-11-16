// Package stag exposes functionality to get struct tags values.
//
// Go offers struct tags which are discoverable via reflection but
// you can get the value associated with key in the tag string and
// convert it to approperiate type by stag library easy, clean and quickly.
// it supports most of the data types: int, int8, int32, int64, uint, uint8,
// uint16, uint32, uint64, float32, float64, bool
package stag

import (
	"reflect"
	"strconv"
)

const (
	vINT     = 0
	vINT8    = 8
	vINT16   = 16
	vINT32   = 32
	vINT64   = 64
	vUINT    = 1
	vUINT8   = 9
	vUINT16  = 17
	vUINT32  = 33
	vUINT64  = 65
	vFLOAT32 = 34
	vFLOAT64 = 66
	vBOOL    = 67
)

// Get returns the string value associated with key in the tag string if the key
// is present in the tag then it return the value otherwise it returns empty string.
func Get(s interface{}, field, tag string) (value string) {
	value, _ = get(s, field, tag)
	return
}

// GetInt returns the int value associated with key in the tag string if the key
// is present in the tag then it return the value otherwise it returns zero.
func GetInt(s interface{}, field, tag string) (value int) {
	n, _ := getNum(s, field, tag, vINT)
	value = int(n.(int64))
	return
}

// GetInt8 is like GetInt but for int8 numbers.
func GetInt8(s interface{}, field, tag string) (value int8) {
	n, _ := getNum(s, field, tag, vINT8)
	value = int8(n.(int64))
	return
}

// GetInt16 is like GetInt but for int16 numbers.
func GetInt16(s interface{}, field, tag string) (value int16) {
	n, _ := getNum(s, field, tag, vINT16)
	value = int16(n.(int64))
	return
}

// GetInt32 is like GetInt but for int32 numbers.
func GetInt32(s interface{}, field, tag string) (value int32) {
	n, _ := getNum(s, field, tag, vINT32)
	value = int32(n.(int64))
	return
}

// GetInt64 is like GetInt but for int64 numbers.
func GetInt64(s interface{}, field, tag string) (value int64) {
	n, _ := getNum(s, field, tag, vINT64)
	value = n.(int64)
	return
}

// GetUint is like GetInt but for uint numbers.
func GetUint(s interface{}, field, tag string) (value uint) {
	n, _ := getNum(s, field, tag, vUINT)
	value = uint(n.(uint64))
	return
}

// GetUint8 is like GetInt but for uint8 numbers.
func GetUint8(s interface{}, field, tag string) (value uint8) {
	n, _ := getNum(s, field, tag, vUINT8)
	value = uint8(n.(uint64))
	return
}

// GetUint16 is like GetInt but for uint16 numbers.
func GetUint16(s interface{}, field, tag string) (value uint16) {
	n, _ := getNum(s, field, tag, vUINT16)
	value = uint16(n.(uint64))
	return
}

// GetUint32 is like GetInt but for uint32 numbers.
func GetUint32(s interface{}, field, tag string) (value uint32) {
	n, _ := getNum(s, field, tag, vUINT32)
	value = uint32(n.(uint64))
	return
}

// GetUint64 is like GetInt but for uint64 numbers.
func GetUint64(s interface{}, field, tag string) (value uint64) {
	n, _ := getNum(s, field, tag, vUINT64)
	value = n.(uint64)
	return
}

// GetFloat32 is like GetInt but for float32 numbers.
func GetFloat32(s interface{}, field, tag string) (value float32) {
	n, _ := getNum(s, field, tag, vFLOAT32)
	value = float32(n.(float64))
	return
}

// GetFloat64 is like GetInt but for float64 numbers.
func GetFloat64(s interface{}, field, tag string) (value float64) {
	n, _ := getNum(s, field, tag, vFLOAT64)
	value = n.(float64)
	return
}

// GetBool returns the boolean value associated with key in the tag string if the key
// is present in the tag then it return the value otherwise it returns false.
func GetBool(s interface{}, field, tag string) bool {
	n, _ := getNum(s, field, tag, vBOOL)
	return n.(bool)
}

// Lookup returns the string value associated with key in the tag string if the key
// is present in the tag then it return the string value otherwise the ok returns false.
func Lookup(s interface{}, field, tag string) (value string, ok bool) {
	value, ok = get(s, field, tag)
	return
}

// LookupInt returns the int value associated with key in the tag string if the key
// is present in the tag then it return the value otherwise the ok returns false.
func LookupInt(s interface{}, field, tag string) (value int, ok bool) {
	n, ok := getNum(s, field, tag, vINT)
	value = int(n.(int64))
	return
}

// LookupInt8 is like LookupInt but for int8 numbers.
func LookupInt8(s interface{}, field, tag string) (value int8, ok bool) {
	n, ok := getNum(s, field, tag, vINT8)
	value = int8(n.(int64))
	return
}

// LookupInt16 is like LookupInt but for int16 numbers.
func LookupInt16(s interface{}, field, tag string) (value int16, ok bool) {
	n, ok := getNum(s, field, tag, vINT16)
	value = int16(n.(int64))
	return
}

// LookupInt32 is like LookupInt but for int32 numbers.
func LookupInt32(s interface{}, field, tag string) (value int32, ok bool) {
	n, ok := getNum(s, field, tag, vINT32)
	value = int32(n.(int64))
	return
}

// LookupInt64 is like LookupInt but for int64 numbers.
func LookupInt64(s interface{}, field, tag string) (value int64, ok bool) {
	n, ok := getNum(s, field, tag, vINT64)
	value = n.(int64)
	return
}

// LookupUint is like LookupInt but for uint numbers.
func LookupUint(s interface{}, field, tag string) (value uint, ok bool) {
	n, ok := getNum(s, field, tag, vUINT)
	value = uint(n.(uint64))
	return
}

// LookupUint8 is like LookupInt but for uint8 numbers.
func LookupUint8(s interface{}, field, tag string) (value uint8, ok bool) {
	n, ok := getNum(s, field, tag, vUINT8)
	value = uint8(n.(uint64))
	return
}

// LookupUint16 is like LookupInt but for uint16 numbers
func LookupUint16(s interface{}, field, tag string) (value uint16, ok bool) {
	n, ok := getNum(s, field, tag, vUINT16)
	value = uint16(n.(uint64))
	return
}

// LookupUint32 is like LookupInt but for uint32 numbers
func LookupUint32(s interface{}, field, tag string) (value uint32, ok bool) {
	n, ok := getNum(s, field, tag, vUINT32)
	value = uint32(n.(uint64))
	return
}

// LookupUint64 is like LookupInt but for uint64 numbers
func LookupUint64(s interface{}, field, tag string) (value uint64, ok bool) {
	n, ok := getNum(s, field, tag, vUINT64)
	value = n.(uint64)
	return
}

// LookupFloat32 is like LookupInt but for float32 numbers
func LookupFloat32(s interface{}, field, tag string) (value float32, ok bool) {
	n, ok := getNum(s, field, tag, vFLOAT32)
	value = float32(n.(float64))
	return
}

// LookupFloat64 is like LookupInt but for float64 numbers
func LookupFloat64(s interface{}, field, tag string) (value float64, ok bool) {
	n, ok := getNum(s, field, tag, vFLOAT64)
	value = n.(float64)
	return
}

// LookupBool returns the boolean value associated with key in the tag string if the key
// is present in the tag then it return the value otherwise the ok returns false.
func LookupBool(s interface{}, field, tag string) (bool, bool) {
	n, ok := getNum(s, field, tag, vBOOL)
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
	case vINT, vINT8, vINT16, vINT32, vINT64:
		n, err = strconv.ParseInt(v, 10, t)
		if err != nil {
			return 0, !ok
		}
	case vUINT, vUINT8, vUINT16, vUINT32, vUINT64:
		n, err = strconv.ParseUint(v, 10, t-1)
		if err != nil {
			return 0, !ok
		}
	case vFLOAT32, vFLOAT64:
		n, err = strconv.ParseFloat(v, t-2)
		if err != nil {
			return 0, !ok
		}
	case vBOOL:
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
