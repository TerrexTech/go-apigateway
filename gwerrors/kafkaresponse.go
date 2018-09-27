package gwerrors

import "fmt"

// KRError represents a KafkaResponse error.
type KRError struct {
	Code    int16
	Err     error
	Message string
}

// NewKRError creates a new KRError from provided parameters.
func NewKRError(err error, code int16, msg string) *KRError {
	return &KRError{
		Code:    code,
		Err:     err,
		Message: msg,
	}
}

// Error returns string representation of error.
func (kr *KRError) Error() string {
	return fmt.Sprintf("%d: %s", kr.Code, kr.Message)
}
