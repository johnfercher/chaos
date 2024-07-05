package structservices

import "github.com/johnfercher/chaos/struct/structcore/structmodels"

type FunctionInterpreter interface {
	Interpret(file string) []structmodels.Method
}
