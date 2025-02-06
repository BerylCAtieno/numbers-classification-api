// Gets number fact from NumbersAPI

package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type NumberFactResponse struct {
	Text string `json:"text"`
}

func GetNumberFact(number int) (string, error) {
	url := fmt.Sprintf("http://numbersapi.com/%d?json", number)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch fact")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var factResponse NumberFactResponse
	err = json.Unmarshal(body, &factResponse)
	if err != nil {
		return "", err
	}

	return factResponse.Text, nil
}
