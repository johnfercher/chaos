package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParameter_Call(t *testing.T) {
	t.Run("when call, then return name", func(t *testing.T) {
		// Arrange
		parameter := Parameter{
			Name: "name",
			Type: "type",
		}

		// Act
		s := parameter.Call()

		// Assert
		assert.Equal(t, "name", s)
	})
}

func TestParameter_Signature(t *testing.T) {
	t.Run("when parameter has name, then return name parameter", func(t *testing.T) {
		// Arrange
		parameter := Parameter{
			Name: "name",
			Type: "type",
		}

		// Act
		s := parameter.NamedSignature()

		// Assert
		assert.Equal(t, "name type", s)
	})
	t.Run("when parameter hasn´t name, then return name parameter", func(t *testing.T) {
		// Arrange
		parameter := Parameter{
			Type: "type",
		}

		// Act
		s := parameter.NamedSignature()

		// Assert
		assert.Equal(t, "type", s)
	})
}
