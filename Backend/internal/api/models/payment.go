package models

// ProcessPaymentRequest represents a request to process a payment.
// It includes details like the cardholder's name, card number, expiry date, amount, currency, and CVV.
type ProcessPaymentRequest struct {
	FirstName    string  `json:"firstName" example:"John" validate:"required,alpha"`                    // The first name of the cardholder. Required and must be alphabetic.
	LastName     string  `json:"lastName" example:"Doe" validate:"required,alpha"`                      // The last name of the cardholder. Required and must be alphabetic.
	CardNumber   string  `json:"cardNumber" example:"4111111111111111" validate:"required,credit_card"` // The credit card number. Required and must be a valid credit card number.
	ExpiryDate   string  `json:"expiryDate" example:"12/29" validate:"required,expirydate"`             // The expiry date of the credit card in MM/YY format. Required with custom validation.
	Amount       float64 `json:"amount" example:"500" validate:"required,gt=0"`                         // The amount to be charged. Required and must be greater than 0.
	CurrencyCode string  `json:"currencyCode" example:"GBP" validate:"required,len=3,alpha"`            // The currency code for the transaction. Required, must be 3 alphabetic characters.
	CVV          string  `json:"cvv" example:"123" validate:"required,len=3,numeric"`                   // The CVV of the credit card. Required, must be exactly 3 numeric characters.
}

// ProcessPaymentResponse represents a response after processing a payment.
// It includes an ID, status, and a response summary.
type ProcessPaymentResponse struct {
	ID              string `json:"id" example:"PAY-1625843728243722000"` // The unique identifier for the payment transaction.
	Status          string `json:"status" example:"payment_paid"`        // The status of the payment transaction.
	ResponseSummary string `json:"responseSummary" example:"Approved"`   // A summary of the payment response.
}

// PaymentDetails represents the details of a processed payment.
// It includes the payment ID, cardholder's name, masked card number, expiry date, amount, currency, status, and status code.
type PaymentDetails struct {
	ID           string  `json:"id" example:"PAY-1625843728243722000"`  // The unique identifier for the payment transaction.
	FirstName    string  `json:"firstName" example:"John"`              // The first name of the cardholder.
	LastName     string  `json:"lastName" example:"Doe"`                // The last name of the cardholder.
	CardNumber   string  `json:"cardNumber" example:"************1111"` // The masked credit card number.
	ExpiryDate   string  `json:"expiryDate" example:"12/29"`            // The expiry date of the credit card in MM/YY format.
	Amount       float64 `json:"amount" example:"500"`                  // The amount charged in the transaction.
	CurrencyCode string  `json:"currencyCode" example:"GBP"`            // The currency code for the transaction.
	Status       string  `json:"status" example:"payment_paid"`         // The status of the payment transaction.
	StatusCode   int     `json:"statusCode" example:"10000"`            // The status code of the payment transaction.
}
