package main

import (
	"testing"
)

func TestCommonPrefix(t *testing.T) {
	testPaths := []string{"/var/cache/salt", "/var/lib/whatever"}
	common := CommonPrefix(testPaths)
	if common != "/var/" {
		t.Fatalf("Failed to find common path: %s", common)
	}
	testPaths = []string{"/root", "/var/cache/salt", "/var/lib/whatever"}
	common = CommonPrefix(testPaths)
	if common != "/" {
		t.Fatalf("Failed to find common path: %s", common)
	}
	testPaths = []string{"uncommon", "/var/cache/salt", "/var/lib/whatever"}
	common = CommonPrefix(testPaths)
	if common != "" {
		t.Fatalf("Failed to find common path: %s", common)
	}
}
