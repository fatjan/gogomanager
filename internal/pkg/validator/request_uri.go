package validator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func StrictURLValidation(fl validator.FieldLevel) bool {
	urlPattern := `^(http|https)://([a-zA-Z0-9-]+\.)+[a-zA-Z]{2,}(/[\w\-.~:/?#[\]@!$&'()*+,;=%]*)?$`
	matched, _ := regexp.MatchString(urlPattern, fl.Field().String())
	return matched
}
