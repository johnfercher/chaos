package models

import "fmt"

type Interface struct {
	PackageName string
	Imports     []Import
	Name        string
	Methods     []Method
}

func (i *Interface) String() string {
	s := fmt.Sprintf("package %s\n", i.PackageName)
	s += fmt.Sprintf("interface %s\n", i.Name)
	for _, _import := range i.Imports {
		s += fmt.Sprintf("import %s\n", _import.Full)
	}
	for _, method := range i.Methods {
		s += fmt.Sprintf("method %s\n", method)
	}
	return s
}
