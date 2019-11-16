// Package stag exposes functionality to get struct tags values.
//
// Go offers struct tags which are discoverable via reflection but
// you can get the value associated with key in the tag string and
// convert it to approperiate type by stag library easy, clean and quickly.
// it supports most of the data types: int, int8, int32, int64, uint, uint8,
// uint16, uint32, uint64, float32, float64, bool
package stags

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

// STags represents stags
type STags struct {
	t reflect.Type
}

// New returns a new stags for a given struct
func New(s interface{}) *STags {
	return &STags{
		t: reflect.TypeOf(s),
	}
}

// Get returns the string value associated with key in the tag string if the key
// is present in the tag then it return the value otherwise it returns empty string.
func (s STags) Get(tag string, fields ...string) (value string) {
	value, _ = s.get(tag, fields...)
	return
}

// GetInt returns the int value associated with key in the tag string if the key
// is present in the tag then it return the value otherwise it returns zero.
func (s STags) GetInt(tag string, fields ...string) (value int) {
	n, ok := s.getNum(vINT, tag, fields...)
	if ok {
		value = int(n.(int64))
	}
	return
}

// GetInt8 is like GetInt but for int8 numbers.
func (s STags) GetInt8(tag string, fields ...string) (value int8) {
	n, ok := s.getNum(vINT8, tag, fields...)
	if ok {
		value = int8(n.(int64))
	}
	return
}

// GetInt16 is like GetInt but for int16 numbers.
func (s STags) GetInt16(tag string, fields ...string) (value int16) {
	n, ok := s.getNum(vINT16, tag, fields...)
	if ok {
		value = int16(n.(int64))
	}
	return
}

// GetInt32 is like GetInt but for int32 numbers.
func (s STags) GetInt32(tag string, fields ...string) (value int32) {
	n, ok := s.getNum(vINT32, tag, fields...)
	if ok {
		value = int32(n.(int64))
	}
	return
}

// GetInt64 is like GetInt but for int64 numbers.
func (s STags) GetInt64(tag string, fields ...string) (value int64) {
	n, ok := s.getNum(vINT64, tag, fields...)
	if ok {
		value = n.(int64)
	}
	return
}

// GetUint is like GetInt but for uint numbers.
func (s STags) GetUint(tag string, fields ...string) (value uint) {
	n, ok := s.getNum(vUINT, tag, fields...)
	if ok {
		value = uint(n.(uint64))
	}
	return
}

// GetUint8 is like GetInt but for uint8 numbers.
func (s STags) GetUint8(tag string, fields ...string) (value uint8) {
	n, ok := s.getNum(vUINT8, tag, fields...)
	if ok {
		value = uint8(n.(uint64))
	}
	return
}

// GetUint16 is like GetInt but for uint16 numbers.
func (s STags) GetUint16(tag string, fields ...string) (value uint16) {
	n, ok := s.getNum(vUINT16, tag, fields...)
	if ok {
		value = uint16(n.(uint64))
	}
	return
}

// GetUint32 is like GetInt but for uint32 numbers.
func (s STags) GetUint32(tag string, fields ...string) (value uint32) {
	n, ok := s.getNum(vUINT32, tag, fields...)
	if ok {
		value = uint32(n.(uint64))
	}
	return
}

// GetUint64 is like GetInt but for uint64 numbers.
func (s STags) GetUint64(tag string, fields ...string) (value uint64) {
	n, ok := s.getNum(vUINT64, tag, fields...)
	if ok {
		value = n.(uint64)
	}
	return
}

// GetFloat32 is like GetInt but for float32 numbers.
func (s STags) GetFloat32(tag string, fields ...string) (value float32) {
	n, ok := s.getNum(vFLOAT32, tag, fields...)
	if ok {
		value = float32(n.(float64))
	}
	return
}

// GetFloat64 is like GetInt but for float64 numbers.
func (s STags) GetFloat64(tag string, fields ...string) (value float64) {
	n, ok := s.getNum(vFLOAT64, tag, fields...)
	if ok {
		value = n.(float64)
	}
	return
}

// GetBool returns the boolean value associated with key in the tag string if the key
// is present in the tag then it return the value otherwise it returns false.
func (s STags) GetBool(tag string, fields ...string) (value bool) {
	n, ok := s.getNum(vBOOL, tag, fields...)
	if ok {
		value = n.(bool)
	}
	return
}

