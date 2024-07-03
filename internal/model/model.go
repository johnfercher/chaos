package model

import "fmt"

type Parameter struct {
	Name string
	Type string
}

func (p *Parameter) Signature() string {
	if p.Name == "" {
		return p.Type
	}

	return fmt.Sprintf("%s %s", p.Name, p.Type)
}

func (p *Parameter) Call() string {
	if p.Name == "" {
		return ""
	}

	return p.Name
}

type Method struct {
	Name       string
	Parameters []Parameter
	Returns    []Parameter
}

func (m *Method) Signature() string {
	s := m.Name
	s += "("
	for i, parameter := range m.Parameters {
		if i < len(m.Parameters)-1 {
			s += parameter.Signature() + ", "
		} else {
			s += parameter.Signature()
		}
	}
	s += ") ("
	if len(m.Returns) == 1 {
		s += m.Returns[0].Signature()
	} else {
		for i, _return := range m.Returns {
			if i < len(m.Parameters)-1 {
				s += _return.Signature() + ", "
			} else {
				s += _return.Signature()
			}
		}
	}

	s += ")"
	return s
}

func (m *Method) Call() string {
	s := m.Name
	s += "("
	for i, parameter := range m.Parameters {
		if i < len(m.Parameters)-1 {
			s += parameter.Call() + ", "
		} else {
			s += parameter.Call()
		}
	}
	s += ")"
	return s
}

func (m *Method) CallReturn(err string) string {
	s := "return "
	if len(m.Returns) == 0 {
		return s
	}

	for _, _return := range m.Returns {
		if _return.Type == "error" {
			s += err
		} else {
			s += _return.Type
		}
	}

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
