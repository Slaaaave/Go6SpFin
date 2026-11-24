package service

import (
	"errors"
	"strings"
	"unicode"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func IsMorse(input string) (string, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return "", errors.New("пустой файл")
	}

	if justText(input) {
		return morse.ToMorse(input), nil
	}

	result := morse.ToText(input)

	if result == "" && input != "" {
		return "", errors.New("Ошибка ")
	}

	return result, nil
}

func justText(input string) bool {
	for _, r := range input {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			return true
		}
	}
	return false
}
