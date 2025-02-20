package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// - Hello devuelve un saludo para la persona especificada

func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("Nombre Vacio")
	}
	// - Devuelve un saludo que incluye el nobmre en un mensaje
	message := fmt.Sprintf(randomFormat(), name)
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

// - randomFormat devuelvo uno de varios formatos de mensajes
func randomFormat() string {
	formats := []string{
		"Buenos dias %v!!!!",
		"Boa tarde %v!!!",
		"Good night %v!!!",
	}
	return formats[rand.Intn(len(formats))]
}
