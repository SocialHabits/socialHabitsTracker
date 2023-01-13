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
		fmt.Println("validation errors", validationErrors, "translate validation errors", validationErrors.Translate(translation))
		transErr := fmt.Errorf("%s%+v", fieldName, validationErrors[0].Translate(translation))
		fmt.Println(transErr, validationErrors)
		return val, transErr
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
