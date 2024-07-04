package services

import (
	"github.com/johnfercher/chaos/internal/core/models"
)

type DecoratorStrategy interface {
	Generate(*models.Interface) string
}
