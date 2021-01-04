package internal

import (
	"reflect"
	"testing"
)

func assertEquals(t *testing.T, v1, v2 interface{}) {
	if v1 != v2 {
		t.Error()
	}
}

func assertNotEquals(t *testing.T, v1, v2 interface{}) {
	if v1 == v2 {
		t.Error()
	}
}

func assertNil(t *testing.T, v interface{}, expNil bool) {
	val := reflect.ValueOf(v)
	if (val.IsNil() && !expNil) || (!val.IsNil() && expNil) {
		t.Error()
	}
}

func assertError(t *testing.T, err error, expError bool) {
	if (err != nil && !expError) || (err == nil && expError) {
		t.Error()
	}
}
