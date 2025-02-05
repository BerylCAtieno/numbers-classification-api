package main

import (
	"encoding/json"
	"io"
	"math"
	"net/http"
)

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

    // Count the number of digits
    for temp > 0 {
        digits++
        temp /= 10
    }

    // Calculate sum of each digit raised to the power of digits
    temp = number
    for temp > 0 {
        digit := temp % 10
        sum += int(math.Pow(float64(digit), float64(digits)))
        temp /= 10
    }

    return sum == original
}

// Check if a number is a perfect number
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

// Structure to hold the response from the Numbers API
type NumberFactResponse struct {
	Text string `json:"text"`
}

// Handler to fetch fun fact about a number
func getNumberFact(w http.ResponseWriter, r *http.Request) {
	// Extract the number from the URL
	number := r.URL.Query().Get("number")
	if number == "" {
		http.Error(w, "Please provide a number as query parameter", http.StatusBadRequest)
		return
	}

	// Call the Numbers API
	url := fmt.Sprintf("http://numbersapi.com/%s?json", number)
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Failed to fetch data from Numbers API", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Check if the response is OK
	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Failed to get valid response from Numbers API", resp.StatusCode)
		return
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response body", http.StatusInternalServerError)
		return
	}

	// Parse the response from Numbers API into our struct
	var factResponse NumberFactResponse
	err = json.Unmarshal(body, &factResponse)
	if err != nil {
		http.Error(w, "Failed to parse Numbers API response", http.StatusInternalServerError)
		return
	}

	// Set the content-type as JSON and send the response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Return the fun fact as a JSON response
	json.NewEncoder(w).Encode(factResponse)
}
