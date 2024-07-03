package internal

import (
	"fmt"
	"github.com/johnfercher/chaos/internal/model"
	"strings"
)

type DecoratorGenerator struct {
	decoratorTemplate string
	methodTemplate    string
}

func NewDecoratorGenerator(decoratorTemplate string, methodTemplate string) *DecoratorGenerator {
	return &DecoratorGenerator{
		decoratorTemplate: decoratorTemplate,
		methodTemplate:    methodTemplate,
	}
}

func (d *DecoratorGenerator) Generate(_interface *model.Interface) string {
	template := strings.ReplaceAll(d.decoratorTemplate, "{{package}}", _interface.PackageName)
	template = strings.ReplaceAll(template, "{{imports}}", d.buildImports(_interface))
	template = strings.ReplaceAll(template, "{{implementation}}", fmt.Sprintf("%sChaos", _interface.Name))
	template = strings.ReplaceAll(template, "{{interface}}", _interface.Name)
	template = strings.ReplaceAll(template, "{{methods}}", d.buildMethods(_interface))

	return template
}

func (d *DecoratorGenerator) buildImports(_interface *model.Interface) string {
	s := `import (`
	for _, _import := range _interface.Imports {
		s += "\n\t" + fmt.Sprintf(`"%s"`, _import)
	}
	s += "\n)"
	return s
}

func (d *DecoratorGenerator) buildMethods(_interface *model.Interface) string {
	methods := []string{}
	for _, method := range _interface.Methods {
		template := strings.ReplaceAll(d.methodTemplate, "{{implementation}}", fmt.Sprintf("%sChaos", _interface.Name))
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
