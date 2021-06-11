package errors

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

type Error struct {
	Code    int
	Message string
	Cause   *Error
	Line    int
	File    string
}

func (e Error) GetStack(builder *strings.Builder) {
	builder.WriteString(e.File + ":" + strconv.Itoa(e.Line) + " " + e.Message + "\n")
	if e.Cause != nil {
		e.Cause.GetStack(builder)
	}
}

func (e Error) Error() string {

	stackSb := strings.Builder{}
	e.GetStack(&stackSb)

	return stackSb.String()
}
func CausedError(cause error, format string, args ...interface{}) Error {
	msg := fmt.Sprintf(format, args...)

	var ourError *Error = nil
	if cause != nil {
		errCast, ok := cause.(Error)
		if !ok {
			ourError = &Error{
				Message: cause.Error(),
				Cause:   nil,
				Line:    0,
				File:    "",
			}
		} else {
			ourError = &errCast
		}
	}

	_, file, line, _ := runtime.Caller(2)

	return Error{
		Message: msg,
		Cause:   ourError,
		File:    file,
		Line:    line,
	}

}
func BasicError(format string, args ...interface{}) Error {
	return CausedError(nil, format, args...)
}

func CausedErrorTrue(errP *error, cause error, format string, args ...interface{}) bool {
	if cause != nil {
		*errP = CausedError(cause, format, args...)
		return true
	}
	return false
}
