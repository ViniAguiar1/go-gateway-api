package domain

import "errors"

var (
	// ErrAccountNotFound is returned when attempting to find an account that doesn't exist
	ErrAccountNotFound = errors.New("account not found")

	// ErrDuplicateAPIKey is returned when trying to create an account with an API key that already exists
	ErrDuplicateAPIKey = errors.New("duplicate API key")

	// ErrInvoiceNotFound is returned when attempting to find an invoice that doesn't exist
	ErrInvoiceNotFound = errors.New("invoice not found")

	// ErrUnauthorizedAccess is returned when a user tries to access a resource they don't have permission for
	ErrUnauthorizedAccess = errors.New("unauthorized access")

	// ErrInvalidAPIKey is returned when the provided API key is malformed or invalid
	ErrInvalidAPIKey = errors.New("invalid API key")

	// ErrInvalidRequest is returned when the request payload is malformed or contains invalid data
	ErrInvalidRequest = errors.New("invalid request")

	// ErrInternalServerError is returned when an unexpected error occurs on the server
	ErrInternalServerError = errors.New("internal server error")

	// ErrInvalidAmount is returned when the amount provided is negative, zero, or exceeds allowed limits
	ErrInvalidAmount = errors.New("invalid amount")

	// ErrInvalidCurrency is returned when the currency code is not supported or invalid
	ErrInvalidCurrency = errors.New("invalid currency")

	// ErrInvalidStatus is returned when the status value is not one of the allowed enum values
	ErrInvalidStatus = errors.New("invalid status")

	// ErrInvalidDate is returned when the date format is invalid or the date is in the past when not allowed
	ErrInvalidDate = errors.New("invalid date")
)
