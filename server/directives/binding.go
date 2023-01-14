package directives

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
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

func msgForTag(field validator.FieldError) string {
	switch field.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", field)
	case "email":
		return fmt.Sprintf("%s must be a valid email", field)
	case "gte":
		return fmt.Sprintf("%s must be greater than or equal to 1", field)
	default:
		return fmt.Sprintf("%s is invalid", field)
	}
}

func Binding(ctx context.Context, obj interface{}, next graphql.Resolver, constraint string) (interface{}, error) {
	// the parameter passed by the gqlgen will be validated through the directives using the validator
	val, err := next(ctx)
	if err != nil {
		panic(err)
	}

	fieldName := *graphql.GetPathContext(ctx).Field

	err = validate.Var(val, constraint)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)

		if fieldName == "city" || fieldName == "state" || fieldName == "country" {
			// check if the field is address and return err and transErr for its fields
			var address = obj.(map[string]interface{})

			for _, value := range address {
				err = validate.Var(value, constraint)
				if err != nil {
					for _, fe := range err.(validator.ValidationErrors) {
						return val, fmt.Errorf("%s %s", fe.Field(), msgForTag(fe))
					}
				}
			}
		}
		return val, fmt.Errorf("%s %s", fieldName, validationErrors[0].Translate(translation))

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
