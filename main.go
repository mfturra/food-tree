package main

import (
	"fmt"
)

func main() {
	// Acquire protein input
	fmt.Println("Welcome to the food recipe generator! What do you have?")
	fmt.Println("Enter the protein:")
	var protein string
	fmt.Scanln(&protein)

	// Acquire veggie input
	fmt.Println("Enter the veggie:")
	var veggie string
	fmt.Scanln(&veggie)

	// Output ingredients
	fmt.Print("Looks like you have: " + protein + " & " + veggie)

}
