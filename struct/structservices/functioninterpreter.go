package structservices

import (
	"github.com/johnfercher/chaos/struct/regex"
	"github.com/johnfercher/chaos/struct/structcore/structmodels"
)

type FunctionInterpreter struct {
}

func NewFunctionInterpreter() *FunctionInterpreter {
	return &FunctionInterpreter{}
}

func (f *FunctionInterpreter) Interpret(content string) []structmodels.Method {
	return regex.GetFunctions(content)
}
