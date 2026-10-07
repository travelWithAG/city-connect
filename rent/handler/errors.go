package handler

import (
	"errors"
	"fmt"
)

var (
	NotFound = func(a string, b string) error {
		return fmt.Errorf("%s: %s", a, b)
	}
	ExceptionError = func () error  {
		return errors.New("Faced exception")
	}
)