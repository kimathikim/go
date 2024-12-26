package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	// create a message using a random format

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}
  // Hello returns a greeting for the named person.
func Hellos(names []string) (map[string]string, error) {
  // a map to associate names with messages
  messages := make(map[string]string)
  // loop thought the provided slice names with the message
for _, name := range names {
    message, err := Hello(name)
    if err != nil {
      return nil, err
    }

  // 
	}

	// create a message using a random format

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func randomFormat() string {
	// a slice of message format
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Bonjour, %v!",
	}
	// return a randomly selected message format by specifying a random index
	return formats[rand.Intn(len(formats))]
}
