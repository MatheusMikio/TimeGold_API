package models

type ErrorMessage struct {
	Property string
	Message  string
}

func CreateErrorMessage(property, message string) *ErrorMessage {
	return &ErrorMessage{
		Property: property,
		Message:  message,
	}
}
