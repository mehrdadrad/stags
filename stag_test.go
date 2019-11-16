package stag

import (
	"reflect"
	"testing"
)

type MyStruct struct {
	Field1 bool   `key1:"true" key2:"3.14"`
	Field2 int    `key1:"value1" key2:"-1.6180"`
	Field3 string `key1:"911" key2:"-55667788"`
	Field4 int    `key1:"-55" key2:"66"`
}

var s MyStruct

func TestGetNumWrongField(t *testing.T) {
	_, ok := getNum(s, "Field101", "key1", vBOOL)
	if ok {
		t.Error("expected false but got,", ok)
	}
}

func TestGetNumWrongInt(t *testing.T) {
	_, ok := getNum(s, "Field1", "key1", vINT)
	if ok {
		t.Error("expected false but got,", ok)
	}
}

func TestGetNumWrongUint(t *testing.T) {
	_, ok := getNum(s, "Field1", "key1", vUINT)
	if ok {
		t.Error("expected false but got,", ok)
	}
}

func TestGetNumWrongFloat(t *testing.T) {
	_, ok := getNum(s, "Field1", "key1", vFLOAT32)
	if ok {
		t.Error("expected false but got,", ok)
	}
}

func TestGetNumWrongBool(t *testing.T) {
	_, ok := getNum(s, "Field1", "key2", vBOOL)
	if ok {
		t.Error("expected false but got,", ok)
	}
}

func TestGetNumWrongType(t *testing.T) {
	_, ok := getNum(s, "Field1", "key2", 100)
	if ok {
		t.Error("expected false but got,", ok)
	}
}

func TestGet(t *testing.T) {
	v := Get(s, "Field2", "key1")
	if v != "value1" {
		t.Error("expected key1 but got,", v)
	}
}

func TestUGet(t *testing.T) {
	v, ok := get(s, "Field2", "key1")
	if !ok {
		t.Error("couldn't find key1")
	}

	if v != "value1" {
		t.Error("expected key1 but got,", v)
	}
}

func TestGetInt(t *testing.T) {
	t.Parallel()

	n := GetInt(s, "Field3", "key1")
	if n != 911 {
		t.Error("expected 911 but got,", n)
	}

	k := reflect.TypeOf(n).Kind()
	if k != reflect.Int {
		t.Error("expected type int but got,", k.String())
	}
}

func TestGetInt8(t *testing.T) {
	t.Parallel()

	n := GetInt8(s, "Field4", "key1")
	if n != -55 {
		t.Error("expected -55 but got,", n)
	}

	k := reflect.TypeOf(n).Kind()

	if k != reflect.Int8 {
		t.Error("expected type int8 but got,", k.String())
	}
}

func TestGetInt16(t *testing.T) {
	t.Parallel()

	n := GetInt16(s, "Field4", "key1")
	if n != -55 {
		t.Error("expected -55 but got,", n)
	}

	k := reflect.TypeOf(n).Kind()

	if k != reflect.Int16 {
		t.Error("expected type int16 but got,", k.String())
	}
}

func TestGetInt32(t *testing.T) {
	t.Parallel()

	n := GetInt32(s, "Field3", "key2")
	if n != -55667788 {
		t.Error("expected -55667788 but got,", n)
	}

	k := reflect.TypeOf(n).Kind()

	if k != reflect.Int32 {
		t.Error("expected type int32 but got,", k.String())
	}
}

func TestGetInt64(t *testing.T) {
	t.Parallel()

	n := GetInt64(s, "Field3", "key2")
	if n != -55667788 {
		t.Error("expected -55667788 but got,", n)
	}

	k := reflect.TypeOf(n).Kind()

	if k != reflect.Int64 {
		t.Error("expected type int64 but got,", k.String())
	}
}

func TestGetUint(t *testing.T) {
	t.Parallel()

	n := GetUint(s, "Field3", "key1")
	if n != 911 {
		t.Error("expected 911 but got,", n)
	}

	k := reflect.TypeOf(n).Kind()
	if k != reflect.Uint {
		t.Error("expected type uint but got,", k.String())
	}
}

