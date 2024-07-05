package services

import (
	"errors"
	"github.com/johnfercher/chaos/deco/core/services"
	"github.com/johnfercher/chaos/struct/structcore/structmodels"
	"github.com/johnfercher/chaos/struct/structcore/structservices"
	"regexp"
)

var ErrInterfaceNotFound = errors.New("interface_not_found_error")

type GenerationOrchestrator struct {
	file                 structservices.File
	interfaceInterpreter structservices.InterfaceInterpreter
	decoratorStrategy    services.DecoratorStrategy
}

func NewGenerationOrchestrator(file structservices.File, interfaceInterpreter structservices.InterfaceInterpreter,
	decoratorStrategy services.DecoratorStrategy,
) *GenerationOrchestrator {
	return &GenerationOrchestrator{
		file:                 file,
		interfaceInterpreter: interfaceInterpreter,
		decoratorStrategy:    decoratorStrategy,
	}
}

func (o *GenerationOrchestrator) Generate(filePath string, name string) error {
	file, err := o.file.Read(filePath)
	if err != nil {
		return err
	}

	outputRegex := regexp.MustCompile(`\w+\.go`)

	if name == "" {
		outputFile := outputRegex.ReplaceAllString(filePath, "")
		return o.generateAll(outputFile, file)
	}

	outputFile := outputRegex.ReplaceAllString(filePath, name+".go")
	return o.generateOne(outputFile, file, name)
}

func (o *GenerationOrchestrator) generateAll(outputFile string, file string) error {
	interfaces := o.interfaceInterpreter.Interpret(file)

	for _, _interface := range interfaces {
		decorator := o.decoratorStrategy.Generate(_interface)
		err := o.file.Write(outputFile+_interface.Name+".go", decorator)
		if err != nil {
			return err
		}
	}

	return nil
}

func (o *GenerationOrchestrator) generateOne(outputFile string, file string, name string) error {
	_interface, err := o.getInterface(file, name)
	if err != nil {
		return err
	}

	decorator := o.decoratorStrategy.Generate(_interface)
	return o.file.Write(outputFile, decorator)
}

func (o *GenerationOrchestrator) getInterface(file string, name string) (*structmodels.Interface, error) {
	interfaces := o.interfaceInterpreter.Interpret(file)
	for _, _interface := range interfaces {
		if _interface.Name == name {
			return _interface, nil
		}
	}

	return nil, ErrInterfaceNotFound
}
