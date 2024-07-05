package services

import (
	"fmt"
	"github.com/johnfercher/chaos/deco/core/models"
	"github.com/johnfercher/chaos/struct/structcore/structmodels"
	"strings"
)

type DecoratorGenerator struct {
	decoratorTemplate string
	methodTemplate    string
	suffix            string
}

func NewDecoratorGenerator(suffix string, decoratorTemplate string, methodTemplate string) *DecoratorGenerator {
	return &DecoratorGenerator{
		suffix:            suffix,
		decoratorTemplate: decoratorTemplate,
		methodTemplate:    methodTemplate,
	}
}

func (d *DecoratorGenerator) Generate(_interface *structmodels.Interface) string {
	template := strings.ReplaceAll(d.decoratorTemplate, "{{package}}", _interface.PackageName)
	template = strings.ReplaceAll(template, "{{imports}}", d.buildImports(_interface))
	template = strings.ReplaceAll(template, "{{implementation}}", fmt.Sprintf("%s%s", _interface.Name, d.suffix))
	template = strings.ReplaceAll(template, "{{interface}}", _interface.Name)
	template = strings.ReplaceAll(template, "{{methods}}", d.buildMethods(_interface))

	return template
}

func (d *DecoratorGenerator) buildImports(_interface *structmodels.Interface) string {
	if len(_interface.Imports) == 0 {
		return ""
	}

	s := `import (`
	for _, _import := range _interface.Imports {
		s += "\n\t" + _import.Import()
	}
	s += "\n)"
	return s
}

func (d *DecoratorGenerator) buildMethods(_interface *structmodels.Interface) string {
	methods := []string{}
	for _, method := range _interface.Methods {
		m := models.NewMethod(method)
		template := strings.ReplaceAll(d.methodTemplate, "{{implementation}}", fmt.Sprintf("%s%s", _interface.Name, d.suffix))
		template = strings.ReplaceAll(template, "{{method_signature}}", m.Signature())
		template = strings.ReplaceAll(template, "{{method_call}}", m.Call())
		template = strings.ReplaceAll(template, "{{method_return}}", m.CallReturn("err"))
		methods = append(methods, template)
	}
	s := ""
	for _, method := range methods {
		s += method
		s += "\n"
	}
	return s
}
