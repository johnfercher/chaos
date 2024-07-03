package example

import (
	"context"
)
import "time"
import "math/rand/v2"

type ExampleChaos struct {
	inner          Example
	errPercentage  float64
	possibleErrors []error
	maxDelay       time.Duration
}

func NewExampleChaos(inner Example) *ExampleChaos {
	return &ExampleChaos{
		inner: inner,
	}
}

func (i *ExampleChaos) WithErrors(percentage float64, errs ...error) *ExampleChaos {
	i.errPercentage = percentage
	i.possibleErrors = append(i.possibleErrors, errs...)
	return i
}

func (i *ExampleChaos) WithMaxDelay(maxDelay time.Duration) *ExampleChaos {
	i.maxDelay = maxDelay
	return i
}

func (i *ExampleChaos) getErr() error {
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

func (i *ExampleChaos) applyDelay() {
	if i.maxDelay == 0 {
		return
	}

	time.Sleep(i.maxDelay)
}

func (i *ExampleChaos) Add(ctx context.Context, id string) error {
	i.applyDelay()

	err := i.getErr()
	if err != nil {
		return err
	}

	return i.inner.Add(ctx, id)
}
