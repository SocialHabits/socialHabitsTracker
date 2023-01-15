package directives

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/AntonioTrupac/socialHabitsTracker/graph/customTypes"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	validate    *validator.Validate
	translation ut.Translator
)

func init() {
	validate = validator.New()
	en := en.New()
	uni := ut.New(en, en)
	translation, _ = uni.GetTranslator("en")
	en_translations.RegisterDefaultTranslations(validate, translation)
}

func msgForTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "is required"
	case "email":
		return "must be a valid email address"
	case "min":
		return fmt.Sprintf("must have a minimum value of %s", fe.Param())
	case "max":
		return fmt.Sprintf("must have a maximum value of %s", fe.Param())
	case "len":
		return fmt.Sprintf("must have a length of %s", fe.Param())
	case "gte":
		return fmt.Sprintf("must be greater than or equal to %s", fe.Param())
	case "lte":
		return fmt.Sprintf("must be less than or equal to %s", fe.Param())
	case "eqfield":
		return fmt.Sprintf("must match the value of %s", fe.Param())
	case "uuid":
		return "must be a valid UUID"
	case "url":
		return "must be a valid URL"
	default:
		return "is invalid"
	}
}

var fieldValidators = map[string]func(obj map[string]interface{}, val interface{}) error{
	"address": validateAddress,
}

func validateAddress(obj map[string]interface{}, val interface{}) error {

	addressInput, ok := obj["address"].(map[string]interface{})

	address, _ := json.Marshal(addressInput)

	var addressInputType []*customTypes.AddressInput
	err := json.Unmarshal(address, &addressInputType)
	if err != nil {
		return err
	}

	fmt.Println(addressInputType, ok)
	if !ok || len(addressInputType) == 0 {
		return fmt.Errorf("addressInput is required")
	}
	for _, address := range addressInputType {
		err := validate.Struct(address)
		if err != nil {
			for _, fe := range err.(validator.ValidationErrors) {
				return fmt.Errorf("%s %s", fe.Field(), msgForTag(fe))
			}
		}
	}
	return nil
}

func Binding(ctx context.Context, obj interface{}, next graphql.Resolver, constraint string) (interface{}, error) {
	val, err := next(ctx)
	if err != nil {
		return val, err
	}
	fieldName := *graphql.GetPathContext(ctx).Field

	if validateFunc, ok := fieldValidators[fieldName]; ok {
		if err := validateFunc(obj.(map[string]interface{}), val); err != nil {
			return val, err
		}
	} else {
		err = validate.Var(val, constraint)
		if err != nil {
			validationErrors := err.(validator.ValidationErrors)
			return val, fmt.Errorf("%s %s", fieldName, validationErrors.Translate(translation))
		}
	}

	return val, nil
}

func ValidateAddTranslation(tag string, message string) {
	validate.RegisterTranslation(tag, translation, func(ut ut.Translator) error {
		return ut.Add(tag, message, true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(tag, fe.Field())

		return t
	})
}
