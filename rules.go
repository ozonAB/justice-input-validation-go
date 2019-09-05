package validator

import (
	"github.com/asaskevich/govalidator"
	"github.com/pariz/gountries"
	"strconv"
	"time"
)

func IsTag(str string) bool {
	return rxTag.MatchString(str)
}

func IsLanguage(str string) bool {
	return rxLanguage.MatchString(str)
}

func IsTopic(str string) bool {
	return rxTopic.MatchString(str)
}

func IsDisplayName(str string) bool {
	return rxDisplayName.MatchString(str)
}

func IsPersonName(str string) bool {
	return rxPersonName.MatchString(str)
}

func IsUUID4WithoutHyphens(str string) bool {
	return rxUUIDv4WithoutHyphen.MatchString(str)
}

func IsEmail(str string) bool {
	return rxOWASPEmail.MatchString(str)
}

func IsPermissionResource(str string) bool {
	return rxResourcePermission.MatchString(str)
}

func IsPath(str string) bool {
	return rxPath.MatchString(str)
}

func IsURL(str string) bool {
	return rxOWASPUrl.MatchString(str)
}

func IsDateTime(str string) bool {
	return govalidator.IsRFC3339(str)
}

func IsDate(str string) bool {
	_, err := time.Parse(ISO8601TimeFormat, str)
	return err == nil
}

func IsJWT(str string) bool {
	return rxJWT.MatchString(str)
}

func IsNumeric(str string) bool {
	return govalidator.IsNumeric(str)
}

func IsIn(str string, params ...string) bool {
	return govalidator.IsIn(str, params...)
}

func IsLowerCase(str string) bool {
	return govalidator.IsLowerCase(str)
}

func IsPassword(str string) bool {
	valid, err := rxOWASPComplexPassword.MatchString(str)
	if err != nil {
		return false
	}
	return valid
}

func StringLength(input string, min int, max int) bool {
	return govalidator.StringLength(input, strconv.Itoa(min), strconv.Itoa(max))
}

func IsNotContainWhitespace(str string) bool {
	return !rxContainWhitespace.MatchString(str)
}

func IsContainWhitespace(str string) bool {
	return rxContainWhitespace.MatchString(str)
}

func IsCountry(str string) bool {
	_, err := gountries.New().FindCountryByAlpha(str)
	return err == nil
}

func IsCodeChallenge(str string) bool {
	return rxCodeChallenge.MatchString(str)
}
