package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestProcessPayment(t *testing.T) {
	app := &application{}

	// Create a valid payment request
	paymentRequest := ProccessPaymentRequest{
		FirstName:  "John",
		LastName:   "Doe",
		CardNumber: "4658587360641032",
		ExpiryDate: "12/24",
		Amount:     100.0,
		Currency:   "USD",
		CVV:        "123",
	}
	requestBody, err := json.Marshal(paymentRequest)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP request with the valid payment request
	req, err := http.NewRequest("POST", "/process-payment", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP response recorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler
	handler := http.HandlerFunc(app.proccessPayment)
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	// Check the response body
	var response ProccessPaymentResponse
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	if response.Status != "payment_paid" {
		t.Errorf("expected status to be 'payment_paid', got %v", response.Status)
	}
}

func TestRetrievePaymentDetails(t *testing.T) {
	app := &application{}

	// Simulate storing a payment
	id := "PAY-12345"
	payments[id] = PaymentDetails{
		ID:         id,
		FirstName:  "Jane",
		LastName:   "Doe",
		CardNumber: maskCardNumber("4111111111111111"),
		ExpiryDate: "12/24",
		Amount:     200.0,
		Currency:   "USD",
		Status:     "payment_paid",
		StatusCode: 10000,
	}

	// Create a new HTTP request to retrieve payment details
	req, err := http.NewRequest("GET", "/retrieve-payment?id=PAY-12345", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP response recorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler
	handler := http.HandlerFunc(app.retrievePaymentDetails)
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body
	var response PaymentDetails
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	if response.ID != id {
		t.Errorf("expected ID to be '%v', got %v", id, response.ID)
	}
	if response.FirstName != "Jane" {
		t.Errorf("expected FirstName to be 'Jane', got %v", response.FirstName)
	}
	if response.LastName != "Doe" {
		t.Errorf("expected LastName to be 'Doe', got %v", response.LastName)
	}
	if response.CardNumber != maskCardNumber("4111111111111111") {
		t.Errorf("expected CardNumber to be '%v', got %v", maskCardNumber("4111111111111111"), response.CardNumber)
	}
	if response.ExpiryDate != "12/24" {
		t.Errorf("expected ExpiryDate to be '12/24', got %v", response.ExpiryDate)
	}
	if response.Amount != 200.0 {
		t.Errorf("expected Amount to be 200.0, got %v", response.Amount)
	}
	if response.Currency != "USD" {
		t.Errorf("expected Currency to be 'USD', got %v", response.Currency)
	}
	if response.Status != "payment_paid" {
		t.Errorf("expected Status to be 'payment_paid', got %v", response.Status)
	}
	if response.StatusCode != 10000 {
		t.Errorf("expected StatusCode to be 10000, got %v", response.StatusCode)
	}
}

func TestProcessInvalidPayment(t *testing.T) {
	app := &application{}

	// Create an invalid payment request (e.g., missing required fields)
	invalidPaymentRequest := ProccessPaymentRequest{
		FirstName:  "John",
		CardNumber: "4111111111111111",
		Amount:     -100.0, // Invalid amount
		Currency:   "USD",
		CVV:        "123",
	}
	requestBody, err := json.Marshal(invalidPaymentRequest)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP request with the invalid payment request
	req, err := http.NewRequest("POST", "/process-payment", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP response recorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler
	handler := http.HandlerFunc(app.proccessPayment)
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusUnprocessableEntity {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusUnprocessableEntity)
	}

	// Check the response body for error message
	expectedError := "Field validation for 'Amount' failed on the 'gt' tag"
	if !strings.Contains(rr.Body.String(), expectedError) {
		t.Errorf("expected error message to contain '%v', got %v", expectedError, rr.Body.String())
	}
}

func TestRetrieveNonExistentPaymentDetails(t *testing.T) {
	app := &application{}

	// Create a new HTTP request to retrieve payment details with a non-existent ID
	req, err := http.NewRequest("GET", "/retrieve-payment?id=PAY-99999", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP response recorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler
	handler := http.HandlerFunc(app.retrievePaymentDetails)
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}

	// Check the response body for error message
	expectedError := "Payment not found"
	if !strings.Contains(rr.Body.String(), expectedError) {
		t.Errorf("expected error message to contain '%v', got %v", expectedError, rr.Body.String())
	}
}
