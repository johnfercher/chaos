package structmodels

import (
	"fmt"
)

type Interface struct {
	PackageName string
	Imports     Imports
	Name        string
	Methods     []Method
}

func (i *Interface) String() string {
	s := fmt.Sprintf("package %s\n", i.PackageName)

	s += i.Imports.String()

	s += fmt.Sprintf("type %s interface {\n", i.Name)
	for _, method := range i.Methods {
		s += "\t" + method.String() + "\n"
	}
	s += "}"
	return s
}
