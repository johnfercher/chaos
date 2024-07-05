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

func (f *FunctionInterpreter) Interpret(file string) []structmodels.Method {
	packageName := regex.GetPackageName(file)
	methods := regex.GetFunctions(file)

	for i := 0; i < len(methods); i++ {
		methods[i].Package = packageName
	}

	return methods
}
