# Number Classification API

## Overview

The **Number Classification API** is a simple RESTful web service that classifies a given number by its properties.

## Features

- Determines if a number is a **prime number**.
- Checks whether the number is **even or odd**.
- Computes the **sum of the digits** of the number.
- Retrieves a **fact about the number** from [Numbers API](http://numbersapi.com/).

## API Endpoint

- `GET /api/classify-number?number=42/math`

### Request Example

```bash
GET http://localhost:8000/api/classify-number?number=371/math
```

### Response Example

```json
{
  "number": 371,
  "is_prime": false,
  "is_perfect": false,
  "properties": ["armstrong", "odd"],
  "digit_sum": 11,
  "fun_fact": "371 is a pokemon ⚡, just kidding."
}
```

## Installation & Setup

### Prerequisites

- [Go](https://go.dev/) installed on your machine.
- IDE eg.vs code

### Steps

1. **Clone the repository:**

   ```sh
   git clone https://github.com/ikennarichard/number-classification-api.git
   cd number-classification-api
   ```

2. **Run the API server:**

   ```sh
   go run main.go
   ```

3. The server will start on `http://localhost:8000`

## Error Handling

If an invalid or missing number is provided, the API returns a JSON error response:

```json
{
  "number": "abc",
  "error": true,
}
```
