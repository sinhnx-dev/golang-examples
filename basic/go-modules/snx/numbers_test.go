package numbers

import "testing"

func TestIsPrime(t *testing.T) {
	expected := true
	if ret := IsPrime(5); ret != expected {
		t.Errorf("IsPrime(5) = %t, want %t", ret, expected)
	}
}
