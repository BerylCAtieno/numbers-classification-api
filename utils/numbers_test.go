package utils

import (
	"testing"
)


func TestSumOfDigits(t *testing.T) {
	result := SumOfDigits(51)
	expected := 6

	if result != expected {
		t.Errorf("SumOfDigits(51) = %d; want %d", result, expected)
	}

}

func TestEvenOrOdd(t *testing.T) {
	result := IsEvenOrOdd(4)
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
			result := IsArmstrong(tc.input)
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
			result := IsPerfect(tc.input)
			if result != tc.expected {
				t.Errorf("isPerfect(%d) = %v; want %v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestIsPrime(t *testing.T) {
    testCases := []struct {
        name     string
        input    int
        expected bool
    }{
        {"Prime number 2", 2, true},
        {"Prime number 3", 3, true},
        {"Not a prime number 4", 4, false},
        {"Prime number 5", 5, true},
        {"Not a prime number 10", 10, false},
        {"Prime number 11", 11, true},
        {"Prime number 13", 13, true},
        {"Not a prime number 16", 16, false},
        {"Prime number 17", 17, true},
        {"Prime number 19", 19, true},
        {"Prime number 23", 23, true},
        {"Prime number 29", 29, true},
        {"Prime number 97", 97, true},
        {"Not a prime number 100", 100, false},
        {"Not a prime number 1", 1, false},
        {"Not a prime number 0", 0, false},
        {"Not a prime number -5", -5, false},
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            result := IsPrime(tc.input)
            if result != tc.expected {
                t.Errorf("isPrime(%d) = %v; want %v", tc.input, result, tc.expected)
            }
        })
    }
}

