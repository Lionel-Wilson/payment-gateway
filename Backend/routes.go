package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	// using bmizernay/pat package to implement RESTful routes
	router := pat.New()

	router.Post("/payments", http.HandlerFunc(app.proccessPayment))
	router.Get("/payments", http.HandlerFunc(app.retrievePaymentDetails))

	return standardMiddleware.Then(router)
}
