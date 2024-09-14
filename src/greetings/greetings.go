package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", "", errors.New("empty name")
	}

	// Return a greeting that embeds the name in a message,
	// if a name was given.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	message2 := fmt.Sprintf(randomFormat(), name)
	return message, message2, nil
}

// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages.
	messages := make(map[string]string)
	// Loop through the received slice of names, calling
	// the Hello function to get a message for each name.
	for _, name := range names {
		_, message2, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// In the map, associate the retrieved message with
		// the name.
		messages[name] = message2
	}

	return messages, nil
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}
