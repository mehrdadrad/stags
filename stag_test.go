package stags

import (
	"reflect"
	"testing"
)

type MyStruct struct {
	Field1 bool   `key1:"true" key2:"3.14"`
	Field2 int    `key1:"value1" key2:"-1.6180"`
	Field3 string `key1:"911" key2:"-55667788"`
	Field4 int    `key1:"-55" key2:"66"`
	Field5 struct {
		Field1 int `key1:"value1" key2:"178"`
	}
}

var s MyStruct
var st = New(s)

func TestGetNumWrongField(t *testing.T) {
	_, ok := st.getNum(vBOOL, "key1", "Field101")
	if ok {
		t.Error("expected false but got,", ok)
	}
}

func TestGetNumWrongInt(t *testing.T) {
	_, ok := st.getNum(vINT, "key1", "Field1")
	if ok {
		t.Error("expected false but got,", ok)
	}
}

func TestGetNumWrongUint(t *testing.T) {
	_, ok := st.getNum(vUINT, "key1", "Field1")
	if ok {
		t.Error("expected false but got,", ok)
	}
}

func TestGetNumWrongFloat(t *testing.T) {
	_, ok := st.getNum(vFLOAT32, "key1", "Field1")
	if ok {
		t.Error("expected false but got,", ok)
	}
}

func TestGetNumWrongBool(t *testing.T) {
	_, ok := st.getNum(vBOOL, "key2", "Field1")
	if ok {
		t.Error("expected false but got,", ok)
	}
}

func TestGetNumWrongType(t *testing.T) {
	_, ok := st.getNum(100, "key2", "Field1")
	if ok {
		t.Error("expected false but got,", ok)
	}
}

func TestGet(t *testing.T) {
	v := st.Get("key1", "Field2")
	if v != "value1" {
		t.Error("expected key1 but got,", v)
	}
}

func TestGetNested(t *testing.T) {
	v := st.Get("key1", "Field5", "Field1")
	if v != "value1" {
		t.Error("expected key1 but got,", v)
	}
}

func TestUGet(t *testing.T) {
	v, ok := st.get("key1", "Field2")
	if !ok {
		t.Error("couldn't find key1")
	}

	if v != "value1" {
		t.Error("expected key1 but got,", v)
	}
}

func TestGetInt(t *testing.T) {
	t.Parallel()

	n := st.GetInt("key1", "Field3")
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

	n := st.GetInt8("key1", "Field4")
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

	n := st.GetInt16("key1", "Field4")
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

	n := st.GetInt32("key2", "Field3")
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

	n := st.GetInt64("key2", "Field3")
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

	n := st.GetUint("key1", "Field3")
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

	n := st.GetUint8("key2", "Field4")
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

	n := st.GetUint16("key2", "Field4")
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

	n := st.GetUint32("key1", "Field3")
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

	n := st.GetUint64("key1", "Field3")
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

	n := st.GetFloat32("key2", "Field2")
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

	n := st.GetFloat64("key2", "Field1")
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

	n := st.GetBool("key1", "Field1")
	if n != true {
		t.Error("expected true but got,", n)
	}

	k := reflect.TypeOf(n).Kind()

	if k != reflect.Bool {
		t.Error("expected type bool but got,", k.String())
	}
}

func TestLookup(t *testing.T) {
	v, ok := st.Lookup("key1", "Field2")
	if !ok {
		t.Error("couldn't find field / key")
	}

	if v != "value1" {
		t.Error("expected key1 but got,", v)
	}
}

func TestLookupInt(t *testing.T) {
	t.Parallel()

	n, ok := st.LookupInt("key1", "Field3")
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

	n, ok := st.LookupInt8("key1", "Field4")
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

	n, ok := st.LookupInt16("key1", "Field4")
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

	n, ok := st.LookupInt32("key2", "Field3")
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

	n, ok := st.LookupInt64("key2", "Field3")
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

	n, ok := st.LookupUint("key1", "Field3")
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

	n, ok := st.LookupUint8("key2", "Field4")
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

	n, ok := st.LookupUint16("key2", "Field4")
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

	n, ok := st.LookupUint32("key1", "Field3")
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

	n, ok := st.LookupUint64("key1", "Field3")
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

	n, ok := st.LookupFloat32("key2", "Field2")
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

	n, ok := st.LookupFloat64("key2", "Field1")
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

	n, ok := st.LookupBool("key1", "Field1")
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
