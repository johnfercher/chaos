package regex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPackageName(t *testing.T) {
	t.Run("when has package name, then return it correctly", func(t *testing.T) {
		// Arrange
		content := `package example

import "context"

type Example interface {
	Add(ctx context.Context, id string) error
}`
		// Act
		pkg := GetPackageName(content)

		// Assert
		assert.Equal(t, "example", pkg)
	})
	t.Run("when has not package name, then return empty", func(t *testing.T) {
		// Arrange
		content := `pkg example

import "context"

type Example interface {
	Add(ctx context.Context, id string) error
}`
		// Act
		pkg := GetPackageName(content)

		// Assert
		assert.Empty(t, pkg)
	})
}
