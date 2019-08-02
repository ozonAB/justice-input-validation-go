package validator

import (
	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.TagMap["displayName"] = IsDisplayName
}

func Validate(s interface{}) (bool, error) {
	valid, err := govalidator.ValidateStruct(s)

	return valid, err
}
