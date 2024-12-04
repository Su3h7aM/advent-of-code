package main

import "testing"

func TestHello(t *testing.T) {
	result := Hello()

	if result != "hello" {
		t.Errorf("Hello() should return 'hello', got %s", result)
	}
}
