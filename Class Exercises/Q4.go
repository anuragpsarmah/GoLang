/*
You are tasked with creating a grade calculator program in GoLang.
The program should prompt the user to input their scores in three subjects: Math, Science, and English.
Based on these scores, calculate the average grade (assuming each subject is equally weighted) and display the corresponding letter grade (A, B, C, D, or F) using control flow.
*/

package main

import (
	"fmt"
)

func main() {
	var mathScore, scienceScore, englishScore float64

	// I am prompting the user to input scores for Math, Science, and English
	fmt.Print("Enter your Math score: ")
	fmt.Scanln(&mathScore)

	fmt.Print("Enter your Science score: ")
	fmt.Scanln(&scienceScore)

	fmt.Print("Enter your English score: ")
	fmt.Scanln(&englishScore)

	// Here i am calculating the average score
	averageScore := (mathScore + scienceScore + englishScore) / 3.0

	// Here I am determining the corresponding letter grade based on the average score
	var letterGrade string
	switch {
	case averageScore >= 90:
		letterGrade = "A"
	case averageScore >= 80:
		letterGrade = "B"
	case averageScore >= 70:
		letterGrade = "C"
	case averageScore >= 60:
		letterGrade = "D"
	default:
		letterGrade = "F"
	}

	// Here I am displaying the average score and corresponding letter grade
	fmt.Printf("Average Score: %.2f\n", averageScore)
	fmt.Printf("Letter Grade: %s\n", letterGrade)
}
