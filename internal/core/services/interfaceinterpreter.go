package services

import (
	"github.com/johnfercher/chaos/internal/core/models"
)

type InterfaceInterpreter interface {
	Interpret(file string) []*models.Interface
}
