package example

import (
	"testing"
)

func testUnexportedFunction(t *testing.T) {
	t.Fatal("this function doesn't run")
}

func TestUnexportedFunction(t *testing.T) {
	_, err := unexportedFunction(true)
	if err != nil {
		t.Errorf("wanted %q, got %v", Message, err)
	}
}
