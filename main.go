package main

import (
	"encoding/json"
	"fmt"
	"math"

	"net/http"
	"strconv"
	"strings"
)

type Response struct {
	Number     int      `json:"number"`
	IsPrime    bool    `json:"is_prime"`
	IsPerfect  bool     `json:"is_perfect"`
	Properties []string `json:"properties"`
	DigitSum   int      `json:"digit_sum"`
	FunFact    string   `json:"fun_fact"`
}


type ErrorResponse struct {
	Number string `json:"number"`
	Error  bool   `json:"error"`
}

func main() {
	http.HandleFunc("/api/classify-number", routeHandler)
	fmt.Println("server running on port 8000...")
	http.ListenAndServe(":8000", nil)
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func isPerfect(n int) bool {
	sum := 1
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			sum += i
			if i != n/i {
				sum += n / i
			}
		}
	}
	return sum == n && n != 1
}

func isArmstrong(n int) bool {
	original := n
	sum := 0
	numDigits := int(math.Log10(float64(n))) + 1

	for n > 0 {
		digit := n % 10
		sum += int(math.Pow(float64(digit), float64(numDigits)))
		n /= 10
	}
	return sum == original
}

func calculateDigitSum(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}


func fetchFunFact(number int) string {
	url := fmt.Sprintf("http://numbersapi.com/%d/math", number)
	resp, err := http.Get(url)
	if err != nil {
		return "Fun fact unavailable"
	}
	defer resp.Body.Close()

	var fact strings.Builder // grow string dynamically
	
	buf := make([]byte, 512) // a buffer to read the chunks of data 
	for {
		n, _ := resp.Body.Read(buf) // read up to 512 bytes
		if n == 0 {
			break
		}
		fact.Write(buf[:n]) // add the data to fact builder
	}

	return fact.String()
}


func routeHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "https://*, http://")
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		w.WriteHeader(405)
		json.NewEncoder(w).Encode(ErrorResponse{Number: "Invalid request method", Error: true})
		return
	}

	numberStr := r.URL.Query().Get("number")
	if numberStr == "" {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(ErrorResponse{Number: numberStr, Error: true})
		return
	}

	number, err := strconv.Atoi(numberStr)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(ErrorResponse{Number: numberStr, Error: true})
		return
	}

	// properties
	properties := []string{}
	if isArmstrong(number) {
		properties = append(properties, "armstrong")
	}
	if number%2 == 0 {
		properties = append(properties, "even")
	} else {
		properties = append(properties, "odd")
	}

	digitSum := calculateDigitSum(number)
	funFact := fetchFunFact(number)

	// response
	response := Response{
		Number:      number,
		IsPrime:     isPrime(number),
		IsPerfect:   isPerfect(number),
		Properties:  properties,
		DigitSum:    digitSum,
		FunFact:     funFact,
	}
	json.NewEncoder(w).Encode(response)
}