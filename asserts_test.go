package clo

import (
	"reflect"
	"testing"
)

func assertValEquals(t *testing.T, v1, v2 interface{}) {
	if v1 != v2 {
		t.Errorf("expected equal values: %v and %v", v1, v2)
	}
}

func assertValNotEquals(t *testing.T, v1, v2 interface{}) {
	if v1 == v2 {
		t.Errorf("expected different values: %v and %v", v1, v2)
	}
}

func assertNil(t *testing.T, v interface{}, expNil bool) {
	val := reflect.ValueOf(v)
	if (val.IsNil() && !expNil) || (!val.IsNil() && expNil) {
		t.Errorf("v: %v; expNil: %v", v, expNil)
	}
}

func assertError(t *testing.T, err error, expError bool) {
	if err != nil && !expError {
		t.Error("unexpected error: ", err)
	}
	if err == nil && expError {
		t.Error("missing expected error")
	}
}

func assertInterfaceEquals(t *testing.T, r1, r2 interface{}) {
	if !reflect.DeepEqual(r1, r2) {
		t.Errorf("expected %v and %v to be equal", r1, r2)
	}
}
