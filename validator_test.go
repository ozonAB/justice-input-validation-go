package validator_test

import (
	validator "github.com/AccelByte/justice-input-validation-go"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Validate(t *testing.T) {
	inputs := map[string]interface{}{
		"displayName": struct {
			Input string `valid:"displayName"`
		}{Input: "Display Name"},
		"alphaNoWhiteSpace": struct {
			Input string `valid:"alphaNoWhiteSpace"`
		}{Input: "alphaNoWhiteSpace"},
		"alphaNoWhiteSpaceWithDash": struct {
			Input string `valid:"alphaNoWhiteSpaceWithDash"`
		}{Input: "alphaNoWhiteSpaceWithDash"},
	}

	for k, v := range inputs {
		t.Run("Test_"+k, func(t *testing.T) {
			valid, err := validator.Validate(v)
			assert.Nil(t, err)
			assert.True(t, valid)
		})
	}
}