func TestGetUint8(t *testing.T) {
	t.Parallel()

	n := GetUint8(s, "Field4", "key2")
	if n != 66 {
		t.Error("expected 66 but got,", n)
	}

	k := reflect.TypeOf(n).Kind()

	if k != reflect.Uint8 {
		t.Error("expected type uint8 but got,", k.String())
	}
}

func TestGetUint16(t *testing.T) {
	t.Parallel()

	n := GetUint16(s, "Field4", "key2")
	if n != 66 {
		t.Error("expected 66 but got,", n)
	}

	k := reflect.TypeOf(n).Kind()

	if k != reflect.Uint16 {
		t.Error("expected type uint16 but got,", k.String())
	}
}

func TestGetUint32(t *testing.T) {
	t.Parallel()

	n := GetUint32(s, "Field3", "key1")
	if n != 911 {
		t.Error("expected 911 but got,", n)
	}

	k := reflect.TypeOf(n).Kind()

	if k != reflect.Uint32 {
		t.Error("expected type uint32 but got,", k.String())
	}
}

func TestGetUint64(t *testing.T) {
	t.Parallel()

	n := GetUint64(s, "Field3", "key1")
	if n != 911 {
		t.Error("expected 911 but got,", n)
	}

	k := reflect.TypeOf(n).Kind()

	if k != reflect.Uint64 {
		t.Error("expected type uint64 but got,", k.String())
	}
}

func TestGetFloat32(t *testing.T) {
	t.Parallel()

	n := GetFloat32(s, "Field2", "key2")
	if n != -1.6180 {
		t.Error("expected -1.6180 but got,", n)
	}

	k := reflect.TypeOf(n).Kind()

	if k != reflect.Float32 {
		t.Error("expected type float32 but got,", k.String())
	}
}
func TestGetFloat64(t *testing.T) {
	t.Parallel()

	n := GetFloat64(s, "Field1", "key2")
	if n != 3.14 {
		t.Error("expected 3.14 but got,", n)
	}

	k := reflect.TypeOf(n).Kind()

	if k != reflect.Float64 {
		t.Error("expected type float64 but got,", k.String())
	}
}

func TestGetBool(t *testing.T) {
	t.Parallel()

	n := GetBool(s, "Field1", "key1")
	if n != true {
		t.Error("expected true but got,", n)
	}

	k := reflect.TypeOf(n).Kind()

	if k != reflect.Bool {
		t.Error("expected type bool but got,", k.String())
	}
}

func TestLookup(t *testing.T) {
	v, ok := Lookup(s, "Field2", "key1")
	if !ok {
		t.Error("couldn't find field / key")
	}

	if v != "value1" {
		t.Error("expected key1 but got,", v)
	}
}

func TestLookupInt(t *testing.T) {
	t.Parallel()

	n, ok := LookupInt(s, "Field3", "key1")
	if !ok {
		t.Error("couldn't find field / key")
	}

	if n != 911 {
		t.Error("expected 911 but got,", n)
	}

	k := reflect.TypeOf(n).Kind()
	if k != reflect.Int {
		t.Error("expected type int but got,", k.String())
	}
}

func TestLookupInt8(t *testing.T) {
	t.Parallel()

	n, ok := LookupInt8(s, "Field4", "key1")
	if !ok {
		t.Error("couldn't find field / key")
	}

	if n != -55 {
		t.Error("expected -55 but got,", n)
	}

	k := reflect.TypeOf(n).Kind()

	if k != reflect.Int8 {
		t.Error("expected type int8 but got,", k.String())
	}
}

func TestLookupInt16(t *testing.T) {
	t.Parallel()

	n, ok := LookupInt16(s, "Field4", "key1")
	if !ok {
		t.Error("couldn't find field / key")
	}

	if n != -55 {
		t.Error("expected -55 but got,", n)
	}

	k := reflect.TypeOf(n).Kind()

	if k != reflect.Int16 {
		t.Error("expected type int16 but got,", k.String())
	}
}

