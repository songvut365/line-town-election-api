package validation

import (
	"fmt"
	"line-town-election-api/model"
	"strings"

	"github.com/gobeam/stringy"

	"github.com/go-playground/validator/v10"
)

//Validation
var validate = validator.New()

func ValidInput(input interface{}) []*model.ErrorResponse {
	var errors []*model.ErrorResponse
	err := validate.Struct(input)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element model.ErrorResponse
			element.FailedField = formatFailedField(err.StructNamespace())
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func formatFailedField(s string) string {
	field := strings.Split(s, ".")[1]
	field = strings.ToLower(field)

	//convert to camel case
	str := stringy.New(field)
	camel := str.CamelCase()

	fmt.Println(camel)

	// to lower first letter
	first := strings.ToLower(string(camel[0]))
	field = first + camel[1:]

	return field
}