// Lookup returns the string value associated with key in the tag string if the key
// is present in the tag then it return the string value otherwise the ok returns false.
func (s STags) Lookup(tag string, fields ...string) (value string, ok bool) {
	value, ok = s.get(tag, fields...)
	return
}

// LookupInt returns the int value associated with key in the tag string if the key
// is present in the tag then it return the value otherwise the ok returns false.
func (s STags) LookupInt(tag string, fields ...string) (value int, ok bool) {
	n, ok := s.getNum(vINT, tag, fields...)
	if ok {
		value = int(n.(int64))
	}
	return
}

// LookupInt8 is like LookupInt but for int8 numbers.
func (s STags) LookupInt8(tag string, fields ...string) (value int8, ok bool) {
	n, ok := s.getNum(vINT8, tag, fields...)
	if ok {
		value = int8(n.(int64))
	}
	return
}

// LookupInt16 is like LookupInt but for int16 numbers.
func (s STags) LookupInt16(tag string, fields ...string) (value int16, ok bool) {
	n, ok := s.getNum(vINT16, tag, fields...)
	if ok {
		value = int16(n.(int64))
	}
	return
}

// LookupInt32 is like LookupInt but for int32 numbers.
func (s STags) LookupInt32(tag string, fields ...string) (value int32, ok bool) {
	n, ok := s.getNum(vINT32, tag, fields...)
	if ok {
		value = int32(n.(int64))
	}
	return
}

// LookupInt64 is like LookupInt but for int64 numbers.
func (s STags) LookupInt64(tag string, fields ...string) (value int64, ok bool) {
	n, ok := s.getNum(vINT64, tag, fields...)
	if ok {
		value = n.(int64)
	}
	return
}

// LookupUint is like LookupInt but for uint numbers.
func (s STags) LookupUint(tag string, fields ...string) (value uint, ok bool) {
	n, ok := s.getNum(vUINT, tag, fields...)
	if ok {
		value = uint(n.(uint64))
	}
	return
}

// LookupUint8 is like LookupInt but for uint8 numbers.
func (s STags) LookupUint8(tag string, fields ...string) (value uint8, ok bool) {
	n, ok := s.getNum(vUINT8, tag, fields...)
	if ok {
		value = uint8(n.(uint64))
	}
	return
}

// LookupUint16 is like LookupInt but for uint16 numbers
func (s STags) LookupUint16(tag string, fields ...string) (value uint16, ok bool) {
	n, ok := s.getNum(vUINT16, tag, fields...)
	if ok {
		value = uint16(n.(uint64))
	}
	return
}

// LookupUint32 is like LookupInt but for uint32 numbers
func (s STags) LookupUint32(tag string, fields ...string) (value uint32, ok bool) {
	n, ok := s.getNum(vUINT32, tag, fields...)
	if ok {
		value = uint32(n.(uint64))
	}
	return
}

// LookupUint64 is like LookupInt but for uint64 numbers
func (s STags) LookupUint64(tag string, fields ...string) (value uint64, ok bool) {
	n, ok := s.getNum(vUINT64, tag, fields...)
	if ok {
		value = n.(uint64)
	}
	return
}

// LookupFloat32 is like LookupInt but for float32 numbers
func (s STags) LookupFloat32(tag string, fields ...string) (value float32, ok bool) {
	n, ok := s.getNum(vFLOAT32, tag, fields...)
	if ok {
		value = float32(n.(float64))
	}
	return
}

// LookupFloat64 is like LookupInt but for float64 numbers
func (s STags) LookupFloat64(tag string, fields ...string) (value float64, ok bool) {
	n, ok := s.getNum(vFLOAT64, tag, fields...)
	if ok {
		value = n.(float64)
	}
	return
}

// LookupBool returns the boolean value associated with key in the tag string if the key
// is present in the tag then it return the value otherwise the ok returns false.
func (s STags) LookupBool(tag string, fields ...string) (value bool, ok bool) {
	n, ok := s.getNum(vBOOL, tag, fields...)
	if ok {
		value = n.(bool)
	}
	return
}

func (s STags) getNum(t int, tag string, fields ...string) (interface{}, bool) {
	var (
		n   interface{}
		err error
	)

	v, ok := s.get(tag, fields...)
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

func (s STags) get(tag string, fields ...string) (string, bool) {
	var (
		ok bool
		sf reflect.StructField
	)

	for _, f := range fields {
		sf, ok = s.t.FieldByName(f)
		if !ok {
			return "", ok
		}

		s.t = sf.Type
	}

	return sf.Tag.Lookup(tag)
}