func TestLookupInt32(t *testing.T) {
	t.Parallel()

	n, ok := LookupInt32(s, "Field3", "key2")
	if !ok {
		t.Error("couldn't find field / key")
	}

	if n != -55667788 {
		t.Error("expected -55667788 but got,", n)
	}

	k := reflect.TypeOf(n).Kind()

	if k != reflect.Int32 {
		t.Error("expected type int32 but got,", k.String())
	}
}

func TestLookupInt64(t *testing.T) {
	t.Parallel()

	n, ok := LookupInt64(s, "Field3", "key2")
	if !ok {
		t.Error("couldn't find field / key")
	}

	if n != -55667788 {
		t.Error("expected -55667788 but got,", n)
	}

	k := reflect.TypeOf(n).Kind()

	if k != reflect.Int64 {
		t.Error("expected type int64 but got,", k.String())
	}
}

func TestLookupUint(t *testing.T) {
	t.Parallel()

	n, ok := LookupUint(s, "Field3", "key1")
	if !ok {
		t.Error("couldn't find field / key")
	}

	if n != 911 {
		t.Error("expected 911 but got,", n)
	}

	k := reflect.TypeOf(n).Kind()
	if k != reflect.Uint {
		t.Error("expected type uint but got,", k.String())
	}
}

func TestLookupUint8(t *testing.T) {
	t.Parallel()

	n, ok := LookupUint8(s, "Field4", "key2")
	if !ok {
		t.Error("couldn't find field / key")
	}

	if n != 66 {
		t.Error("expected 66 but got,", n)
	}

	k := reflect.TypeOf(n).Kind()

	if k != reflect.Uint8 {
		t.Error("expected type uint8 but got,", k.String())
	}
}

func TestLookupUint16(t *testing.T) {
	t.Parallel()

	n, ok := LookupUint16(s, "Field4", "key2")
	if !ok {
		t.Error("couldn't find field / key")
	}

	if n != 66 {
		t.Error("expected 66 but got,", n)
	}

	k := reflect.TypeOf(n).Kind()

	if k != reflect.Uint16 {
		t.Error("expected type uint16 but got,", k.String())
	}
}

func TestLookupUint32(t *testing.T) {
	t.Parallel()

	n, ok := LookupUint32(s, "Field3", "key1")
	if !ok {
		t.Error("couldn't find field / key")
	}

	if n != 911 {
		t.Error("expected 911 but got,", n)
	}

	k := reflect.TypeOf(n).Kind()

	if k != reflect.Uint32 {
		t.Error("expected type uint32 but got,", k.String())
	}
}

func TestLookupUint64(t *testing.T) {
	t.Parallel()

	n, ok := LookupUint64(s, "Field3", "key1")
	if !ok {
		t.Error("couldn't find field / key")
	}

	if n != 911 {
		t.Error("expected 911 but got,", n)
	}

	k := reflect.TypeOf(n).Kind()

	if k != reflect.Uint64 {
		t.Error("expected type uint64 but got,", k.String())
	}
}

func TestLookupFloat32(t *testing.T) {
	t.Parallel()

	n, ok := LookupFloat32(s, "Field2", "key2")
	if !ok {
		t.Error("couldn't find field / key")
	}

	if n != -1.6180 {
		t.Error("expected -1.6180 but got,", n)
	}

	k := reflect.TypeOf(n).Kind()

	if k != reflect.Float32 {
		t.Error("expected type float32 but got,", k.String())
	}
}
func TestLookupFloat64(t *testing.T) {
	t.Parallel()

	n, ok := LookupFloat64(s, "Field1", "key2")
	if !ok {
		t.Error("couldn't find field / key")
	}

	if n != 3.14 {
		t.Error("expected 3.14 but got,", n)
	}

	k := reflect.TypeOf(n).Kind()

	if k != reflect.Float64 {
		t.Error("expected type float64 but got,", k.String())
	}
}

func TestLookupBool(t *testing.T) {
	t.Parallel()

	n, ok := LookupBool(s, "Field1", "key1")
	if !ok {
		t.Error("couldn't find field / key")
	}

	if n != true {
		t.Error("expected true but got,", n)
	}

	k := reflect.TypeOf(n).Kind()

	if k != reflect.Bool {
		t.Error("expected type bool but got,", k.String())
	}
}
