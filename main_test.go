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


