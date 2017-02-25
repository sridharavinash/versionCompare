package versionCompare

import (
	"testing"
)

func TestEqualFail(t *testing.T) {
	v := versionCompare("2.4", "2.5", "=")
	if v  {
		t.Error("Expected false, got", v)
	}
}

func TestEqualPass(t *testing.T) {
	v := versionCompare("2.4", "2.4", "=")
	if !v {
		t.Error("Expected true, got", v)
	}
}
