package validator

import (
	"microservice/shared/dto"

	validator "gopkg.in/go-playground/validator.v9"
)

var validate = validator.New()

func ValidateStruct(s interface{}) []*dto.ErrorValidateResponse {
	var errors []*dto.ErrorValidateResponse
	err := validate.Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element dto.ErrorValidateResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
