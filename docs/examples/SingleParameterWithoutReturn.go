package examples

import "time"
import "math/rand/v2"

type SingleParameterWithoutReturnChaos struct {
	inner          SingleParameterWithoutReturn
	errPercentage  float64
	possibleErrors []error
	maxDelay       time.Duration
}

func NewSingleParameterWithoutReturnChaos(inner SingleParameterWithoutReturn) *SingleParameterWithoutReturnChaos {
	return &SingleParameterWithoutReturnChaos{
		inner: inner,
	}
}

func (i *SingleParameterWithoutReturnChaos) WithErrors(percentage float64, errs ...error) *SingleParameterWithoutReturnChaos {
	i.errPercentage = percentage
	i.possibleErrors = append(i.possibleErrors, errs...)
	return i
}

func (i *SingleParameterWithoutReturnChaos) WithMaxDelay(maxDelay time.Duration) *SingleParameterWithoutReturnChaos {
	i.maxDelay = maxDelay
	return i
}

func (i *SingleParameterWithoutReturnChaos) getErr() error {
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

func (i *SingleParameterWithoutReturnChaos) applyDelay() {
	if i.maxDelay == 0 {
		return
	}

	time.Sleep(i.maxDelay)
}

func (i *SingleParameterWithoutReturnChaos) Add(id string) {
	i.applyDelay()

	err := i.getErr()
	if err != nil {
		return
	}

	i.inner.Add(id)
}
