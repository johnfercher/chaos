package regex

import (
	"fmt"
	"testing"
)

var content = `package structservices

import (
	"github.com/johnfercher/chaos/struct/structcore/structmodels"
)

type FunctionInterpreter interface {
	Interpret(file string) []*structmodels.Method
    Interpret[T](file string) []*structmodels.Method
}


type FunctionInterpreter struct {
}

func NewFunctionInterpreter() *FunctionInterpreter {
	return &FunctionInterpreter{}
}

func NewFunctionInterpreter(id string) *FunctionInterpreter {
	return &FunctionInterpreter{}
}

func NewFunctionInterpreter[T](id string) *FunctionInterpreter {
	return &FunctionInterpreter{}
}

func (f *FunctionInterpreter) Interpret(file string) []*structmodels.Method {
	//TODO implement me
	panic("implement me")
}
`

func TestGetFunctions(t *testing.T) {
	t.Run("", func(t *testing.T) {
		methods := GetFunctions(content)
		fmt.Println(methods)
	})
}
