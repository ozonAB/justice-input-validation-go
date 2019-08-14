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

func ValidateStruct(s interface{}) (bool, error) {
	valid, err := govalidator.ValidateStruct(s)

	return valid, err
}
