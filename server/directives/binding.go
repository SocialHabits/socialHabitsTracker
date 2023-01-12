package directives

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/go-playground/validator/v10"
)

var (
	validate *validator.Validate
)

func init() {
	validate = validator.New()
}

func Binding(ctx context.Context, obj interface{}, next graphql.Resolver, constraint string) (interface{}, error) {
	// the parameter passed by the gqlgen will be validated through the directives using the validator
	val, err := next(ctx)
	if err != nil {
		panic(err)
	}

	err = validate.Var(val, constraint)
	if err != nil {
		return nil, err
	}

	return val, nil
}
