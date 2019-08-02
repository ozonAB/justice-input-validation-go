package validator_test

import (
	validator "github.com/AccelByte/justice-input-validation-go"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_DisplayName(t *testing.T) {
	t.Run("Test_DisplayNameValid", func(t *testing.T) {
		displayName := "ValidDisplayName"
		valid := validator.IsDisplayName(displayName)

		assert.True(t, valid)
	})
	t.Run("Test_DisplayNameInValidSymbol", func(t *testing.T) {
		displayName := "InvalidDisplayName @#$%^&*()"
		valid := validator.IsDisplayName(displayName)

		assert.False(t, valid)
	})
	t.Run("Test_DisplayNameInValidLength", func(t *testing.T) {
		displayName := `
			Lorem ipsum dolor sit amet, consectetur adipiscing elit. 
			Ut at elit eget risus maximus luctus eu ac dui. 
			Praesent convallis consequat enim eget gravida. 
			Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas.
		`
		valid := validator.IsDisplayName(displayName)
		assert.False(t, valid)
	})
}

func TestIsAlphaNoWhiteSpace(t *testing.T) {
	t.Run("Test_AlphaNoWhiteSpaceValid", func(t *testing.T) {
		input := "validString"
		valid := validator.IsAlphaNoWhiteSpace(input)
		assert.True(t, valid)
	})
	t.Run("Test_AlphaNoWhiteSpaceInValidWhiteSpaceExists", func(t *testing.T) {
		input := "invalid string"
		valid := validator.IsAlphaNoWhiteSpace(input)
		assert.False(t, valid)
	})
	t.Run("Test_AlphaNoWhiteSpaceInValidLength", func(t *testing.T) {
		input := "invalidLengthTooLongStringMustBeInValid"
		valid := validator.IsAlphaNoWhiteSpace(input)
		assert.False(t, valid)
	})
}

func TestIsAlphaNoWhiteSpaceWithDash(t *testing.T) {
	t.Run("Test_AlphaNoWhiteSpaceWithDashValid", func(t *testing.T) {
		valid := validator.IsAlphaNoWhiteSpaceWithDash("valid-string")
		assert.True(t, valid)
	})
}
