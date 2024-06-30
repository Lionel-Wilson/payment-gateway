package handlers

import (
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/Lionel-Wilson/payment-gateway/internal/api/models"
	"github.com/Lionel-Wilson/payment-gateway/internal/api/validators"
	"github.com/Lionel-Wilson/payment-gateway/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Application represents the application with its logging configurations.
type Application struct {
	ErrorLog *log.Logger // Logger for error messages
	InfoLog  *log.Logger // Logger for informational messages
}

var (
	validate *validator.Validate              // Validator for struct validation
	payments map[string]models.PaymentDetails // In-memory store for payment details
	mu       sync.Mutex                       // Mutex to ensure thread-safe access to the payments map
)

func init() {
	validate = validator.New()
	validate.RegisterValidation("expirydate", validators.ExpiryDateValidation)
	payments = make(map[string]models.PaymentDetails)
}

// ProcessPayment handles the processing of a payment.
//
// @Summary      Process a Payment
// @Description  Processes a payment through the payment gateway.
// @Tags         Payments
// @Accept       json
// @Produce      json
// @Param ProccessPaymentRequestBody body ProcessPaymentRequest true "A JSON body" ProccessPaymentRequest()
// @Success      201  {object}  ProcessPaymentResponse
// @Failure      402  {object}  ErrorResponse
// @Failure      422  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /payments [post]
func (app *Application) ProcessPayment(c *gin.Context) {
	var paymentDetails models.ProcessPaymentRequest

	if err := c.ShouldBindJSON(&paymentDetails); err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, "Invalid request payload", nil)
		return
	}

	response, err := createPayment(&paymentDetails)
	if err != nil {
		errMsg := validators.TranslateValidationErrors(err)
		utils.NewErrorResponse(c, http.StatusUnprocessableEntity, "Validation failed", errMsg)
		return
	}

	if response.Status == "payment_paid" {
		c.JSON(http.StatusCreated, response)
	} else if response.Status == "payment_declined" {
		c.JSON(http.StatusPaymentRequired, response)
	} else {
		utils.NewErrorResponse(c, http.StatusInternalServerError, "Something went wrong. Please try again later.", nil)
	}
}

// RetrievePayment retrieves the details of a previously made payment using its identifier.
//
// @Summary      Retrieve Payment Details
// @Description  Retrieves the details of a previously made payment using its identifier.
// @Tags         Payments
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Payment ID"
// @Success      200  {object}  PaymentDetails
// @Failure      404  {object}  ErrorResponse
// @Router       /payments/{id} [get]
func (app *Application) RetrievePayment(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		utils.NewErrorResponse(c, http.StatusBadRequest, "Invalid id provided", nil)
		return
	}

	id = strings.TrimSpace(id)

	mu.Lock()
	payment, exists := payments[id]
	mu.Unlock()

	if !exists {
		utils.NewErrorResponse(c, http.StatusNotFound, "Payment not found", nil)
		return
	}

	c.JSON(http.StatusOK, payment)
}

// AllPayments retrieves the details of all previously made payments.
//
// @Summary      Retrieve all payments
// @Description  Retrieves the details of all previously made payments.
// @Tags         Payments
// @Accept       json
// @Produce      json
// @Success      200  {array}  PaymentDetails
// @Failure      404  {object}  ErrorResponse
// @Router       /payments [get]
func (app *Application) AllPayments(c *gin.Context) {
	mu.Lock()
	defer mu.Unlock()

	if len(payments) == 0 {
		utils.NewErrorResponse(c, http.StatusNotFound, "No payments available", nil)
		return
	}

	paymentsList := []models.PaymentDetails{}
	for _, payment := range payments {
		paymentsList = append(paymentsList, payment)
	}

	c.JSON(http.StatusOK, paymentsList)
}

func createPayment(paymentDetails *models.ProcessPaymentRequest) (models.ProcessPaymentResponse, error) {
	// Trim whitespace from payment details
	utils.TrimWhitespace(paymentDetails)

	// Validate payment details
	err := validate.Struct(paymentDetails)
	if err != nil {
		return models.ProcessPaymentResponse{}, err
	}

	// Simulate bank processing
	id, status, statusCode, summary := utils.SimulateBank()

	// Lock payments map before updating
	mu.Lock()
	defer mu.Unlock()

	// Update payments map with new payment details
	payments[id] = models.PaymentDetails{
		ID:           id,
		FirstName:    paymentDetails.FirstName,
		LastName:     paymentDetails.LastName,
		CardNumber:   utils.MaskCardNumber(paymentDetails.CardNumber),
		ExpiryDate:   paymentDetails.ExpiryDate,
		Amount:       paymentDetails.Amount,
		CurrencyCode: paymentDetails.CurrencyCode,
		Status:       status,
		StatusCode:   statusCode,
	}

	// Prepare response
	response := models.ProcessPaymentResponse{
		ID:              id,
		Status:          status,
		ResponseSummary: summary,
	}

	return response, nil
}
