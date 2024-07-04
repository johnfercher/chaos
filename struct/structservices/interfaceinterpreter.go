package structservices

import (
	regex2 "github.com/johnfercher/chaos/struct/regex"
	"github.com/johnfercher/chaos/struct/structcore/structmodels"
	"strings"
)

type InterfaceInterpreter struct{}

func NewInterfaceInterpreter() *InterfaceInterpreter {
	return &InterfaceInterpreter{}
}

func (int *InterfaceInterpreter) Interpret(file string) []*structmodels.Interface {
	pkgName := regex2.GetPackageName(file)
	imports := regex2.GetImports(file)
	interfaces := regex2.GetInterfaces(file)

	for i := 0; i < len(interfaces); i++ {
		interfaces[i].PackageName = pkgName
		interfaces[i].Imports = int.getUsedImports(imports, interfaces[i].Methods)
	}

	return interfaces
}

func (int *InterfaceInterpreter) getUsedImports(allImports []structmodels.Import, methods []structmodels.Method) []structmodels.Import {
	filtered := make(map[string]structmodels.Import)

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

	var list []structmodels.Import
	for _, f := range filtered {
		list = append(list, f)
	}

	return list
}
