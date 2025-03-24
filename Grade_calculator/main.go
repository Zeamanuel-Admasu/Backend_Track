package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
)

func CalculateAverage(grade float64, totalSubjects int) (float64, error) {
	if totalSubjects <= 0 {
		return 0, errors.New("invalid number of subjects: must be greater than zero")
	}
	if grade < 0 {
		return 0, errors.New("not a valid grade: can't be negative")
	}
	if grade > (100 * float64(totalSubjects)) {
		return 0, errors.New("not valid grade: exceeds maximum grade")
	}

	average := grade / float64(totalSubjects)
	return math.Round(average*100000) / 100000, nil

}
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter your name")
	scanner.Scan()
	name := scanner.Text()

	var totalSubjects int
	fmt.Println("Enter the number of subjects")
	_, err := fmt.Scanln(&totalSubjects)
	if err != nil || totalSubjects <= 0 {
		fmt.Println("Error: Invalid number of subjects.")
		return
	}

	if totalSubjects < 0 {
		fmt.Println(errors.New("invalid number of subjects"))
		return
	}
	grades := make(map[string]float64)
	var totalGrade float64
	totalGrade = 0
	for i := 0; i < totalSubjects; i++ {
		var subject string
		var grade float64

		fmt.Println("Enter the subject's name: ")
		scanner.Scan()
		subject = scanner.Text()
		fmt.Println("Enter grade for this subject(0 - 100)")
		_, err := fmt.Scanln(&grade)
		if err != nil || grade < 0 || grade > 100 {
			fmt.Println("not a valid grade")
			i--
			continue
		}
		grades[subject] = grade
		totalGrade += (grade)
	}
	averageGrade, err := CalculateAverage(float64(totalGrade), totalSubjects)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Printf("Hey %s, you have taken %d course/s and the average grade point you scored is: %.2f", name, totalSubjects, averageGrade)
	fmt.Println("\nSubjects & Grades:")
	for subject, grade := range grades {
		fmt.Printf("---> %s: %.2f\n", subject, grade)
	}
}
