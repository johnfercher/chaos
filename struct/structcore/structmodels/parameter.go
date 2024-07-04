package structmodels

import "fmt"

type Parameter struct {
	Name string
	Type string
}

func (p *Parameter) String() string {
	if p.Name == "" {
		return p.Type
	}

	return fmt.Sprintf("%s %s", p.Name, p.Type)
}

func NewParameter(name string, _type string) Parameter {
	return Parameter{
		Name: name,
		Type: _type,
	}
}
