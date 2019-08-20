package validator

import (
	"regexp"

	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.ParamTagMap["alphaNumeric"] = IsAlphaNumeric
	govalidator.ParamTagRegexMap["alphaNumeric"] = regexp.MustCompile("^alphaNumeric\\((\\w+)\\)$") //nolint
	govalidator.TagMap["tag"] = IsTag
	govalidator.TagMap["language"] = IsLanguage
	govalidator.TagMap["topic"] = IsTopic
	govalidator.TagMap["displayName"] = IsDisplayName
	govalidator.TagMap["userDisplayName"] = IsUserDisplayName
	govalidator.TagMap["uuid4WithoutHyphens"] = IsUUID4WithoutHyphens
	govalidator.TagMap["orderNumber"] = IsOrderNumber
	govalidator.TagMap["dockerImage"] = IsDockerImage
	govalidator.TagMap["permissionResource"] = IsPermissionResource
	govalidator.TagMap["path"] = IsPath
	govalidator.TagMap["url"] = IsURL
	govalidator.TagMap["memorySize"] = IsMemorySize
	govalidator.TagMap["time"] = IsTime
	govalidator.TagMap["dob"] = IsDate
	govalidator.TagMap["jwt"] = IsJWT
	govalidator.TagMap["password"] = IsPassword
	govalidator.TagMap["emailOWASP"] = IsOWASPEmail
}

// ValidateStruct is used to check the Struct based on the rule of each field
func ValidateStruct(s interface{}) (bool, error) {
	valid, err := govalidator.ValidateStruct(s)

	return valid, err
}

// BaseValidator is used to extend validate function in GoValidator
func BaseValidator(validationRule string, input string) bool {
	tagMap := govalidator.TagMap
	if _, ok := tagMap[validationRule]; !ok {
		return false
	}
	return tagMap[validationRule](input)
}
