package validator_test

import (
	"testing"

	validator "github.com/AccelByte/justice-input-validation-go"
	"github.com/stretchr/testify/assert"
)

var invalidLengthInput256 = `
	Lorem ipsum dolor sit amet, consectetur adipiscing elit.
	Nunc vitae justo porta, dapibus nisl et, tempus ex. 
	Etiam dolor sem, blandit nec ultricies ultrices, aliquam eu tortor. 
	Maecenas bibendum diam elit. Cras malesuada et nibh et viverra.
`

func Test_IsAlphaNumeric(t *testing.T) {
	t.Run("Test_AlphaNumericValid", func(t *testing.T) {
		input := "validString"
		valid := validator.IsAlphaNumeric(input, "256")
		assert.True(t, valid)
	})
	t.Run("Test_AlphaNumericValidWithoutParams", func(t *testing.T) {
		input := "validString"
		valid := validator.IsAlphaNumeric(input)
		assert.True(t, valid)
	})
	t.Run("Test_AlphaNumericInValid", func(t *testing.T) {
		inputs := []string{"invalid string", "inv@lidString"}
		for _, input := range inputs {
			valid := validator.IsAlphaNumeric(input)
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
	t.Run("Test_TagInValidLength", func(t *testing.T) {
		input := "moreThan30Charsaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
		valid := validator.IsTag(input)
		assert.False(t, valid)
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
	t.Run("Test_LanguageInValidLength", func(t *testing.T) {
		input := invalidLengthInput256
		valid := validator.IsLanguage(input)
		assert.False(t, valid)
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
	t.Run("Test_TopicInValidLength", func(t *testing.T) {
		input := "moreThan5Chars"
		valid := validator.IsTopic(input)
		assert.False(t, valid)
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
	t.Run("Test_DisplayNameInValidLength", func(t *testing.T) {
		input := invalidLengthInput256
		valid := validator.IsDisplayName(input)
		assert.False(t, valid)
	})
	t.Run("Test_DisplayNameInValidChars", func(t *testing.T) {
		inputs := []string{"In_Valid", "In-Valid", "Inv@lid"}
		for _, input := range inputs {
			valid := validator.IsDisplayName(input)
			assert.False(t, valid)
		}
	})
}

func Test_IsUserDisplayName(t *testing.T) {
	t.Run("Test_UserDisplayNameValid", func(t *testing.T) {
		inputs := []string{"Valid Name", "Valid.Name", "valid,name", "valid-name", "VALID NAME"}
		for _, input := range inputs {
			valid := validator.IsUserDisplayName(input)
			assert.True(t, valid)
		}
	})
	t.Run("Test_UserDisplayNameInValidLength", func(t *testing.T) {
		input := invalidLengthInput256
		valid := validator.IsUserDisplayName(input)
		assert.False(t, valid)
	})
	t.Run("Test_UserDisplayNameInValidChars", func(t *testing.T) {
		inputs := []string{"inv4lid", "in_valid", "Inv@lid"}
		for _, input := range inputs {
			valid := validator.IsUserDisplayName(input)
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
	t.Run("Test_UUID4WithoutHyphensInValidLength", func(t *testing.T) {
		input := "c82fd812449942d2bd74806e32bda3861"
		valid := validator.IsUUID4WithoutHyphens(input)
		assert.False(t, valid)
	})
	t.Run("Test_UUID4WithoutHyphensInValidChars", func(t *testing.T) {
		inputs := []string{"inv4lid", "in_valid", "Inv@lid"}
		for _, input := range inputs {
			valid := validator.IsUUID4WithoutHyphens(input)
			assert.False(t, valid)
		}
	})
}

func Test_OrderNumber(t *testing.T) {
	t.Run("Test_OrderNumberValid", func(t *testing.T) {
		inputs := []string{"O1234567890123456"}
		for _, input := range inputs {
			valid := validator.IsOrderNumber(input)
			assert.True(t, valid)
		}
	})
	t.Run("Test_OrderNumberInvalidLength", func(t *testing.T) {
		input := "012345678912345671"
		valid := validator.IsOrderNumber(input)
		assert.False(t, valid)
	})
	t.Run("Test_OrderNumberInvalidChars", func(t *testing.T) {
		inputs := []string{"01234567891234567", "O123456789123456a", "O1234abcd6789123"}
		for _, input := range inputs {
			valid := validator.IsOrderNumber(input)
			assert.False(t, valid)
		}
	})
}

func Test_IsDockerImage(t *testing.T) {
	t.Run("Test_IsDockerImageValid", func(t *testing.T) {
		inputs := []string{"alpine", "alpine:latest", "alpine-new:latest", "alpine:123"}
		for _, input := range inputs {
			valid := validator.IsDockerImage(input)
			assert.True(t, valid)
		}
	})
	t.Run("Test_IsDockerImageInvalidLength", func(t *testing.T) {
		input := invalidLengthInput256
		valid := validator.IsDockerImage(input)
		assert.False(t, valid)
	})
	t.Run("Test_IsDockerImageInvalidChars", func(t *testing.T) {
		inputs := []string{"Alpine", "Alpine latest", "@lpine"}
		for _, input := range inputs {
			valid := validator.IsDockerImage(input)
			assert.False(t, valid)
		}
	})
}

func Test_IsOWASPEmail(t *testing.T) {
	t.Run("Test_IsOWASPEmailValid", func(t *testing.T) {
		inputs := []string{"email@mail.com", "email_new@mail.com", "email-new@mail.com", "email.new@mail.com"}
		for i, input := range inputs {
			valid := validator.IsOWASPEmail(input)
			assert.True(t, valid, i)
		}
	})
	t.Run("Test_IsOWASPEmailInvalidLength", func(t *testing.T) {
		input := invalidLengthInput256
		valid := validator.IsOWASPEmail(input)
		assert.False(t, valid)
	})
	t.Run("Test_IsOWASPEmailInvalidChars", func(t *testing.T) {
		inputs := []string{"em@il@mail.com", "email @gmail.com"}
		for _, input := range inputs {
			valid := validator.IsOWASPEmail(input)
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
	t.Run("Test_IsPermissionResourceInvalidLength", func(t *testing.T) {
		input := invalidLengthInput256
		valid := validator.IsPermissionResource(input)
		assert.False(t, valid)
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
	t.Run("Test_IsPathInvalidLength", func(t *testing.T) {
		input := invalidLengthInput256
		valid := validator.IsPath(input)
		assert.False(t, valid)
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
	t.Run("Test_PasswordInvalidLength", func(t *testing.T) {
		inputs := []string{"Paswrd1", "p@swrd1", "P@SWRD1", "Paswrd!"}
		for i, input := range inputs {
			valid := validator.IsPassword(input)
			assert.False(t, valid, i)
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
