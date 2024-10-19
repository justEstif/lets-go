package datatypes

import (
	"testing"
)

func TestConcatenateExample(t *testing.T) {
	result := concatenateExample("Ruby")
	expected := "Classic Ruby"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestConcatenate(t *testing.T) {
	t.Run("returns 'Hello world!'", func(t *testing.T) {
		result := concatenate("world")
		expected := "Hello world!"
		if result != expected {
			t.Errorf("Expected %q, got %q", expected, result)
		}
	})

	t.Run("returns 'Hello universe!'", func(t *testing.T) {
		result := concatenate("universe")
		expected := "Hello universe!"
		if result != expected {
			t.Errorf("Expected %q, got %q", expected, result)
		}
	})
}

func TestSubstrings(t *testing.T) {
	t.Run("returns the first 4 letters of the word", func(t *testing.T) {
		result := substrings("chocolate")
		expected := "choc"
		if result != expected {
			t.Errorf("Expected %q, got %q", expected, result)
		}
	})
}

func TestCapitalize(t *testing.T) {
	t.Run("capitalizes a word", func(t *testing.T) {
		result := capitalize("paris")
		expected := "Paris"
		if result != expected {
			t.Errorf("Expected %q, got %q", expected, result)
		}
	})

	t.Run("only capitalizes the first word if there are multiple words", func(t *testing.T) {
		result := capitalize("miami in the summer")
		expected := "Miami in the summer"
		if result != expected {
			t.Errorf("Expected %q, got %q", expected, result)
		}
	})

	t.Run("leaves an already capitalized word as is", func(t *testing.T) {
		result := capitalize("London")
		expected := "London"
		if result != expected {
			t.Errorf("Expected %q, got %q", expected, result)
		}
	})
}

func TestUppercase(t *testing.T) {
	t.Run("uppercases a word", func(t *testing.T) {
		result := uppercase("small")
		expected := "SMALL"
		if result != expected {
			t.Errorf("Expected %q, got %q", expected, result)
		}
	})

	t.Run("uppercases multiple words", func(t *testing.T) {
		result := uppercase("make me bigger")
		expected := "MAKE ME BIGGER"
		if result != expected {
			t.Errorf("Expected %q, got %q", expected, result)
		}
	})
}

func TestDowncase(t *testing.T) {
	t.Run("downcases a word", func(t *testing.T) {
		result := downcase("LARGE")
		expected := "large"
		if result != expected {
			t.Errorf("Expected %q, got %q", expected, result)
		}
	})

	t.Run("downcases multiple words", func(t *testing.T) {
		result := downcase("MAKE ME SMALLER")
		expected := "make me smaller"
		if result != expected {
			t.Errorf("Expected %q, got %q", expected, result)
		}
	})
}

func TestEmptyString(t *testing.T) {
	t.Run("returns true if string is empty", func(t *testing.T) {
		result := emptyString("")
		expected := true
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("returns false if string is not empty", func(t *testing.T) {
		result := emptyString("wow")
		expected := false
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}

func TestStringLength(t *testing.T) {
	t.Run("returns the length of a word", func(t *testing.T) {
		result := stringLength("longitude")
		expected := 9
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("returns the length of a string with multiple words", func(t *testing.T) {
		result := stringLength("i am quite long")
		expected := 15
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})
}

func TestReverse(t *testing.T) {
	t.Run("reverses a word", func(t *testing.T) {
		result := reverse("desrever")
		expected := "reversed"
		if result != expected {
			t.Errorf("Expected %q, got %q", expected, result)
		}
	})

	t.Run("reverses multiple words", func(t *testing.T) {
		result := reverse("dnuora kcab")
		expected := "back around"
		if result != expected {
			t.Errorf("Expected %q, got %q", expected, result)
		}
	})
}

func TestSpaceRemover(t *testing.T) {
	t.Run("removes a single space", func(t *testing.T) {
		result := spaceRemover("white space")
		expected := "whitespace"
		if result != expected {
			t.Errorf("Expected %q, got %q", expected, result)
		}
	})

	t.Run("removes multiple spaces", func(t *testing.T) {
		result := spaceRemover("many white spaces")
		expected := "manywhitespaces"
		if result != expected {
			t.Errorf("Expected %q, got %q", expected, result)
		}
	})
}
