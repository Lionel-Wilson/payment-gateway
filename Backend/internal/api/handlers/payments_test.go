package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Lionel-Wilson/payment-gateway/internal/api/models"
	"github.com/Lionel-Wilson/payment-gateway/pkg/utils"
	"github.com/gin-gonic/gin"

	"github.com/stretchr/testify/assert"
)

func setupTestApp() *Application {
	return &Application{}
}

func TestProcessPayment(t *testing.T) {
	tests := []struct {
		name                string
		input               models.ProcessPaymentRequest
		expectedStatusCodes []int
		expectedResponse    interface{}
	}{
		{
			name: "Valid Payment",
			input: models.ProcessPaymentRequest{
				FirstName:    "John",
				LastName:     "Doe",
				CardNumber:   "4658587360641032",
				ExpiryDate:   "12/24",
				Amount:       100.0,
				CurrencyCode: "USD",
				CVV:          "123",
			},
			expectedStatusCodes: []int{http.StatusCreated, http.StatusPaymentRequired},
			expectedResponse:    "payment_paid",
		},
		{
			name: "Invalid Card Number",
			input: models.ProcessPaymentRequest{
				FirstName:    "John",
				LastName:     "Doe",
				CardNumber:   "1234567890123456",
				ExpiryDate:   "12/29",
				Amount:       500,
				CurrencyCode: "USD",
				CVV:          "123",
			},
			expectedStatusCodes: []int{http.StatusUnprocessableEntity},
			expectedResponse:    "Validation failed",
		},
		{
			name: "Invalid Payment Request",
			input: models.ProcessPaymentRequest{
				FirstName:    "John",
				CardNumber:   "4111111111111111",
				Amount:       -100.0,
				CurrencyCode: "USD",
				CVV:          "123",
			},
			expectedStatusCodes: []int{http.StatusUnprocessableEntity},
			expectedResponse:    "Validation failed",
		},
	}

	app := setupTestApp()
	router := gin.New()
	router.POST("/api/v1/payments", app.ProcessPayment)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reqBody, _ := json.Marshal(tt.input)
			req, _ := http.NewRequest("POST", "/api/v1/payments", bytes.NewBuffer(reqBody))
			req.Header.Set("Content-Type", "application/json")

			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			assert.Contains(t, tt.expectedStatusCodes, rr.Code)

			if rr.Code == http.StatusCreated {
				var response models.ProcessPaymentResponse
				err := json.Unmarshal(rr.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResponse, response.Status)
			} else if rr.Code == http.StatusPaymentRequired {
				var response models.ProcessPaymentResponse
				err := json.Unmarshal(rr.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, "payment_declined", response.Status)
			} else {
				var errorResponse utils.ErrorResponse
				err := json.Unmarshal(rr.Body.Bytes(), &errorResponse)
				assert.NoError(t, err)
				assert.Contains(t, errorResponse.Message, tt.expectedResponse)
			}
		})
	}
}

func TestRetrievePaymentDetails(t *testing.T) {
	tests := []struct {
		name               string
		setupPayments      map[string]models.PaymentDetails
		paymentID          string
		expectedStatusCode int
		expectedResponse   interface{}
	}{
		{
			name: "Valid Payment",
			setupPayments: map[string]models.PaymentDetails{
				"PAY-12345": {
					ID:           "PAY-12345",
					FirstName:    "Jane",
					LastName:     "Doe",
					CardNumber:   utils.MaskCardNumber("4111111111111111"),
					ExpiryDate:   "12/24",
					Amount:       200.0,
					CurrencyCode: "USD",
					Status:       "payment_paid",
					StatusCode:   10000,
				},
			},
			paymentID:          "PAY-12345",
			expectedStatusCode: http.StatusOK,
			expectedResponse: models.PaymentDetails{
				ID:           "PAY-12345",
				FirstName:    "Jane",
				LastName:     "Doe",
				CardNumber:   utils.MaskCardNumber("4111111111111111"),
				ExpiryDate:   "12/24",
				Amount:       200.0,
				CurrencyCode: "USD",
				Status:       "payment_paid",
				StatusCode:   10000,
			},
		},
		{
			name:               "Non-Existent Payment",
			setupPayments:      map[string]models.PaymentDetails{},
			paymentID:          "PAY-99999",
			expectedStatusCode: http.StatusNotFound,
			expectedResponse:   "Payment not found",
		},
	}

	app := setupTestApp()
	router := gin.New()
	router.GET("/api/v1/payments/:id", app.RetrievePayment)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mu.Lock()
			payments = tt.setupPayments
			mu.Unlock()

			req, _ := http.NewRequest("GET", "/api/v1/payments/"+tt.paymentID, nil)
			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			assert.Equal(t, tt.expectedStatusCode, rr.Code)

			if tt.expectedStatusCode == http.StatusOK {
				var response models.PaymentDetails
				err := json.Unmarshal(rr.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResponse, response)
			} else {
				var errorResponse utils.ErrorResponse
				err := json.Unmarshal(rr.Body.Bytes(), &errorResponse)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResponse, errorResponse.Message)
			}
		})
	}
}

func TestAllPayments(t *testing.T) {
	tests := []struct {
		name               string
		setupPayments      map[string]models.PaymentDetails
		expectedStatusCode int
		expectedResponse   interface{}
	}{
		{
			name:               "No Payments",
			setupPayments:      map[string]models.PaymentDetails{},
			expectedStatusCode: http.StatusNotFound,
			expectedResponse:   "No payments available",
		},
		{
			name: "With Payments",
			setupPayments: map[string]models.PaymentDetails{
				"PAY-12345": {
					ID:           "PAY-12345",
					FirstName:    "Jane",
					LastName:     "Doe",
					CardNumber:   utils.MaskCardNumber("4111111111111111"),
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
					CardNumber:   utils.MaskCardNumber("4222222222222222"),
					ExpiryDate:   "11/23",
					Amount:       150.0,
					CurrencyCode: "EUR",
					Status:       "payment_paid",
					StatusCode:   10000,
				},
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse: []models.PaymentDetails{
				{
					ID:           "PAY-12345",
					FirstName:    "Jane",
					LastName:     "Doe",
					CardNumber:   utils.MaskCardNumber("4111111111111111"),
					ExpiryDate:   "12/24",
					Amount:       200.0,
					CurrencyCode: "USD",
					Status:       "payment_paid",
					StatusCode:   10000,
				},
				{
					ID:           "PAY-67890",
					FirstName:    "John",
					LastName:     "Smith",
					CardNumber:   utils.MaskCardNumber("4222222222222222"),
					ExpiryDate:   "11/23",
					Amount:       150.0,
					CurrencyCode: "EUR",
					Status:       "payment_paid",
					StatusCode:   10000,
				},
			},
		},
	}

	app := setupTestApp()
	router := gin.New()
	router.GET("/api/v1/payments", app.AllPayments)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mu.Lock()
			payments = tt.setupPayments
			mu.Unlock()

			req, _ := http.NewRequest("GET", "/api/v1/payments", nil)
			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			assert.Equal(t, tt.expectedStatusCode, rr.Code)

			if tt.expectedStatusCode == http.StatusOK {
				var response []models.PaymentDetails
				err := json.Unmarshal(rr.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResponse, response)
			} else {
				var errorResponse utils.ErrorResponse
				err := json.Unmarshal(rr.Body.Bytes(), &errorResponse)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResponse, errorResponse.Message)
			}
		})
	}
}
