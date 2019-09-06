package validator_test

import (
	"fmt"
	"testing"

	validator "github.com/AccelByte/justice-input-validation-go"
	"github.com/stretchr/testify/assert"
)

func Test_IsAlphaNumeric(t *testing.T) {
	t.Run("Test_AlphaNumericValid", func(t *testing.T) {
		input := "validString"
		valid := validator.BaseValidator("alphanum", input)
		assert.True(t, valid)
	})
	t.Run("Test_AlphaNumericValidWithoutParams", func(t *testing.T) {
		input := "validString"
		valid := validator.BaseValidator("alphanum", input)
		assert.True(t, valid)
	})
	t.Run("Test_AlphaNumericInValid", func(t *testing.T) {
		inputs := []string{"invalid string", "inv@lidString"}
		for _, input := range inputs {
			valid := validator.BaseValidator("alphanum", input)
			assert.False(t, valid)
		}
	})
}

func Test_IsTag(t *testing.T) {
	t.Run("Test_TagValid", func(t *testing.T) {
		inputs := []string{"valid", "Valid", "valiD", "vAlid"}
		for _, input := range inputs {
			valid := validator.IsTag(input)
			assert.True(t, valid)
		}
	})
	t.Run("Test_TagInValidChars", func(t *testing.T) {
		inputs := []string{"1nv@lid", "inv a lid"}
		for _, input := range inputs {
			valid := validator.IsTag(input)
			assert.False(t, valid)
		}
	})
}

func Test_IsLanguage(t *testing.T) {
	t.Run("Test_LanguageValid", func(t *testing.T) {
		inputs := []string{"valid", "Valid", "valiD", "vAlid", "va-lid"}
		for _, input := range inputs {
			valid := validator.IsLanguage(input)
			assert.True(t, valid)
		}
	})
	t.Run("Test_LanguageInValidChars", func(t *testing.T) {
		inputs := []string{"1nv@lid", "inv a lid"}
		for _, input := range inputs {
			valid := validator.IsLanguage(input)
			assert.False(t, valid)
		}
	})
}

func Test_IsTopic(t *testing.T) {
	t.Run("Test_TopicValid", func(t *testing.T) {
		inputs := []string{"VALID", "VA_LID"}
		for _, input := range inputs {
			valid := validator.IsTopic(input)
			assert.True(t, valid)
		}
	})
	t.Run("Test_TopicInValidChars", func(t *testing.T) {
		inputs := []string{"invalid", "INVALId", "iNVALID", "INVALID_", "_INVALID", "INV@LID"}
		for _, input := range inputs {
			valid := validator.IsTopic(input)
			assert.False(t, valid)
		}
	})
}

func Test_IsDisplayName(t *testing.T) {
	t.Run("Test_DisplayNameValid", func(t *testing.T) {
		inputs := []string{"Valid Name", "Valid Name 12345", "VALID NAME", "DisplayName"}
		for _, input := range inputs {
			valid := validator.IsDisplayName(input)
			assert.True(t, valid)
		}
	})
	t.Run("Test_DisplayNameInvalidChars", func(t *testing.T) {
		inputs := []string{"In+Valid", "!nValid", "Inv@lid"}
		for _, input := range inputs {
			valid := validator.IsDisplayName(input)
			assert.False(t, valid)
		}
	})
}

func Test_IsPersonName(t *testing.T) {
	t.Run("Test_IsPersonNameValid", func(t *testing.T) {
		inputs := []string{"Valid Name", "Valid.Name", "valid,name", "valid-name", "VALID NAME"}
		for _, input := range inputs {
			valid := validator.IsPersonName(input)
			assert.True(t, valid)
		}
	})
	t.Run("Test_IsPersonNameInValidChars", func(t *testing.T) {
		inputs := []string{"inv4lid", "in_valid", "Inv@lid"}
		for _, input := range inputs {
			valid := validator.IsPersonName(input)
			assert.False(t, valid)
		}
	})
}

func Test_UUID4WithoutHyphens(t *testing.T) {
	t.Run("Test_UUID4WithoutHyphensValid", func(t *testing.T) {
		inputs := []string{"c82fd812449942d2bd74806e32bda386"}
		for _, input := range inputs {
			valid := validator.IsUUID4WithoutHyphens(input)
			assert.True(t, valid)
		}
	})
	t.Run("Test_UUID4WithoutHyphensInValidChars", func(t *testing.T) {
		inputs := []string{"inv4lid", "in_valid", "Inv@lid"}
		for _, input := range inputs {
			valid := validator.IsUUID4WithoutHyphens(input)
			assert.False(t, valid)
		}
	})
}

func Test_IsEmail(t *testing.T) {
	t.Run("Test_IsEmailValid", func(t *testing.T) {
		inputs := []string{"email@mail.com", "email_new@mail.com", "email-new@mail.com", "email.new@mail.com"}
		for i, input := range inputs {
			valid := validator.IsEmail(input)
			assert.True(t, valid, i)
		}
	})
	t.Run("Test_IsEmailInvalidChars", func(t *testing.T) {
		inputs := []string{"em@il@mail.com", "email @gmail.com"}
		for _, input := range inputs {
			valid := validator.IsEmail(input)
			assert.False(t, valid)
		}
	})
}

func Test_IsPermissionResource(t *testing.T) {
	t.Run("Test_IsPermissionResourceValid", func(t *testing.T) {
		inputs := []string{"PERMISSION", "PERMISSION:RESOURCE"}
		for i, input := range inputs {
			valid := validator.IsPermissionResource(input)
			assert.True(t, valid, i)
		}
	})
	t.Run("Test_IsPermissionResourceInvalidChars", func(t *testing.T) {
		inputs := []string{"permission", "PERMISSION RESOURCE"}
		for _, input := range inputs {
			valid := validator.IsPermissionResource(input)
			assert.False(t, valid)
		}
	})
}

