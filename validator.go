package validator

import (
	"github.com/asaskevich/govalidator"
	"github.com/pariz/gountries"
)

var gountriesQuery *gountries.Query

func init() {
	// setup countries firstly to avoid concurrent initialize
	gountriesQuery = gountries.New()

	govalidator.TagMap["tag"] = IsTag
	govalidator.TagMap["language"] = IsLanguage
	govalidator.TagMap["topic"] = IsTopic
	govalidator.TagMap["displayName"] = IsDisplayName
	govalidator.TagMap["personName"] = IsPersonName
	govalidator.TagMap["uuid4WithoutHyphens"] = IsUUID4WithoutHyphens
	govalidator.TagMap["permissionResource"] = IsPermissionResource
	govalidator.TagMap["permissionResourceWithUUID"] = IsPermissionResourceWithUUID
	govalidator.TagMap["path"] = IsPath
	govalidator.TagMap["url"] = IsURL
	govalidator.TagMap["uri"] = IsURI
	govalidator.TagMap["dateTime"] = IsDateTime
	govalidator.TagMap["date"] = IsDate
	govalidator.TagMap["jwt"] = IsJWT
	govalidator.TagMap["password"] = IsPassword
	govalidator.TagMap["email"] = IsEmail
	govalidator.TagMap["codeChallenge"] = IsCodeChallenge
	govalidator.TagMap["notContainWhitespace"] = IsNotContainWhitespace
	govalidator.TagMap["containWhitespace"] = IsContainWhitespace
	govalidator.TagMap["country"] = IsCountry
	govalidator.TagMap["namespace"] = IsNamespace
}

// ValidateStruct is used to check the Struct based on the rule of each field
func ValidateStruct(s interface{}) (bool, error) {
	return govalidator.ValidateStruct(s)
}

// BaseValidator is used to extend validate function in GoValidator
func BaseValidator(validationRule string, input string) bool {
	tagMap := govalidator.TagMap
	if _, ok := tagMap[validationRule]; !ok {
		return false
	}
	return tagMap[validationRule](input)
}
