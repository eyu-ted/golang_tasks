package main

import (
	"fmt"
)

func CalculateAverageGrade(grades map[string]float64) float64 {
	sum := 0.0
	for _, grade := range grades {
		sum += grade
	}
	return sum / float64(len(grades))
}

func AcceptInput() (string, map[string]float64) {
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

	record := make(map[string]float64)

	for i := 0; i < numSubjects; i++ {
		var subject string
		var grade float64

		fmt.Printf("Enter the name of subject %d: ", i+1)
		fmt.Scanln(&subject)

		fmt.Printf("Enter the grade for %s: ", subject)
		_, err := fmt.Scanln(&grade)
		for err != nil || grade < 0 || grade > 100 {
			fmt.Println("Invalid grade. Please enter a number between 0 and 100.")
			_, err = fmt.Scanln(&grade)
		}
		record[subject] = grade
	}

	return studentName, record
}

func DisplayResult(studentName string, record map[string]float64, averageGrade float64) {
	fmt.Printf("\nStudent Name: %s\n", studentName)
	fmt.Println("Subject Grades:")
	for subject, grade := range record {
		fmt.Printf("%s: %.2f\n", subject, grade)
	}
	fmt.Printf("Average Grade: %.2f\n", averageGrade)
}

func main() {
	studentName, record := AcceptInput()
	averageGrade := CalculateAverageGrade(record)
	DisplayResult(studentName, record, averageGrade)
}
