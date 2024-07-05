package structservices

import (
	"github.com/johnfercher/chaos/struct/structcore/structmodels"
)

type InterfaceInterpreter interface {
	Interpret(content string) []*structmodels.Interface
}
