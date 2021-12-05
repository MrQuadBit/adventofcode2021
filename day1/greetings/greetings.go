package greetings

import (
	"errors"
	"fmt"
)

//capital letter means exports (exported name)
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	//var message string
	//message = fmt.Sprintf("Hi, %v. Welcome!", name)
	//:= shortcut for declaring and initializing
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
