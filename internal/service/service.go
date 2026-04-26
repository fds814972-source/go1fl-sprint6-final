package service

import (
	"fmt"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func AutoConvertation(data string) (string, error) {
	td := strings.TrimSpace(data)
	if len(td) == 0 {
		return "", fmt.Errorf("Empty line")
	}

	isMorse := true
	// функция определения строки или морзе с конвертацией
	for _, r := range td {
		if r != '.' && r != '-' && r != ' ' {
			isMorse = false
			break
		}
	}

	if isMorse {
		td = morse.ToText(td)
	} else {
		td = morse.ToMorse(td)
	}

	return td, nil
}
