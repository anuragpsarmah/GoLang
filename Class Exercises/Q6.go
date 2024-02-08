/*
Develop a scenario based on your domain that incorporates looping, control structures, variables, and constants.
Calculate the factorial of a number.
*/

package main

import (
	"fmt"
)

// Here I am defining constants for BMI categories
const (
	Underweight  = "underweight"
	NormalWeight = "normal weight"
	Overweight   = "overweight"
)

// Recursive function to calculate factorial
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	// Here I am creating a structure to store health data of users
	type userHealthData struct {
		Name   string
		Height float32
		Weight float32
	}

	// Here I am creating a slice for the structure objects
	var users []userHealthData
	var numUsers int
	fmt.Print("Enter the number of users: ")
	fmt.Scan(&numUsers)

	for i := 0; i < numUsers; i++ {
		fmt.Printf("\nENTER THE DETAILS OF USER-%d:\n", i+1)
		var (
			name   string
			height float32
			weight float32
		)
		fmt.Print("Enter your name: ")
		fmt.Scan(&name)
		fmt.Print("Enter your height: ")
		fmt.Scan(&height)
		fmt.Print("Enter your weight: ")
		fmt.Scan(&weight)

		temp := userHealthData{Name: name, Height: height, Weight: weight}
		users = append(users, temp)
	}

	for i, temp := range users {
		fmt.Printf("\nRecommendation for User-%d:\n", i+1)
		var BMI float32 = temp.Weight / (temp.Height * temp.Height)

		// Here I am determining BMI category based on the calculated BMI
		var BMIcategory string
		switch {
		case BMI < 18.5:
			BMIcategory = Underweight
		case BMI < 24.9:
			BMIcategory = NormalWeight
		default:
			BMIcategory = Overweight
		}

		// Here I am displaying recommendation based on BMI category
		switch BMIcategory {
		case Underweight:
			fmt.Printf("Sorry %s, your BMI is less than 18.5. You are underweight!", temp.Name)
		case NormalWeight:
			fmt.Printf("Congratulations %s, your BMI is between 18.5 and 24.9. You are of normal weight!", temp.Name)
		case Overweight:
			fmt.Printf("Sorry %s, your BMI is more than 24.9. You are overweight!", temp.Name)
		}
		fmt.Println()

		// Here I am calculating factorial of the number of users
		result := factorial(numUsers)

		// Here i am displaying the factorial
		fmt.Printf("\nFactorial of %d (number of users) is: %d\n", numUsers, result)
	}
}
