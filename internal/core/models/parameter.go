package models

import (
	"fmt"
	"github.com/johnfercher/chaos/struct/structcore/structmodels"
	"strings"
)

var nilZeroCases = []string{
	"*",
	"[]",
	"...",
	"[]*",
	"map[",
	"context.Context",
	"interface{}",
}

var zeroZeroCases = []string{
	"byte",
	"int",
	"int8",
	"int16",
	"int32",
	"int64",
	"uint",
	"uint8",
	"uint16",
	"uint32",
	"uint64",
	"float32",
	"float64",
}

var emptyZeroCases = []string{
	"string",
	"rune",
}

func getZeroValue(s string) string {
	for _, c := range nilZeroCases {
		if strings.Contains(s, c) {
			return "nil"
		}
	}

	for _, c := range zeroZeroCases {
		if strings.Contains(s, c) {
			return "0"
		}
	}

	for _, c := range emptyZeroCases {
		if strings.Contains(s, c) {
			return `""`
		}
	}

	return s + "{}"
}

type Parameter struct {
	Name      string
	Type      string
	ZeroValue string
}

func NewParameter(name string, _type string) Parameter {
	return Parameter{
		Name:      name,
		Type:      _type,
		ZeroValue: getZeroValue(_type),
	}
}

func NewFromParameter(parameter structmodels.Parameter) Parameter {
	return Parameter{
		Name:      parameter.Name,
		Type:      parameter.Type,
		ZeroValue: getZeroValue(parameter.Type),
	}
}

func (p *Parameter) NamedSignature() string {
	if p.Name == "" {
		return p.Type
	}

	return fmt.Sprintf("%s %s", p.Name, p.Type)
}

func (p *Parameter) UnnamedSignature() string {
	return p.Type
}

func (p *Parameter) Call() string {
	return p.Name
}
