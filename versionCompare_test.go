package versionCompare

import (
	"testing"
)

func TestLesser(t *testing.T) {
	v := versionCompare("2.4", "2.5")
	if v != -1 {
		t.Error("Expected -1, got", v)
	}
}

func TestEqual(t *testing.T) {
	v := versionCompare("2.4", "2.4")
	if v != 0 {
		t.Error("Expected 0, got", v)
	}
}

func TestGreater(t *testing.T) {
	v := versionCompare("2.4.1", "2.4")
	if v != 1 {
		t.Error("Expected 1, got", v)
	}
}
