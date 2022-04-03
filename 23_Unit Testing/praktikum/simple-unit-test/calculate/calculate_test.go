package calculate

import (
	"testing"
)

func TestAddition(t *testing.T) {
	if Addition(20, 30) != 50 {
		t.Error("Expected (20)+(30) to equal 50")
	}

	if Addition(-30, 10) != -20 {
		t.Error("Expected (-30)+(10) to equal -20")
	}
}

func TestSubtraction(t *testing.T) {
	if Subtraction(20, 30) != -10 {
		t.Error("Expected (20)-(30) to equal -10")
	}

	if Subtraction(-30, 10) != -40 {
		t.Error("Expected (-30)-(10) to equal -40")
	}
}

func TestDivision(t *testing.T) {
	if Division(8, 2) != 4 {
		t.Error("Expected (8)/(2) to equal 4")
	}

	if Division(30, 6) != 5 {
		t.Error("Expected (30)/(6) to equal 5")
	}
}

func TestMultiplication(t *testing.T) {
	if Multiplication(8, 8) != 64 {
		t.Error("Expected (8)/(8) to equal 64")
	}

	if Multiplication(2, 6) != 12 {
		t.Error("Expected (2)/(6) to equal 12")
	}
}
