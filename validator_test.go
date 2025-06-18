package validator_test

import (
	"testing"

	validator "github.com/AccelByte/justice-input-validation-go"
	"github.com/stretchr/testify/assert"
)

func Test_Validate(t *testing.T) {
	inputs := map[string]interface{}{
		"alphaNumeric": struct {
			Input string `valid:"alphanum"`
		}{
			Input: "AlphaNumericChars",
		},
		"tag": struct {
			Input string `valid:"tag"`
		}{
			Input: "Tag",
		},
		"language": struct {
			Input string `valid:"language"`
		}{
			Input: "Language",
		},
		"topic": struct {
			Input string `valid:"topic"`
		}{
			Input: "TOPIC",
		},
		"displayName": struct {
			Input string `valid:"displayName"`
		}{
			Input: "DisplayName",
		},
		"personName": struct {
			Input string `valid:"personName"`
		}{
			Input: "PersonName",
		},
		"uuid4WithoutHyphens": struct {
			Input string `valid:"uuid4WithoutHyphens"`
		}{
			Input: "c82fd812449942d2bd74806e32bda386",
		},
		"permissionResource": struct {
			Input string `valid:"permissionResource"`
		}{
			Input: "PERMISSION:RESOURCE",
		},
		"path": struct {
			Input string `valid:"path"`
		}{
			Input: "/a/path/to/somewhere",
		},
		"dateTime": struct {
			Input string `valid:"dateTime"`
		}{
			Input: "2002-10-02T10:00:00-05:00",
		},
		"date": struct {
			Input string `valid:"date"`
		}{
			Input: "1993-11-16",
		},
		"password": struct {
			Input string `valid:"password"`
		}{
			Input: "Password99!@",
		},
		"uri": struct {
			Input string `valid:"uri"`
		}{
			Input: "justice-launcher://",
		},
		"namespace": struct {
			Namespace1 string `valid:"namespace"`
			Namespace2 string `valid:"namespace"`
		}{
			Namespace1: "game1",
			Namespace2: "studio1-game1",
		},
		"country": struct {
			Country string `valid:"country"`
		}{
			Country: "XK",
		},
	}

	for i, v := range inputs {
		t.Run("Test_"+i, func(t *testing.T) {
			valid, err := validator.ValidateStruct(v)
			if err != nil {
				t.Fatal(err.Error())
			}
			assert.True(t, valid, i)
		})
	}
}

func Test_InvalidValidate(t *testing.T) {
	inputs := map[string]interface{}{
		"invalid_country": struct {
			Country string `valid:"country"`
		}{
			Country: "USA",
		},
	}

	for i, v := range inputs {
		t.Run("Test_"+i, func(t *testing.T) {
			valid, _ := validator.ValidateStruct(v)
			assert.False(t, valid, i)
		})
	}
}
