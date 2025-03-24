package main

import (
	"fmt"
	"math"
	"testing"
)

func TestCalculateAverage(t *testing.T) {

	var tests = []struct {
		totalGrade    float64
		totalSubjects int
		expected      float64
		expectError   bool
	}{
		{300, 3, 100.0, false},
		{0, 0, 0.0, true},
		{250, 5, 50.0, false},
		{0, 3, 0.0, false},
		{100, 1, 100.0, false},
		{-10, 2, 0, true},
		{200, 0, 0, true},
		{500, 4, 0, true},
		{1000, 9, 0, true},
		{299, 3, 99.66667, false},
		{0.0001, 1, 0.0001, false},
		{100, 1, 100.0, false},
		{50, 0, 0.0, true},
		{250.5, 5, 50.1, false},
		{1e6, 10000, 100.0, false},
		{1e12, 1e6, 0.0, true},
		{299.9999999, 3, 100.00, false},
		{299, 3, 99.66667, false},
	}

	for _, test := range tests {
		result, err := CalculateAverage(test.totalGrade, test.totalSubjects)
		if test.expectError {
			if err == nil {
				t.Errorf("Expected an error for totalGrade=%.2f, totalSubjects=%d, but got none", test.totalGrade, test.totalSubjects)
			}
		} else {
			if err != nil {
				t.Errorf("Did not expect an error for totalGrade=%.2f, totalSubjects=%d, but got: %v", test.totalGrade, test.totalSubjects, err)
			}
			const epsilon = 0.00001
			if abs(result-test.expected) > epsilon {
				fmt.Printf("Debug: For totalGrade=%.10f, totalSubjects=%d\n", test.totalGrade, test.totalSubjects)
				fmt.Printf("Expected: %.10f\n", test.expected)
				fmt.Printf("Got:      %.10f\n", result)
				fmt.Printf("Difference: %.10f\n", math.Abs(result-test.expected))
				t.Errorf("For totalGrade=%.2f, totalSubjects=%d, expected %.5f but got %.5f", test.totalGrade, test.totalSubjects, test.expected, result)
			}
		}
	}
}
func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
