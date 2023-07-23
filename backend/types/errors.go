package types

import "errors"

const ServerErrorMessage = "Something's wrong. Please contact admin."
const BadRequestErrorMessage = "Bad Request"

var (
	ErrNotFound  = errors.New("not found")
	ErrDuplicate = errors.New("duplicate entry")
	ErrInvalid   = errors.New("invalid")
)
