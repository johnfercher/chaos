package services

import (
	"github.com/johnfercher/chaos/internal/core/models"
	"github.com/johnfercher/chaos/internal/regex"
	"strings"
)

type InterfaceInterpreter struct{}

func NewInterfaceInterpreter() *InterfaceInterpreter {
	return &InterfaceInterpreter{}
}

func (int *InterfaceInterpreter) Interpret(file string) []*models.Interface {
	pkgName := regex.GetPackageName(file)
	imports := regex.GetImports(file)
	interfaces := regex.GetInterfaces(file)

	for i := 0; i < len(interfaces); i++ {
		interfaces[i].PackageName = pkgName
		interfaces[i].Imports = int.getUsedImports(imports, interfaces[i].Methods)
	}

	return interfaces
}

func (int *InterfaceInterpreter) getUsedImports(allImports []models.Import, methods []models.Method) []models.Import {
	filtered := make(map[string]models.Import)

	for _, _import := range allImports {

		for _, method := range methods {
			for _, parameter := range method.Parameters {
				if strings.Contains(parameter.Type, _import.Alias) {
					filtered[_import.Alias] = _import
				}
			}
			for _, _return := range method.Returns {
				if strings.Contains(_return.Type, _import.Alias) {
					filtered[_import.Alias] = _import
				}
			}
		}

	}

	var list []models.Import
	for _, f := range filtered {
		list = append(list, f)
	}

	return list
}
