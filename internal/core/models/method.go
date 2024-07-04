package models

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
			s += parameter.NamedSignature() + ", "
		} else {
			s += parameter.NamedSignature()
		}
	}
	s += ")"
	if len(m.Returns) == 0 {
		return s
	}

	if len(m.Returns) == 1 {
		s += " " + m.Returns[0].UnnamedSignature()
		return s
	}

	s += " ("
	for i, _return := range m.Returns {
		if i < len(m.Returns)-1 {
			s += _return.UnnamedSignature() + ", "
		} else {
			s += _return.UnnamedSignature()
		}
	}
	s += ")"

	return s
}

func (m *Method) Call() string {
	s := ""
	if len(m.Returns) > 0 {
		s += "return "
	}

	s += "i.inner."
	s += m.Name
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
	s := "return"
	if len(m.Returns) == 0 {
		return s
	}

	s += " "

	for index, _return := range m.Returns {
		if _return.Type == "error" {
			s += err
		} else {
			s += _return.ZeroValue
		}
		if index < len(m.Returns)-1 {
			s += ", "
		}
	}

	return s
}
