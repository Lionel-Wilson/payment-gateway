package handlers

import "log"

// Application represents the application with its logging configurations.
type Application struct {
	ErrorLog *log.Logger // Logger for error messages
	InfoLog  *log.Logger // Logger for informational messages
}
