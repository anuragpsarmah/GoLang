/*
You want to build a simple interest calculator in GoLang.
The program should ask the user to input multiple sets of data,each containing the principal amount, the annual interest rate, and the number of years for which the interest is to be calculated.
Use arrays to store the input data for each set, calculate the simple interest for each set using the formula: Simple Interest = (Principal Amount * Annual Interest Rate * Number of Years) / 100, and display the results.
*/

package main

import (
	"fmt"
)

func main() {
	var numSets int

	// Here I am prompting the user to input the number of sets of data
	fmt.Print("Enter the number of sets of data: ")
	fmt.Scanln(&numSets)

	// Here I am declaring arrays to store principal amount, annual interest rate, and number of years for each set of data
	principal := make([]float64, numSets)
	interestRate := make([]float64, numSets)
	years := make([]float64, numSets)

	// Here I am taking input data for each set
	for i := 0; i < numSets; i++ {
		fmt.Printf("\nEnter data for set %d:\n", i+1)
		fmt.Print("Principal amount: ")
		fmt.Scanln(&principal[i])
		fmt.Print("Annual interest rate: ")
		fmt.Scanln(&interestRate[i])
		fmt.Print("Number of years: ")
		fmt.Scanln(&years[i])
	}

	// Here I am calculating and displaying simple interest for each set
	fmt.Println("\nSimple Interest for Each Set:")
	for i := 0; i < numSets; i++ {
		simpleInterest := (principal[i] * interestRate[i] * years[i]) / 100.0
		fmt.Printf("Set %d: %.2f\n", i+1, simpleInterest)
	}
}
