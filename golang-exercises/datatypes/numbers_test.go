package datatypes

import (
	"testing"
)

func TestAdd(t *testing.T) {
	t.Run("adds two positive numbers", func(t *testing.T) {
		result := add(1, 2)
		expected := 3
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("adds two negative numbers", func(t *testing.T) {
		result := add(-2, -3)
		expected := -5
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})
}

func TestSubtract(t *testing.T) {
	t.Run("subtracts two numbers", func(t *testing.T) {
		result := subtract(5, 3)
		expected := 2
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})
}

func TestMultiply(t *testing.T) {
	t.Run("multiplies two numbers", func(t *testing.T) {
		t.Skip("Skipping test")
		result := multiply(5, 5)
		expected := 25
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})
}

func TestDivide(t *testing.T) {
	t.Run("divides two numbers", func(t *testing.T) {
		t.Skip("Skipping test")
		result := divide(25, 5)
		expected := 5
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})
}

func TestRemainder(t *testing.T) {
	t.Run("returns the remainder using modulo", func(t *testing.T) {
		t.Skip("Skipping test")
		result := remainder(25, 5)
		expected := 0
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("returns the remainder when it is not 0", func(t *testing.T) {
		t.Skip("Skipping test")
		result := remainder(13, 5)
		expected := 3
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})
}

func TestFloatDivision(t *testing.T) {
	t.Run("returns a float", func(t *testing.T) {
		t.Skip("Skipping test")
		result := floatDivision(10, 2)
		expected := 5.0
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}

func TestStringToNumber(t *testing.T) {
	t.Run("returns an integer from a string", func(t *testing.T) {
		t.Skip("Skipping test")
		result, err := stringToNumber("1")
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		expected := 1
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("returns a negative integer from a string", func(t *testing.T) {
		t.Skip("Skipping test")
		result, err := stringToNumber("-5")
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		expected := -5
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})
}

func TestEven(t *testing.T) {
	t.Run("returns true when the number is even", func(t *testing.T) {
		t.Skip("Skipping test")
		result := isEven(6)
		expected := true
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("returns false when the number is not even", func(t *testing.T) {
		t.Skip("Skipping test")
		result := isEven(5)
		expected := false
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}

func TestOdd(t *testing.T) {
	t.Run("returns true when the number is odd", func(t *testing.T) {
		t.Skip("Skipping test")
		result := isOdd(9)
		expected := true
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("returns false when the number is not odd", func(t *testing.T) {
		t.Skip("Skipping test")
		result := isOdd(6)
		expected := false
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}
