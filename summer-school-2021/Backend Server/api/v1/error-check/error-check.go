package error_check

import (
	"encoding/json"
)

type ErrorWithID struct {
	Message string `json:"message"`
}

type ErrorCheck struct {
	ErrorMessage string      `json:"error_msg"`
	Error        ErrorWithID `json:"error"`
	Message      string      `json:"message"`
}

func CheckError(check string) string {
	var errCheck ErrorCheck
	if err := json.Unmarshal([]byte(check), &errCheck); err != nil {
		return err.Error()
	}

	if len(errCheck.ErrorMessage) != 0 {
		return errCheck.ErrorMessage
	} else if len(errCheck.Error.Message) != 0 {
		return errCheck.Error.Message
	} else if len(errCheck.Message) != 0 {
		return errCheck.Message
	}
	return ""
}
