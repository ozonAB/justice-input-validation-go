package validator

import (
	"regexp"
	"unicode/utf8"
)

func IsDisplayName(str string) bool {
	if utf8.RuneCountInString(str) > 256 {
		return false
	}

	matched, err := regexp.Match("^[a-zA-Z0-9]+(([a-zA-Z0-9 ])?[a-zA-Z0-9]*)*$", []byte(str))
	if !matched || err != nil {
		return false
	}

	return true
}

func IsAlphaNoWhiteSpace(str string) bool {
	return true
}
