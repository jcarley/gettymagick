package testing

import "testing"

func AssertTrue(t *testing.T, expression bool, msg string) {
	if !expression {
		t.Error(msg)
	}
}

func AssertFalse(t *testing.T, expression bool, msg string) {
	if expression {
		t.Error(msg)
	}
}
