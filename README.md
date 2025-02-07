# Numbers Classifier API

A simple RESTful API for classifying numbers based on various properties such as primality, Armstrong numbers, perfect numbers, and more.

## Features
- Check if a number is prime, Armstrong, or perfect.
- Determine if a number is even or odd.
- Compute the sum of digits of a number.
- Retrieve fun facts about numbers.

## Table of Contents
- [Installation](#installation)
- [Running the API](#running-the-api)
- [API Endpoints](#api-endpoints)
- [Testing the API](#testing-the-api)

## Installation
1. Clone the repository:
   ```sh
   git clone https://github.com/BerylCAtieno/numbers-classification-api.git
   cd numbers-classifier-api
   ```
2. Install dependencies:
   ```sh
   go mod tidy
   ```

## Running the API
### Using Go
Run the application locally with:
```sh
go run main.go
```

## API Endpoints
### Classify a Number
- **URL:** `/api/classify-number?number={number}`
- **Method:** `GET`
- **Parameters:**
  - `number` (integer) - The number to classify.
- **Response:**
  ```json
  {
    "number": 153,
    "isPrime": false,
    "isPerfect": false,
    "properties": ["armstrong", "odd"],
    "digitSum": 9,
    "funFact": "153 is an Armstrong number."
  }
  ```

## Testing the API
You can test the API using `curl`:
```sh
curl "https://numbers-classification-api-3b6h.onrender.com/api/classify-number?number=153"
```
Or using Postman by sending a `GET` request to `http://localhost:8080/api/classify-number?number=153`.


Author

`Beryl Atieno`

GitHub: [BerylCAtieno](https://github.com/BerylCAtieno)