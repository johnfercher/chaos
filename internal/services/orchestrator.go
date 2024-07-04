package services

import (
	"errors"
	"regexp"

	"github.com/johnfercher/chaos/internal/core/models"
	"github.com/johnfercher/chaos/internal/core/services"
)

var ErrInterfaceNotFound = errors.New("interface_not_found_error")

type Orchestrator struct {
	file                 services.File
	interfaceInterpreter services.InterfaceInterpreter
	decoratorStrategy    services.DecoratorStrategy
}

func NewOrchestrator(file services.File, interfaceInterpreter services.InterfaceInterpreter,
	decoratorStrategy services.DecoratorStrategy,
) *Orchestrator {
	return &Orchestrator{
		file:                 file,
		interfaceInterpreter: interfaceInterpreter,
		decoratorStrategy:    decoratorStrategy,
	}
}

func (o *Orchestrator) Generate(filePath string, name string) error {
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

func (o *Orchestrator) generateAll(outputFile string, file string) error {
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

func (o *Orchestrator) generateOne(outputFile string, file string, name string) error {
	_interface, err := o.getInterface(file, name)
	if err != nil {
		return err
	}

	decorator := o.decoratorStrategy.Generate(_interface)
	return o.file.Write(outputFile, decorator)
}

func (o *Orchestrator) getInterface(file string, name string) (*models.Interface, error) {
	interfaces := o.interfaceInterpreter.Interpret(file)
	for _, _interface := range interfaces {
		if _interface.Name == name {
			return _interface, nil
		}
	}

	return nil, ErrInterfaceNotFound
}
