package mypkg

import "fmt"

// Greet prints a friendly greeting.
func Greet(salutation string) {
	fmt.Println(salutation, "to", Name)
}

// Name of whom you want to greet.
var Name = "Milton"
