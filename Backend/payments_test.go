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
		FirstName:    "John",
		LastName:     "Doe",
		CardNumber:   "4658587360641032",
		ExpiryDate:   "12/24",
		Amount:       100.0,
		CurrencyCode: "USD",
		CVV:          "123",
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
		ID:           id,
		FirstName:    "Jane",
		LastName:     "Doe",
		CardNumber:   maskCardNumber("4111111111111111"),
		ExpiryDate:   "12/24",
		Amount:       200.0,
		CurrencyCode: "USD",
		Status:       "payment_paid",
		StatusCode:   10000,
	}

	// Create a new HTTP request to retrieve payment details
	req, err := http.NewRequest("GET", "/payments/PAY-12345", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP response recorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler
	handler := http.HandlerFunc(app.retrievePayment)
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
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
	if response.CurrencyCode != "USD" {
		t.Errorf("expected Currency Code to be 'USD', got %v", response.CurrencyCode)
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
		FirstName:    "John",
		CardNumber:   "4111111111111111",
		Amount:       -100.0, // Invalid amount
		CurrencyCode: "USD",
		CVV:          "123",
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
	expectedErrors := []string{
		"LastName is required",
		"ExpiryDate is required",
		"Amount must be greater than 0",
	}
	var responseBody map[string][]string
	err = json.Unmarshal(rr.Body.Bytes(), &responseBody)
	if err != nil {
		t.Fatalf("failed to unmarshal response body: %v", err)
	}

	errors, ok := responseBody["errors"]
	if !ok {
		t.Fatalf("response body does not contain error key")
	}

	for _, expectedError := range expectedErrors {
		found := false
		for _, err := range errors {
			if strings.Contains(err, expectedError) {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("expected error message to contain '%v', got %v", expectedError, errors)
		}
	}
}

func TestRetrieveNonExistentPaymentDetails(t *testing.T) {
	app := &application{}

	// Create a new HTTP request to retrieve payment details with a non-existent ID
	req, err := http.NewRequest("GET", "/payments/PAY-99999", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP response recorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler
	handler := http.HandlerFunc(app.retrievePayment)
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

func TestAllPaymentsNoPayments(t *testing.T) {
	app := &application{}

	// Clear the payments map to simulate no payments
	mu.Lock()
	payments = make(map[string]PaymentDetails)
	mu.Unlock()

	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/payments", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP response recorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler
	handler := http.HandlerFunc(app.allPayments)
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}

	// Check the response body for error message
	expectedError := "No payments available"
	if !strings.Contains(rr.Body.String(), expectedError) {
		t.Errorf("expected error message to contain '%v', got %v", expectedError, rr.Body.String())
	}
}

func TestAllPaymentsWithPayments(t *testing.T) {
	app := &application{}

	// Add some payments to the payments map
	mu.Lock()
	payments["PAY-12345"] = PaymentDetails{
		ID:           "PAY-12345",
		FirstName:    "Jane",
		LastName:     "Doe",
		CardNumber:   maskCardNumber("4111111111111111"),
		ExpiryDate:   "12/24",
		Amount:       200.0,
		CurrencyCode: "USD",
		Status:       "payment_paid",
		StatusCode:   10000,
	}
	payments["PAY-67890"] = PaymentDetails{
		ID:           "PAY-67890",
		FirstName:    "John",
		LastName:     "Smith",
		CardNumber:   maskCardNumber("4222222222222222"),
		ExpiryDate:   "11/23",
		Amount:       150.0,
		CurrencyCode: "EUR",
		Status:       "payment_paid",
		StatusCode:   10000,
	}
	mu.Unlock()

	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/payments", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP response recorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler
	handler := http.HandlerFunc(app.allPayments)
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	var response []PaymentDetails
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	// Verify the payments
	if len(response) != 2 {
		t.Errorf("expected 2 payments, got %v", len(response))
	}
	expectedPayments := map[string]PaymentDetails{
		"PAY-12345": {
			ID:           "PAY-12345",
			FirstName:    "Jane",
			LastName:     "Doe",
			CardNumber:   maskCardNumber("4111111111111111"),
			ExpiryDate:   "12/24",
			Amount:       200.0,
			CurrencyCode: "USD",
			Status:       "payment_paid",
			StatusCode:   10000,
		},
		"PAY-67890": {
			ID:           "PAY-67890",
			FirstName:    "John",
			LastName:     "Smith",
			CardNumber:   maskCardNumber("4222222222222222"),
			ExpiryDate:   "11/23",
			Amount:       150.0,
			CurrencyCode: "EUR",
			Status:       "payment_paid",
			StatusCode:   10000,
		},
	}
	for _, payment := range response {
		expected, ok := expectedPayments[payment.ID]
		if !ok {
			t.Errorf("unexpected payment ID: %v", payment.ID)
		}
		if expected != payment {
			t.Errorf("expected payment details %v, got %v", expected, payment)
		}
	}
}
