/*
You're tasked with building a student information system in GoLang for a school.
Each student record needs to store details such as student ID, name, age, and grade.
Define variables to store the information of a single student and assign values accordingly.
Pay attention to selecting appropriate data types to represent each piece of information.
*/

package main

import "fmt"

func main() {
	// I have defined a struct for student details.
	type Student struct {
		Name  string
		ID    int
		Age   int
		Grade float32
	}

	// I have create an empty slice to store objects of Student Structure. I used slice because it is of variable size.
	var students []Student

	// I am taking the number of students as input
	var numStudents int
	fmt.Println("Enter the number of students:")
	fmt.Scanln(&numStudents)

	// Using a loop I am taking the student detail inputs
	for i := 0; i < numStudents; i++ {
		var (
			name  string
			id    int
			age   int
			grade float32
		)
		fmt.Printf("ENTER THE STUDENT-%d DETAILS:\n", i+1)
		fmt.Print("Enter the name of the student: ")
		fmt.Scan(&name)
		fmt.Print("Enter the student's ID: ")
		fmt.Scan(&id)
		fmt.Print("Enter the age of the student: ")
		fmt.Scan(&age)
		fmt.Print("Enter the student's grade: ")
		fmt.Scan(&grade)
		fmt.Println()

		temp := Student{Name: name, ID: id, Age: age, Grade: grade}
		students = append(students, temp)
	}

	// Here i am printing the stored inputs
	fmt.Println("\nThe student details are:")
	for i, student := range students {
		fmt.Printf("Student %d:\n", i+1)
		fmt.Println("Name:", student.Name)
		fmt.Println("ID:", student.ID)
		fmt.Println("Age:", student.Age)
		fmt.Println("Grade:", student.Grade)
		fmt.Println()
	}
}
