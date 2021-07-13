package error_check

import (
	"backend/api/v1/messages"
)

type ErrorCheck struct {
	ErrorMessage string               `json:"error_msg"`
	Error        messages.ErrorWithID `json:"error"`
}

func CheckError(check ErrorCheck) string {
	if len(check.ErrorMessage) != 0 {
		return check.ErrorMessage
	} else if len(check.Error.Message) != 0 {
		return check.Error.Message
	}
	return ""
}
