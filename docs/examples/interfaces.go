package examples

import "context"

type SingleParameterWithoutReturn interface {
	Add(id string)
}

type TwoParametersWithoutReturn interface {
	Add(ctx context.Context, id string)
}

type SingleParameterWithSingleReturn interface {
	Add(id string) error
}

type SingleParameterWithTwoReturns interface {
	Add(id string) (context.Context, error)
}

type TwoParametersWithTwoReturns interface {
	Add(ctx context.Context, id string) (context.Context, error)
}
