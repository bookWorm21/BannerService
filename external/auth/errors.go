package auth

import "errors"

var (
	Unauthorized = errors.New("Unauthorized")
	NoAccess     = errors.New("NoAccess")
)
