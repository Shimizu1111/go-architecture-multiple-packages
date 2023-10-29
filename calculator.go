package main

import (
	"fmt"

	"github.com/shimizu/calculator/auth"
	"github.com/shimizu/calculator/internal/add"
	"github.com/shimizu/calculator/internal/subtract"
)

func main() {
	a, b := 10, 5
	sum, diff := performCalculations(a, b)
	fmt.Printf("%d + %d = %d\n", a, b, sum)
	fmt.Printf("%d - %d = %d\n", a, b, diff)

	username, password := "user", "pass"
	isAuthenticated := performAuthentication(username, password)
	fmt.Printf("Authentication successful: %v\n", isAuthenticated)
}

func performCalculations(a, b int) (int, int) {
	return add.Add(a, b), subtract.Subtract(a, b)
}

func performAuthentication(username, password string) bool {
	return auth.Authenticate(username, password)
}
