package datatypes

import (
	"strconv"
)

// Add function
func add(a int, b int) int {
	// return the result of adding a and b
	return a + b
}

// Subtract function
func subtract(a int, b int) int {
	// return the result of subtracting b from a
	return a - b
}

// Multiply function
func multiply(a int, b int) int {
	// return the result of multiplying a times b
	return a * b
}

// Divide function
func divide(a int, b int) int {
	// return the result of dividing a by b
	return a / b
}

// Remainder function
func remainder(a int, b int) int {
	// return the remainder of dividing a by b using the modulo operator
	return a % b
}

// FloatDivision function
func floatDivision(a float64, b float64) float64 {
	// return the result of dividing a by b as a float
	return a / b
}

// StringToNumber function
func stringToNumber(s string) (int, error) {
	// return the result of converting a string into an integer
	n, err := strconv.Atoi(s)

	if err != nil {
		return 0, err
	}

	return n, nil
}

// IsEven function
func isEven(number int) bool {
	// return true if the number is even
	return number%2 == 0
}

// IsOdd function
func isOdd(number int) bool {
	// return true if the number is odd
	return number%2 != 0
}
