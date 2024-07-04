package examples

import "time"
import "math/rand/v2"

type SingleParameterWithSingleReturnChaos struct {
	inner          SingleParameterWithSingleReturn
	errPercentage  float64
	possibleErrors []error
	maxDelay       time.Duration
}

func NewSingleParameterWithSingleReturnChaos(inner SingleParameterWithSingleReturn) *SingleParameterWithSingleReturnChaos {
	return &SingleParameterWithSingleReturnChaos{
		inner: inner,
	}
}

func (i *SingleParameterWithSingleReturnChaos) WithErrors(percentage float64, errs ...error) *SingleParameterWithSingleReturnChaos {
	i.errPercentage = percentage
	i.possibleErrors = append(i.possibleErrors, errs...)
	return i
}

func (i *SingleParameterWithSingleReturnChaos) WithMaxDelay(maxDelay time.Duration) *SingleParameterWithSingleReturnChaos {
	i.maxDelay = maxDelay
	return i
}

func (i *SingleParameterWithSingleReturnChaos) getErr() error {
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

func (i *SingleParameterWithSingleReturnChaos) applyDelay() {
	if i.maxDelay == 0 {
		return
	}

	time.Sleep(i.maxDelay)
}

func (i *SingleParameterWithSingleReturnChaos) Add(id string) error {
	i.applyDelay()

	err := i.getErr()
	if err != nil {
		return err
	}

	return i.inner.Add(id)
}
