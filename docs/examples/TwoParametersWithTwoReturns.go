package examples

import (
	"context"
)
import "time"
import "math/rand/v2"

type TwoParametersWithTwoReturnsChaos struct {
	inner          TwoParametersWithTwoReturns
	errPercentage  float64
	possibleErrors []error
	maxDelay       time.Duration
}

func NewTwoParametersWithTwoReturnsChaos(inner TwoParametersWithTwoReturns) *TwoParametersWithTwoReturnsChaos {
	return &TwoParametersWithTwoReturnsChaos{
		inner: inner,
	}
}

func (i *TwoParametersWithTwoReturnsChaos) WithErrors(percentage float64, errs ...error) *TwoParametersWithTwoReturnsChaos {
	i.errPercentage = percentage
	i.possibleErrors = append(i.possibleErrors, errs...)
	return i
}

func (i *TwoParametersWithTwoReturnsChaos) WithMaxDelay(maxDelay time.Duration) *TwoParametersWithTwoReturnsChaos {
	i.maxDelay = maxDelay
	return i
}

func (i *TwoParametersWithTwoReturnsChaos) getErr() error {
	if len(i.possibleErrors) == 0 {
		return nil
	}

	r := rand.IntN(100)
	if float64(r) > i.errPercentage {
		return nil
	}

	index := rand.IntN(len(i.possibleErrors) - 1)
	return i.possibleErrors[index]
}

func (i *TwoParametersWithTwoReturnsChaos) applyDelay() {
	if i.maxDelay == 0 {
		return
	}

	time.Sleep(i.maxDelay)
}

func (i *TwoParametersWithTwoReturnsChaos) Add(ctx context.Context, id string) (context.Context, error) {
	i.applyDelay()

	err := i.getErr()
	if err != nil {
		return nil, err
	}

	return i.inner.Add(ctx, id)
}
