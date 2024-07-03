package internal

import (
	"github.com/johnfercher/chaos/internal/model"
	"github.com/johnfercher/chaos/internal/regex"
)

type InterfaceInterpreter struct {
}

func NewInterfaceInterpreter() *InterfaceInterpreter {
	return &InterfaceInterpreter{}
}

func (i *InterfaceInterpreter) Interpret(file string) []*model.Interface {
	pkgName := regex.GetPackageName(file)
	imports := regex.GetImports(file)
	interfaces := regex.GetInterfaces(file)

	for i := 0; i < len(interfaces); i++ {
		interfaces[i].PackageName = pkgName
		interfaces[i].Imports = imports
	}

	return interfaces
}
