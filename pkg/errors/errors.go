package errors

import "errors"

var ErrInvalidFirstName = errors.New("empty or invalid first name")
var ErrInvalidSecondName = errors.New("empty or invalid last name")

var ErrFailedHashing = errors.New("failed to hash")
