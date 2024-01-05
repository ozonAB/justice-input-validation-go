// Copyright (c) 2019 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package validator

import (
	"regexp"

	"github.com/dlclark/regexp2"
)

const (
	UUID4WithoutHyphenPattern         = "^[0-9a-f]{12}4[0-9a-f]{3}[89ab][0-9a-f]{15}$"
	LanguagePattern                   = "^[a-zA-Z]+([-]{1}[a-zA-Z]+)*$"
	TagPattern                        = "^[a-zA-Z]+([_:-]{1}[a-zA-Z]+)*$"
	TopicPattern                      = "^[A-Z]+([_]{1}[A-Z]+)*$"
	DisplayNamePattern                = "^[a-zA-Z0-9]+(([',. -][a-zA-Z0-9])?[a-zA-Z0-9]*)*$"
	PersonNamePattern                 = "^[a-zA-Z]+(([',. -][a-zA-Z])?[a-zA-Z]*)*$"
	OWASPEmailPattern                 = `^[a-zA-Z0-9_+&*-]+(?:\.[a-zA-Z0-9_+&*-]+)*@(?:[a-zA-Z0-9-]+\.)+[a-zA-Z]{2,15}$`
	ResourcePermissionPattern         = `^[A-Z]+([:]{1}([A-Z]+|(({[a-zA-Z]+})|[a-zA-Z0-9]+|\*)))*$`
	ResourcePermissionPatternWithUUID = `^[A-Z]+([:]{1}([A-Z]+|(({[a-zA-Z0-9-]+})|[a-zA-Z0-9]+|\*)))*$`
	OWASPURLPattern                   = `^((((https?|ftps?|gopher|telnet|nntp):\/\/)|(mailto:|news:))(%[0-9A-Fa-f]{2}|[-()_.!~*';/?:@&=+$,A-Za-z0-9])+)([).!';/?:,][[:blank:]])?$`
	ContainWhitespacePattern          = `\s`
	JWTPattern                        = `^([A-Za-z0-9\-_~+\/]+[=]{0,2})\.([A-Za-z0-9\-_~+\/]+[=]{0,2})(?:\.([A-Za-z0-9\-_~+\/]+[=]{0,2}))?$`
	PathPattern                       = `^(\/[a-zA-Z0-9]+)+$`
	CodeChallengePattern              = `^[a-zA-Z0-9-._~]*$`
	ISO8601TimeFormat                 = "2006-01-02"
	OWASPComplexPasswordPattern       = `^(?:(?=.*\d)(?=.*[A-Z])(?=.*[a-z])|(?=.*\d)(?=.*[^A-Za-z0-9])(?=.*[a-z])|(?=.*[^A-Za-z0-9])(?=.*[A-Z])(?=.*[a-z])|(?=.*\d)(?=.*[A-Z])(?=.*[^A-Za-z0-9]))(?!.*(.)\1{2,})[A-Za-z0-9!~<>,;:_=?*+#."&§%°()\|\[\]\-\$\^\@\/]{8,32}$`
	NamespacePattern                  = `^[a-zA-Z0-9]{1,256}(\+[a-zA-Z0-9]{1,256})?$`
)

var (
	rxUUIDv4WithoutHyphen        = regexp.MustCompile(UUID4WithoutHyphenPattern)
	rxLanguage                   = regexp.MustCompile(LanguagePattern)
	rxTag                        = regexp.MustCompile(TagPattern)
	rxTopic                      = regexp.MustCompile(TopicPattern)
	rxDisplayName                = regexp.MustCompile(DisplayNamePattern)
	rxPersonName                 = regexp.MustCompile(PersonNamePattern)
	rxOWASPEmail                 = regexp.MustCompile(OWASPEmailPattern)
	rxResourcePermission         = regexp.MustCompile(ResourcePermissionPattern)
	rxResourcePermissionWithUUID = regexp.MustCompile(ResourcePermissionPatternWithUUID)
	rxOWASPUrl                   = regexp.MustCompile(OWASPURLPattern)
	rxContainWhitespace          = regexp.MustCompile(ContainWhitespacePattern)
	rxJWT                        = regexp.MustCompile(JWTPattern)
	rxPath                       = regexp.MustCompile(PathPattern)
	rxCodeChallenge              = regexp.MustCompile(CodeChallengePattern)
	rxOWASPComplexPassword       = regexp2.MustCompile(OWASPComplexPasswordPattern, 0)
	rxNamespace                  = regexp.MustCompile(NamespacePattern)
)
