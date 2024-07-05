package structmodels

type Method struct {
	Name       string
	Source     string
	Package    string
	Path       string
	Parameters []Parameter
	Returns    []Parameter
}

func (m *Method) String() string {
	s := ""
	if m.Source != "" {
		s += m.Source + " "
	}

	if m.Package != "" {
		s += m.Package + " "
	}

	if m.Path != "" {
		s += m.Path + " "
	}

	s += m.Name + "("
	for index, parameter := range m.Parameters {
		if index < len(m.Parameters)-1 {
			s += parameter.String() + ", "
		} else {
			s += parameter.String()
		}
	}
	s += ") "
	if len(m.Returns) == 0 {
		return s
	}
	if len(m.Returns) == 1 {
		s += m.Returns[0].String()
		return s
	}
	s += "("
	for index, _return := range m.Returns {
		if index < len(m.Returns)-1 {
			s += _return.String() + ", "
		} else {
			s += _return.String()
		}
	}
	s += ")"
	return s
}
