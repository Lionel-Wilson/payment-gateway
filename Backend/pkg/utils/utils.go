package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/Lionel-Wilson/payment-gateway/internal/api/models"
	"github.com/gin-gonic/gin"
)

// SimulateBank simulates a bank's response to a payment request.
// It returns a unique payment ID, status, status code, and a response summary.
// The status is randomly chosen between "payment_paid" and "payment_declined".
func SimulateBank() (string, string, int, string) {
	statuses := []string{"payment_paid", "payment_declined"}
	status := statuses[rand.Intn(len(statuses))]

	id := fmt.Sprintf("PAY-%d", time.Now().UnixNano())

	if status == "payment_paid" {
		statusCode := 10000
		summary := "Approved"
		return id, status, statusCode, summary
	} else {
		statusCode := 50280
		summary := "Insufficient funds"
		return id, status, statusCode, summary
	}
}

// MaskCardNumber masks all but the last four digits of a credit card number.
// It replaces the initial digits with asterisks (*).
func MaskCardNumber(cardNumber string) string {
	return strings.Repeat("*", len(cardNumber)-4) + cardNumber[len(cardNumber)-4:]
}

// TrimWhitespace trims leading and trailing whitespace from all string fields in ProcessPaymentRequest.
func TrimWhitespace(paymentDetails *models.ProcessPaymentRequest) {
	paymentDetails.FirstName = strings.TrimSpace(paymentDetails.FirstName)
	paymentDetails.LastName = strings.TrimSpace(paymentDetails.LastName)
	paymentDetails.CardNumber = strings.TrimSpace(paymentDetails.CardNumber)
	paymentDetails.ExpiryDate = strings.TrimSpace(paymentDetails.ExpiryDate)
	paymentDetails.CurrencyCode = strings.TrimSpace(paymentDetails.CurrencyCode)
	paymentDetails.CVV = strings.TrimSpace(paymentDetails.CVV)
}

// NewErrorResponse creates a new error response with the provided status code, message, and errors.
// It sends a JSON response with these details to the client.
func NewErrorResponse(c *gin.Context, statusCode int, message string, errors []string) {
	c.JSON(statusCode, ErrorResponse{
		StatusCode: statusCode,
		Message:    message,
		Errors:     errors,
	})
}

// ErrorResponse represents the structure of an error response.
// It contains a status code, a message, and an optional list of errors.
type ErrorResponse struct {
	StatusCode int      `json:"statusCode" example:"422"`
	Message    string   `json:"message" example:"Validation failed"`
	Errors     []string `json:"errors,omitempty"`
}
