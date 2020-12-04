# justice-input-validation-go

This is AccelByte Justice Golang Input Validation package. This package is extending functionality from [govalidator](https://github.com/asaskevich/govalidator) and add some additional rule to it.

## Usage

#### Importing package

```go
import validator "github.com/AccelByte/justice-input-validation-go"
```

#### Validating struct from request data model

We can validate a struct that have validation tag literal with validateStruct() method

```go
import validator "github.com/AccelByte/justice-input-validation-go"

// example of a request model
type requestModel struct {
  Name  string `valid:"displayName"`
  Email string `valid:"email"`
}

reqData = requestModel{
    Name: "Jhon Doe",
    Email: "jhon@mail.com",
}

// validating struct
if valid, err := validator.ValidateStruct(reqData); !valid || err {
    // do something when reqData is invalid
}
 
```

#### List of available validators with its corresponding function that defined in rules.go file

```go
"tag"                   : IsTag
"language"              : IsLanguage
"topic"                 : IsTopic
"displayName"           : IsDisplayName
"personName"            : IsPersonName
"uuid4WithoutHyphens"   : IsUUID4WithoutHyphens
"permissionResource"    : IsPermissionResource
"path"                  : IsPath
"url"                   : IsURL
"uri"                   : IsURI
"dateTime"              : IsDateTime
"date"                  : IsDate
"jwt"                   : IsJWT
"password"              : IsPassword
"email"                 : IsEmail
"codeChallenge"         : IsCodeChallenge
"notContainWhitespace"  : IsNotContainWhitespace
"containWhitespace"     : IsContainWhitespace
"country"               : IsCountry
```

And of course this package is not limiting the functionality that came from [govalidator](https://github.com/asaskevich/govalidator) package, you can use all available validation rules that supported by [govalidator](https://github.com/asaskevich/govalidator) package.