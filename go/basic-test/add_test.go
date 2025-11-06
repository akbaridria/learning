package main

import "testing"

func testAdd(t *testing.T) {
	result := Add(5, 3)
	expected := 8

	if result != expected {
		t.Errorf("Add(2,3) = %d; want %d", result, expected)
	}
}
