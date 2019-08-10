package puppy

import "fmt"

type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("Error code %d: %s",
		e.Code, e.Message)
}

const (
	ErrorValueFormat = 1000
	NegativeValue    = 1001
	NonExistentPuppy = 1002
)
