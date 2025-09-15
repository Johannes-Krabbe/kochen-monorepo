package errors

type ErrorType string

const (
	InternalServerError ErrorType = "internal_server_error"
	Validation ErrorType = "validation"

	InvalidCredentials ErrorType = "invalid_credentials"
	UsernameAlreadyTaken ErrorType = "username_already_taken"
	EmailAlreadyTaken ErrorType = "username_already_taken"
)
