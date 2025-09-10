package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return name, errors.New("mensaje vacio")
	}

	message := fmt.Sprintf(randomFormat(), "")

	return message, nil
}

func Hellos(names []string) (map[string]string, error) {

	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

func randomFormat() string {

	formats := []string{
		"Hola, %v  Bienvenido",
		"Que bueno verte %v",
		"Saludos %v encantado en conocerte"}

	return formats[rand.Intn(len(formats))]
}
