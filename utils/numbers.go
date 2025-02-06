package utils

import "math"

// isPrime checks if a number is prime
func isPrime(number int) bool {
    if number < 2 {
        return false
    }
    for i := 2; i*i <= number; i++ {
        if number%i == 0 {
            return false
        }
    }
    return true
}


// isArmstrong checks if a number is an Armstrong number
func isArmstrong(number int) bool {
	original := number
    sum := 0
    digits := 0
    temp := number

    for temp > 0 {
        digits++
        temp /= 10
    }

    temp = number
    for temp > 0 {
        digit := temp % 10
        sum += int(math.Pow(float64(digit), float64(digits)))
        temp /= 10
    }

    return sum == original
}

// isPerfect checks if a number is a perfect number
func isPerfect(number int) bool {
	if number < 2 {
        return false
    }

    sum := 1 // Start with 1 as it's a proper divisor of every number
    for i := 2; i*i <= number; i++ {
        if number%i == 0 {
            sum += i
            if i != number/i { // Avoid adding square roots twice (e.g., for 16: add 2 and 8, but not 4 twice)
                sum += number / i
            }
        }
    }
    return sum == number
}

// isEvenOrOdd determines if a number is odd/even
func isEvenOrOdd(number int) string {
	if number%2 == 0 {
		return "even"
	} else {
		return "odd"
	}
}

// sumOfDigits computes the sum of digits in a number
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
