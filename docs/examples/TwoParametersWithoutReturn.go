package examples

import (
	"context"
)
import "time"
import "math/rand/v2"

type TwoParametersWithoutReturnChaos struct {
	inner          TwoParametersWithoutReturn
	errPercentage  float64
	possibleErrors []error
	maxDelay       time.Duration
}

func NewTwoParametersWithoutReturnChaos(inner TwoParametersWithoutReturn) *TwoParametersWithoutReturnChaos {
	return &TwoParametersWithoutReturnChaos{
		inner: inner,
	}
}

func (i *TwoParametersWithoutReturnChaos) WithErrors(percentage float64, errs ...error) *TwoParametersWithoutReturnChaos {
	i.errPercentage = percentage
	i.possibleErrors = append(i.possibleErrors, errs...)
	return i
}

func (i *TwoParametersWithoutReturnChaos) WithMaxDelay(maxDelay time.Duration) *TwoParametersWithoutReturnChaos {
	i.maxDelay = maxDelay
	return i
}

func (i *TwoParametersWithoutReturnChaos) getErr() error {
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

func (i *TwoParametersWithoutReturnChaos) applyDelay() {
	if i.maxDelay == 0 {
		return
	}

	time.Sleep(i.maxDelay)
}

func (i *TwoParametersWithoutReturnChaos) Add(ctx context.Context, id string) {
	i.applyDelay()

	err := i.getErr()
	if err != nil {
		return
	}

	i.inner.Add(ctx, id)
}
