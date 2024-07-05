package structservices

import "github.com/johnfercher/chaos/struct/structcore/structmodels"

type FunctionInterpreter interface {
	Interpret(content string) []structmodels.Method
}
