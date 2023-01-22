package util

import "fmt"

type _Error_ struct {
	ErrCode int
	Message interface{}
}

func (e *_Error_) Error() string {
	return fmt.Sprintf("ErrCode: %d, Message: %s", e.ErrCode, e.Message)
}

func NewError(errCode int, message interface{}) error {
	return &_Error_{
		ErrCode: errCode,
		Message: message,
	}
}

func GetErrorValue(err error) _Error_ {
	if e, ok := err.(*_Error_); ok {
		return *e
	}
	return _Error_{}
}
