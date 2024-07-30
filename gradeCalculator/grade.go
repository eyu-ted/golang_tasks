

package main

import (
	"fmt"
)

// CalculateAverageGrade calculates the average of the given grades.
func CalculateAverageGrade(grades []float64) float64 {
	sum := 0.0
	for _, grade := range grades {
		sum += grade
	}
	return sum / float64(len(grades))
}

// AcceptInput accepts input from the user and returns the student's name, subjects, and grades.
func AcceptInput() (string, []string, []float64) {
	var studentName string
	var numSubjects int

	fmt.Print("Enter your name: ")
	fmt.Scanln(&studentName)

	fmt.Print("Enter the number of subjects: ")
	_, err := fmt.Scanln(&numSubjects)
	for err != nil || numSubjects <= 0 {
		fmt.Println("Invalid number of subjects. Please enter a positive integer.")
		_, err = fmt.Scanln(&numSubjects)
	}

	subjects := make([]string, numSubjects)
	grades := make([]float64, numSubjects)

	for i := 0; i < numSubjects; i++ {
		fmt.Printf("Enter the name of subject %d: ", i+1)
		fmt.Scan(&subjects[i])

		fmt.Printf("Enter the grade for %s: ", subjects[i])
		_, err := fmt.Scan(&grades[i])
		for err != nil || grades[i] < 0 || grades[i] > 100 {
			fmt.Println("Invalid grade. Please enter a number between 0 and 100.")
			_, err = fmt.Scan(&grades[i])
		}
	}

	return studentName, subjects, grades
}

// DisplayResult displays the student's name, subjects, grades, and average grade.
func DisplayResult(studentName string, subjects []string, grades []float64, averageGrade float64) {
	fmt.Printf("\nStudent Name: %s\n", studentName)
	fmt.Println("Subject Grades:")
	for i, subject := range subjects {
		fmt.Printf("%s: %.2f\n", subject, grades[i])
	}
	fmt.Printf("Average Grade: %.2f\n", averageGrade)
}

// main is the entry point of the program.
func main() {
	studentName, subjects, grades := AcceptInput()
	averageGrade := CalculateAverageGrade(grades)
	DisplayResult(studentName, subjects, grades, averageGrade)
}

