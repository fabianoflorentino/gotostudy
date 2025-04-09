package errormsg

import (
	"errors"
)

var (
	ErrEnvNotSet = errors.New("environment variable not set: ")
)
