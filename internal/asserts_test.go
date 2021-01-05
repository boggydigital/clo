package internal

import (
	"reflect"
	"testing"
)

func assertValEquals(t *testing.T, v1, v2 interface{}) {
	if v1 != v2 {
		t.Error()
	}
}

func assertValNotEquals(t *testing.T, v1, v2 interface{}) {
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

func assertInterfaceEquals(t *testing.T, r1, r2 interface{}) {
	if !reflect.DeepEqual(r1, r2) {
		t.Error()
	}
}
