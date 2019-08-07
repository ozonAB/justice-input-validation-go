package validator_test

import (
	validator "github.com/AccelByte/justice-input-validation-go"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Validate(t *testing.T) {
	inputs := map[string]interface{}{
		"alphaNumeric": struct {
			Input string `valid:"alphaNumeric(256)"`
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
		"userDisplayName": struct {
			Input string `valid:"userDisplayName"`
		}{
			Input: "DisplayName",
		},
		"uuid4WithoutHyphens": struct {
			Input string `valid:"uuid4WithoutHyphens"`
		}{
			Input: "c82fd812449942d2bd74806e32bda386",
		},
		"orderNumber": struct {
			Input string `valid:"orderNumber"`
		}{
			Input: "O1234567890123456",
		},
		"password": struct {
			Input string `valid:"password"`
		}{
			Input: "p",
		},
		"dockerImage": struct {
			Input string `valid:"dockerImage"`
		}{
			Input: "alpine:latest",
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
		"memorySize": struct {
			Input string `valid:"memorySize"`
		}{
			Input: "1Mi",
		},
		"time": struct {
			Input string `valid:"time"`
		}{
			Input: "2002-10-02T10:00:00-05:00",
		},
		"dob": struct {
			Input string `valid:"dob"`
		}{
			Input: "1993-11-16",
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
