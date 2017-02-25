package versionCompare

import (
	"testing"
)

func TestLesser(t *testing.T) {
	v,_ := versionCompare("2.4", "2.5")
	if v != -1 {
		t.Error("Expected -1, got", v)
	}
}

func TestEqual(t *testing.T) {
	v,_ := versionCompare("2.4", "2.4")
	if v != 0 {
		t.Error("Expected 0, got", v)
	}
}

func TestGreater(t *testing.T) {
	v , _:= versionCompare("2.4.1", "2.4")
	if v != 1 {
		t.Error("Expected 1, got", v)
	}
}

func TestPatchGreaterAlpha(t *testing.T) {
	v , _:= versionCompare("2.4.1-version", "2.4.1-b")
	if v != 1 {
		t.Error("Expected 1, got", v)
	}
}

func TestPatchGreaterAlpha2(t *testing.T) {
	v, _ := versionCompare("2.4.2-version", "2.4.1-zeta")
	if v != 1 {
		t.Error("Expected 1, got", v)
	}
}

func TestPatchNondigitVersions(t *testing.T) {
	_ , err:= versionCompare("a.b.c-version", "b.a.c-zeta")
	if err == nil {
		t.Error("Expected error 'Unable to parse input' , got", err)
	}
}
