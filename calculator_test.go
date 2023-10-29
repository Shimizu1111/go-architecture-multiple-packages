package main

import (
	"testing"
)

func TestPerformCalculations(t *testing.T) {
	sum, diff := performCalculations(10, 5)
	if sum != 15 || diff != 5 {
		t.Errorf("performCalculations failed")
	}
}

func TestPerformAuthentication(t *testing.T) {
	isAuthenticated := performAuthentication("user", "pass")
	if !isAuthenticated {
		t.Errorf("performAuthentication failed")
	}
}
