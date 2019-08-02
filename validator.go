package validator

import (
	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.TagMap["displayName"] = IsDisplayName
	govalidator.TagMap["alphaNoWhiteSpace"] = IsAlphaNoWhiteSpace
	govalidator.TagMap["alphaNoWhiteSpaceWithDash"] = IsAlphaNoWhiteSpaceWithDash
}

func Validate(s interface{}) (bool, error) {
	valid, err := govalidator.ValidateStruct(s)

	return valid, err
}
