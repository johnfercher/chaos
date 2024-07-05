package chaos

var Decorator = `package {{package}}

{{imports}}
import "time"
import "math/rand/v2"

type {{implementation}} struct {
    inner {{interface}}
    errPercentage float64
    possibleErrors []error
    maxDelay time.Duration
}

func New{{implementation}}(inner {{interface}}) *{{implementation}} {
    return &{{implementation}}{
        inner: inner,
    }
}

func (i *{{implementation}}) WithErrors(percentage float64, errs ...error) *{{implementation}} {
    i.errPercentage = percentage
    i.possibleErrors = append(i.possibleErrors, errs...)
    return i
}

func (i *{{implementation}}) WithMaxDelay(maxDelay time.Duration) *{{implementation}} {
    i.maxDelay = maxDelay
    return i
}

func (i *{{implementation}}) getErr() error {
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

func (i *{{implementation}}) applyDelay() {
	if i.maxDelay == 0 {
		return
	}

	time.Sleep(i.maxDelay)
}

{{methods}}
`
