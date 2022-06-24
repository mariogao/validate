package validate

import (
	"reflect"
	"testing"

	"github.com/gookit/goutil/dump"
	"github.com/stretchr/testify/assert"
)

func TestValueLen(t *testing.T) {
	is := assert.New(t)
	tests := []interface{}{
		"abc",
		123,
		int8(123), int16(123), int32(123), int64(123),
		uint8(123), uint16(123), uint32(123), uint64(123),
		float32(123), float64(123),
		[]int{1, 2, 3}, []string{"a", "b", "c"},
		map[string]string{"k0": "v0", "k1": "v1", "k2": "v2"},
	}

	for _, sample := range tests {
		is.Equal(3, ValueLen(reflect.ValueOf(sample)))
	}

	is.Equal(-1, ValueLen(reflect.ValueOf(nil)))
}

func TestCallByValue(t *testing.T) {
	is := assert.New(t)
	is.Panics(func() {
		CallByValue(reflect.ValueOf("invalid"))
	})
	is.Panics(func() {
		CallByValue(reflect.ValueOf(IsJSON), "age0", "age1")
	})

	rs := CallByValue(reflect.ValueOf(IsNumeric), "123")
	is.Len(rs, 1)
	is.Equal(true, rs[0].Interface())
}

func TestCallByValue_nil_arg(t *testing.T) {
	fn1 := func(in interface{}) interface{} {
		_, ok := in.(NilObject)
		dump.P(in, ok)
		return in
	}

	// runtime error: invalid memory address or nil pointer dereference
	// typ := reflect.TypeOf(interface{}(nil))
	// typ.Kind()

	var nilV interface{}
	nilV = 2

	dump.P(
		reflect.ValueOf(nilV).Kind().String(),
		// reflect.New(reflect.Interface).Kind(),
	)

	rs := CallByValue(reflect.ValueOf(fn1), nil)
	dump.P(rs[0].CanInterface(), rs[0].Interface())
}