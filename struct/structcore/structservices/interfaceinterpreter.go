package structservices

import (
	"github.com/johnfercher/chaos/struct/structcore/structmodels"
)

type InterfaceInterpreter interface {
	Interpret(file string) []*structmodels.Interface
}
