package main

import "testing"

func TestSumOfDigits(t *testing.T) {
	result := sumOfDigits(51)
	expected := 6

	if result != expected {
		t.Errorf("SumOfDigits(51) = %d; want %d", result, expected)
	}

}

func TestEvenOrOdd(t *testing.T) {
	result := isEvenOrOdd(4)
	expected := "even"

	if result != expected {
		t.Errorf("isEvenOrOdd(4) = %s; want %s", result, expected)
	}
}

func TestIsArmstrong(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected bool
	}{
		{"Armstrong number 153", 153, true},
		{"Armstrong number 9474", 9474, true},
		{"Not an Armstrong number 9475", 9475, false},
		{"Armstrong number 370", 370, true},
		{"Armstrong number 407", 407, true},
		{"Single-digit Armstrong number 5", 5, true},
		{"Armstrong number 0", 0, true},
		{"Non-Armstrong number 123", 123, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := isArmstrong(tc.input)
			if result != tc.expected {
				t.Errorf("isArmstrong(%d) = %v; want %v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestIsPerfect(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected bool
	}{
		{"Perfect number 6", 6, true},
		{"Perfect number 28", 28, true},
		{"Not a perfect number 12", 12, false},
		{"Perfect number 496", 496, true},
		{"Perfect number 8128", 8128, true},
		{"Perfect number 33550336", 33550336, true},
		{"Not a perfect number 7", 7, false},
		{"Not a perfect number 0", 0, false},
		{"Not a perfect number 1", 1, false},
		{"Not a perfect number 27", 27, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := isPerfect(tc.input)
			if result != tc.expected {
				t.Errorf("isPerfect(%d) = %v; want %v", tc.input, result, tc.expected)
			}
		})
	}
}
