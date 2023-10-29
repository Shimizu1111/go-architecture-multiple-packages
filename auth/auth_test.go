package auth

import (
	"testing"
)

func TestAuthenticate(t *testing.T) {
	if !Authenticate("user", "pass") {
		t.Errorf("Authentication failed")
	}
}
