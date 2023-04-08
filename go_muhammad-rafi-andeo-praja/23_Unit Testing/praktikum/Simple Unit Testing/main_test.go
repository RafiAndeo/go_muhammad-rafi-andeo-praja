package testproject

import (
	"testing"
)

func TestAddition(t *testing.T) {
	if addition(1, 2) != 3 {
		t.Error("Expected 1 + 2 to equal 3")
	}
}

func TestSubtraction(t *testing.T) {
	if subtraction(2, 1) != 1 {
		t.Error("Expected 2 - 1 to equal 1")
	}
}

func TestMultiplication(t *testing.T) {
	if multiplication(2, 2) != 4 {
		t.Error("Expected 2 * 2 to equal 4")
	}
}

func TestDivision(t *testing.T) {
	if division(4, 2) != 2 {
		t.Error("Expected 4 / 2 to equal 2")
	}
}
