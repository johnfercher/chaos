package structmodels

type Parameter struct {
	Name string
	Type string
}

func NewParameter(name string, _type string) Parameter {
	return Parameter{
		Name: name,
		Type: _type,
	}
}
