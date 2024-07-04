package services

import (
	"fmt"
	"strings"

	"github.com/johnfercher/chaos/internal/core/models"
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

func (d *DecoratorGenerator) Generate(_interface *models.Interface) string {
	template := strings.ReplaceAll(d.decoratorTemplate, "{{package}}", _interface.PackageName)
	template = strings.ReplaceAll(template, "{{imports}}", d.buildImports(_interface))
	template = strings.ReplaceAll(template, "{{implementation}}", fmt.Sprintf("%s%s", _interface.Name, d.suffix))
	template = strings.ReplaceAll(template, "{{interface}}", _interface.Name)
	template = strings.ReplaceAll(template, "{{methods}}", d.buildMethods(_interface))

	return template
}

func (d *DecoratorGenerator) buildImports(_interface *models.Interface) string {
	if len(_interface.Imports) == 0 {
		return ""
	}

	s := `import (`
	for _, _import := range _interface.Imports {
		s += "\n\t" + fmt.Sprintf(`"%s"`, _import.Import())
	}
	s += "\n)"
	return s
}

func (d *DecoratorGenerator) buildMethods(_interface *models.Interface) string {
	methods := []string{}
	for _, method := range _interface.Methods {
		template := strings.ReplaceAll(d.methodTemplate, "{{implementation}}", fmt.Sprintf("%s%s", _interface.Name, d.suffix))
		template = strings.ReplaceAll(template, "{{method_signature}}", method.Signature())
		template = strings.ReplaceAll(template, "{{method_call}}", method.Call())
		template = strings.ReplaceAll(template, "{{method_return}}", method.CallReturn("err"))
		methods = append(methods, template)
	}
	s := ""
	for _, method := range methods {
		s += method
		s += "\n"
	}
	return s
}
