package constants

import "errors"

var (
	BearerTokenHasError = errors.New("Bearer token catch error")
	BearerTokenInvalid  = errors.New("Invalid token")

	UserNotFound = errors.New("User not found")
	FailedLogin  = errors.New("Sorry, the email or password you entered is incorrect.")
)
