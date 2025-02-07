package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"numbers-classifier-api/models"
	"numbers-classifier-api/utils"
)

func ClassifyNumberHandler(w http.ResponseWriter, r *http.Request) {
	// Enable CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	numberStr := r.URL.Query().Get("number")
	number, err := strconv.Atoi(numberStr)
	if err != nil {
		response := map[string]interface{}{
			"number": numberStr,
			"error":  true,
		}
		w.WriteHeader(http.StatusBadRequest) // Set 400 status
		json.NewEncoder(w).Encode(response)
		return
	}
	

	properties := []string{}
	if utils.IsArmstrong(number) {
		properties = append(properties, "armstrong")
	}
	properties = append(properties, utils.IsEvenOrOdd(number))

	fact, err := GetNumberFact(number)
	if err != nil {
		fact = "Fun fact unavailable"
	}

	response := models.NumberResponse{
		Number:     number,
		IsPrime:    utils.IsPrime(number),
		IsPerfect:  utils.IsPerfect(number),
		Properties: properties,
		DigitSum:   utils.SumOfDigits(number),
		FunFact:    fact,
	}

	json.NewEncoder(w).Encode(response)
}
