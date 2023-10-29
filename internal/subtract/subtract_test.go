package subtract

import (
	"testing"
)

func TestSubtract(t *testing.T) {
	if Subtract(5, 3) != 2 {
		t.Errorf("Subtraction failed")
	}
}