func Test_IsPath(t *testing.T) {
	t.Run("Test_IsPathValid", func(t *testing.T) {
		inputs := []string{"/path", "/path/TO/SomeWhere"}
		for i, input := range inputs {
			valid := validator.IsPath(input)
			assert.True(t, valid, i)
		}
	})
	t.Run("Test_IsPathInvalidChars", func(t *testing.T) {
		inputs := []string{"path ", "path to", "/path wow/nice"}
		for _, input := range inputs {
			valid := validator.IsPath(input)
			assert.False(t, valid)
		}
	})
}

func Test_IsPassword(t *testing.T) {
	t.Run("Test_PasswordValid", func(t *testing.T) {
		inputs := []string{"Password1", "p@ssword1", "P@SSWORD1", "Password!"}
		for i, input := range inputs {
			valid := validator.IsPassword(input)
			assert.True(t, valid, i)
		}
	})
	t.Run("Test_PasswordInvalidRules", func(t *testing.T) {
		inputs := []string{
			"password",  // only match 1 rule: lowercase
			"PASSWORD",  // only match 1 rule: uppercase
			"12345678",  // only match 1 rule: number
			"!@#$%^&*",  // only match 1 rule: special chars
			"Password",  // only match 2 rules: uppercase & lowercase
			"password1", // only match 2 rules: lowercase & number
			"PASSWORD1", // only match 2 rules: uppercase & number
			"@#$%^&*1",  // only match 2 rules: special chars & number
			"password!", // only match 2 rules: special chars & lowercase
			"PASSWORD!", // only match 2 rules: special chars & uppercase
		}
		for i, input := range inputs {
			valid := validator.IsPassword(input)
			assert.False(t, valid, i)
		}
	})
}

func Test_BaseValidator(t *testing.T) {
	type tempResult struct {
		inputValue   string
		assertResult bool
	}
	t.Run("Test_AvailableBaseValidator", func(t *testing.T) {
		inputTest := make(map[string]tempResult, 3)
		inputTest["ISO3166Alpha2"] = tempResult{
			inputValue:   "US",
			assertResult: true,
		}
		inputTest["ipv4"] = tempResult{
			inputValue:   "192.168.1.1",
			assertResult: true,
		}
		inputTest["numeric"] = tempResult{
			inputValue:   "inputTest",
			assertResult: false,
		}

		for rule, temp := range inputTest {
			valid := validator.BaseValidator(rule, temp.inputValue)
			assert.Equal(t, temp.assertResult, valid, fmt.Sprintf("%s with input %s should be %t, %t found", rule, temp.inputValue, temp.assertResult, valid))
		}
	})
	t.Run("Test_UnavailableBaseValidator", func(t *testing.T) {
		valid := validator.BaseValidator("ruleNotFound", "string")
		assert.False(t, valid)
	})
}

func Test_IsContainWhitespace(t *testing.T) {
	t.Run("Test_IsContainWhitespaceValid", func(t *testing.T) {
		inputs := []string{"hello world", "he llo"}
		for i, input := range inputs {
			valid := validator.IsContainWhitespace(input)
			assert.True(t, valid, i)
		}
	})
	t.Run("Test_IsContainWhitespaceInvalid", func(t *testing.T) {
		inputs := []string{"helloWorld", "hello", "!@#$%^&*"}
		for _, input := range inputs {
			valid := validator.IsContainWhitespace(input)
			assert.False(t, valid)
		}
	})
}

func Test_IsNotContainWhitespace(t *testing.T) {
	t.Run("Test_IsContainWhitespaceInvalid", func(t *testing.T) {
		inputs := []string{"hello world", "he llo"}
		for i, input := range inputs {
			valid := validator.IsNotContainWhitespace(input)
			assert.False(t, valid, i)
		}
	})
	t.Run("Test_IsContainWhitespaceValid", func(t *testing.T) {
		inputs := []string{"helloWorld", "hello", "!@#$%^&*"}
		for _, input := range inputs {
			valid := validator.IsNotContainWhitespace(input)
			assert.True(t, valid)
		}
	})
}

func Test_IsCodeChallenge(t *testing.T) {
	t.Run("Test_IsCodeChallengeValid", func(t *testing.T) {
		inputs := []string{"47DEQpj8HBSa-_TImW-5JCeuQeRkm5NMpJWZG3hSuFU", "a7_sLldsHuLmHKQzpS55XjADbVowVDS1GagArlnVpT8",
			"1i4jrkEui1pXETxcmXRliqbDhk5voCdpfOxeOHjUBH4", "QYtYTirpqMsy-M4RyOZdT9mVLlFvdSjo-dxyBbKpUME"}
		for i, input := range inputs {
			valid := validator.IsCodeChallenge(input)
			assert.True(t, valid, i)
		}
	})
	t.Run("Test_IsCodeChallengeInvalidChars", func(t *testing.T) {
		inputs := []string{"@#!$#%#%^@@%^^^&&%#"}
		for _, input := range inputs {
			valid := validator.IsCodeChallenge(input)
			assert.False(t, valid)
		}
	})
}

func Test_IsDate(t *testing.T) {
	t.Run("Test_IsDateValid", func(t *testing.T) {
		valid := validator.IsDate("2019-09-03")
		assert.True(t, valid)
	})
	t.Run("Test_IsDateInvalid", func(t *testing.T) {
		inputs := []string{"03-09-19", "09-2019-03"}
		for _, input := range inputs {
			valid := validator.IsDate(input)
			assert.False(t, valid)
		}
	})
}
