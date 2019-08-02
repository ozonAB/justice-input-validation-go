package validator

import (
	"regexp"
	"unicode/utf8"
)

func IsDisplayName(str string) bool {
	valid, err := regexp.Match("^[a-zA-Z0-9]+(([a-zA-Z0-9 ])?[a-zA-Z0-9]*)*$", []byte(str))
	if !valid || err != nil {
		return false
	}
	if utf8.RuneCountInString(str) > 256 {
		return false
	}

	return true
}

func IsAlphaNoWhiteSpace(str string) bool {
	valid, err := regexp.Match("^[a-zA-Z0-9]+([_:-]{1}[a-zA-Z0-9]+)*$", []byte(str))
	if !valid || err != nil {
		return false
	}
	if utf8.RuneCountInString(str) > 30 {
		return false
	}
	return true
}

func IsAlphaNoWhiteSpaceWithDash(str string) bool {
	valid, err := regexp.Match("^[a-zA-Z]+([-]{1}[a-zA-Z]+)*$", []byte(str))
	if !valid || err != nil {
		return false
	}
	if utf8.RuneCountInString(str) > 256 {
		return false
	}
	return true
}
