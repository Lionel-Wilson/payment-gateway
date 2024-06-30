package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	_ "github.com/Lionel-Wilson/payment-gateway/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/Lionel-Wilson/payment-gateway/internal/api/handlers"
	"github.com/Lionel-Wilson/payment-gateway/internal/api/middlewares"
)

func main() {
	addr := os.Getenv("DEV_ADDRESS")

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &handlers.Application{
		ErrorLog: errorLog,
		InfoLog:  infoLog,
	}

	// Set up Gin router
	r := gin.Default()

	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	r.Use(middlewares.SecureHeaders())
	r.Use(middlewares.CorsMiddleware())

	apiV1 := r.Group("/api/v1")
	{
		apiV1.POST("/payments", app.ProcessPayment)
		apiV1.GET("/payments/:id", app.RetrievePayment)
		apiV1.GET("/payments", app.AllPayments)

		apiV1.GET("/health", app.HealthCheck)
	}

	// Serve Swagger documentation
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	infoLog.Printf("Starting server on %s", addr)
	err := r.Run(addr)
	if err != nil {
		errorLog.Fatal(err)
	}
}
