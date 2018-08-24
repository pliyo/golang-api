// main_test.go
package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	var result = 2 + 2
	if result != 4 {
		t.Errorf("Sum is")
	}
}
