package validator_test

import (
	validator "github.com/AccelByte/justice-input-validation-go"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {}

func Test_Validate(t *testing.T) {
	valid, err := validator.Validate(struct {
		String string `valid:"displayName"`
	}{String: "Display Name"})

	if err != nil {
		t.Log(err)
	}
	assert.Nil(t, err)
	assert.True(t, valid)
}
