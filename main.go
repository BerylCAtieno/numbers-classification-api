package main


// Check if a number is prime
// Check if a number is an Armstrong number
// Check if a number is a perfect number
// Determine odd/even

func isEvenOrOdd(number int) string {
	if number%2 == 0 {
		return "even"
	} else {
		return "odd"
	}
}
// Compute the sum of digits

func sumOfDigits(number int) int {
	sum := 0
    if number < 0 {
        number = -number // Handle negative numbers
    }
    for number > 0 {
        sum += number % 10 // Extract last digit and add to sum
        number /= 10       // Remove last digit
    }
    return sum
}

// Fetch number fun facts