package services

import (
	"github.com/johnfercher/chaos/struct/structcore/structmodels"
)

type DecoratorStrategy interface {
	Generate(*structmodels.Interface) string
}
