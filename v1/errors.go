package client

import "encoding/json"

type SimpleError struct {
	Message string `json:"error"`
}

func (se SimpleError) Error() string {
	return se.Message
}

type ValidationErrors struct {
	Errors map[string][]string `json:"errors"`
}

func (ve ValidationErrors) Error() string {
	b, _ := json.Marshal(ve)
	return string(b)
}
