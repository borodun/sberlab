package messages

type ErrorMsg struct {
	Error ErrorWithID `json:"error"`
}

type ErrorWithID struct {
	Message string `json:"message"`
}
