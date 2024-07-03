package model

import "fmt"

type Parameter struct {
	Name string
	Type string
}

func (p *Parameter) String() string {
	return fmt.Sprintf("%s %s", p.Name, p.Type)
}

type Method struct {
	Name       string
	Parameters []Parameter
	Returns    []string
}

func (m *Method) String() string {
	s := m.Name
	s += "("
	for i, parameter := range m.Parameters {
		if i < len(m.Parameters)-1 {
			s += parameter.String() + ", "
		} else {
			s += parameter.String()
		}
	}
	s += ") ("
	for _, _return := range m.Returns {
		s += _return + ", "
	}
	s += ")"
	return s
}

type Interface struct {
	PackageName string
	Imports     []string
	Name        string
	Methods     []Method
}

func (i *Interface) String() string {
	s := fmt.Sprintf("package %s\n", i.PackageName)
	s += fmt.Sprintf("interface %s\n", i.Name)
	for _, _import := range i.Imports {
		s += fmt.Sprintf("import %s\n", _import)
	}
	for _, method := range i.Methods {
		s += fmt.Sprintf("method %s\n", method)
	}
	return s
}
