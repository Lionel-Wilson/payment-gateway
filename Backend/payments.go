package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/go-playground/validator/v10"
)

type ProccessPaymentRequest struct {
	FirstName  string  `json:"firstName" validate:"required,alpha"`
	LastName   string  `json:"lastName" validate:"required,alpha"`
	CardNumber string  `json:"cardNumber" validate:"required,credit_card"`
	ExpiryDate string  `json:"expiryDate" validate:"required,expirydate"` // Custom validation to ensure it matches "MM/YY" format
	Amount     float64 `json:"amount" validate:"required,gt=0"`
	Currency   string  `json:"currency" validate:"required,len=3,alpha"`
	CVV        string  `json:"cvv" validate:"required,len=3,numeric"`
}

type ProccessPaymentResponse struct {
	ID              string `json:"id"`
	Status          string `json:"status"`
	ResponseSummary string `json:"responseSummary"`
}

type PaymentDetails struct {
	ID         string  `json:"id"`
	FirstName  string  `json:"firstName"`
	LastName   string  `json:"lastName"`
	CardNumber string  `json:"cardNumber"`
	ExpiryDate string  `json:"expiryDate"`
	Amount     float64 `json:"amount"`
	Currency   string  `json:"currency"`
	Status     string  `json:"status"`
	StatusCode int     `json:"statusCode"`
}

var validate *validator.Validate
var payments = make(map[string]PaymentDetails)
var mu sync.Mutex

func (app *application) proccessPayment(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	var paymentDetails ProccessPaymentRequest

	err = json.Unmarshal(body, &paymentDetails)
	if err != nil {
		app.serverError(w, err)
		return
	}

	validate = validator.New()
	validate.RegisterValidation("expirydate", expiryDateValidation)

	err = validate.Struct(paymentDetails)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	id, status, statusCode, summary := simulateBank(paymentDetails.Amount)

	mu.Lock()
	payments[id] = PaymentDetails{
		ID:         id,
		FirstName:  paymentDetails.FirstName,
		LastName:   paymentDetails.LastName,
		CardNumber: maskCardNumber(paymentDetails.CardNumber),
		ExpiryDate: paymentDetails.ExpiryDate,
		Amount:     paymentDetails.Amount,
		Currency:   paymentDetails.Currency,
		Status:     status,
		StatusCode: statusCode,
	}
	mu.Unlock()

	response := ProccessPaymentResponse{
		ID:              id,
		Status:          status,
		ResponseSummary: summary,
	}

	if status == "payment_paid" {
		w.WriteHeader(http.StatusCreated)
	} else {
		w.WriteHeader(http.StatusPaymentRequired)
	}

	var jsonResponse []byte
	jsonResponse, err = json.Marshal(response)
	if err != nil {
		app.serverError(w, err)
	}

	w.Write(jsonResponse)
}

func (app *application) retrievePaymentDetails(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Invalid id provided", http.StatusBadRequest)
		return
	}

	mu.Lock()
	payment, exists := payments[id]
	mu.Unlock()

	if !exists {
		http.Error(w, "Payment not found", http.StatusNotFound)
		return
	}

	var jsonResponse []byte
	jsonResponse, err := json.Marshal(payment)
	if err != nil {
		app.serverError(w, err)
	}

	w.Write(jsonResponse)
}

func expiryDateValidation(fl validator.FieldLevel) bool {
	expiryDate := fl.Field().String()
	match, _ := regexp.MatchString(`^(0[1-9]|1[0-2])\/([0-9]{2})$`, expiryDate)
	return match
}

func simulateBank(amount float64) (string, string, int, string) {
	customerAccountRemainingBalance := 800

	id := fmt.Sprintf("PAY-%d", time.Now().UnixNano())

	if amount <= float64(customerAccountRemainingBalance) {
		status := "payment_paid"
		statusCode := 10000
		summary := "Approved"
		return id, status, statusCode, summary
	} else {
		status := "payment_declined"
		statusCode := 50280
		summary := "Insufficient funds"
		return id, status, statusCode, summary
	}
}

func maskCardNumber(cardNumber string) string {
	return strings.Repeat("*", len(cardNumber)-4) + cardNumber[len(cardNumber)-4:]
}
