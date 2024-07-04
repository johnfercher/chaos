package examples

import (
	"context"
)
import "time"
import "math/rand/v2"

type SingleParameterWithTwoReturnsChaos struct {
	inner          SingleParameterWithTwoReturns
	errPercentage  float64
	possibleErrors []error
	maxDelay       time.Duration
}

func NewSingleParameterWithTwoReturnsChaos(inner SingleParameterWithTwoReturns) *SingleParameterWithTwoReturnsChaos {
	return &SingleParameterWithTwoReturnsChaos{
		inner: inner,
	}
}

func (i *SingleParameterWithTwoReturnsChaos) WithErrors(percentage float64, errs ...error) *SingleParameterWithTwoReturnsChaos {
	i.errPercentage = percentage
	i.possibleErrors = append(i.possibleErrors, errs...)
	return i
}

func (i *SingleParameterWithTwoReturnsChaos) WithMaxDelay(maxDelay time.Duration) *SingleParameterWithTwoReturnsChaos {
	i.maxDelay = maxDelay
	return i
}

func (i *SingleParameterWithTwoReturnsChaos) getErr() error {
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

func (i *SingleParameterWithTwoReturnsChaos) applyDelay() {
	if i.maxDelay == 0 {
		return
	}

	time.Sleep(i.maxDelay)
}

func (i *SingleParameterWithTwoReturnsChaos) Add(id string) (context.Context, error) {
	i.applyDelay()

	err := i.getErr()
	if err != nil {
		return nil, err
	}

	return i.inner.Add(id)
}
