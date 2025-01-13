package dto

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func (u *UserPatchRequest) Validate() error {
	// Register the custom URI validation
	validate.RegisterValidation("uri", uriCustomValidate)

	err := validate.Struct(u)
	if err != nil {
		// Handle validation errors here
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "email":
				return fmt.Errorf("field '%s' must be a valid email address", err.Field())
			case "min":
				return fmt.Errorf("field '%s' must be at least %s characters long", err.Field(), err.Param())
			case "max":
				return fmt.Errorf("field '%s' cannot exceed %s characters", err.Field(), err.Param())
			case "uri":
				return fmt.Errorf("field '%s' must be a valid URI", err.Field())
			default:
				return fmt.Errorf("field '%s' failed validation on '%s' tag", err.Field(), err.Tag())
			}
		}
	}
	return nil
}

type User struct {
	Email           string `json:"email"`
	Name            string `json:"name"`
	UserImageUri    string `json:"userImageUri"`
	CompanyName     string `json:"companyName"`
	CompanyImageUri string `json:"companyImageUri"`
}

type UserRequest struct {
	UserID int `json:"id"`
}

type UserPatchRequest struct {
	Email           *string `json:"email" validate:"email"`
	Name            *string `json:"name" validate:"min=4,max=52"`
	UserImageUri    *string `json:"userImageUri" validate:"url"`
	CompanyName     *string `json:"companyName" validate:"min=4,max=52"`
	CompanyImageUri *string `json:"companyImageUri" validate:"url"`
}

// Create a custom URI validation function
func uriCustomValidate(fl validator.FieldLevel) bool {
	uri := fl.Field().String()
	if uri == "" {
		return true // It's valid if the field is empty (omitempty).
	}

	// Try to parse the URI
	_, err := url.ParseRequestURI(uri)
	if err != nil {
		return false // URI is not valid
	}

	// Add custom validation logic https or http:
	if !strings.HasPrefix(uri, "http://") && !strings.HasPrefix(uri, "https://") {
		return false // Custom check to make sure the URI starts with http:// or https://
	}

	return true
}
