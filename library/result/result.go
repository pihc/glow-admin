package result

import "fmt"

type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func Response(data interface{}, err error) Result {
	if err != nil {
		return Error(err)
	}
	return Success(data, "")
}
func Success(data interface{}, message string) Result {
	return Result{
		Code:    0,
		Data:    data,
		Message: message,
	}
}
func Errorf(format string, a ...interface{}) Result {
	return Result{
		Code:    500,
		Message: fmt.Sprintf(format, a...),
	}
}
func Error(err error) Result {
	return Result{
		Code:    500,
		Message: err.Error(),
	}
}

func ErrorCode(code int, message string) Result {
	return Result{
		Code:    code,
		Message: message,
	}
}
