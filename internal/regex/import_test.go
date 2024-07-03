package regex

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetImports(t *testing.T) {
	t.Run("when has one single import, then return correctly", func(t *testing.T) {
		// Arrange
		content := `package example

import "context"

type Example interface {
	Add(ctx context.Context, id string) error
}`
		// Act
		imports := GetImports(content)

		// Assert
		assert.Equal(t, []string{"context"}, imports)
	})
	t.Run("when has two single imports, then return correctly", func(t *testing.T) {
		// Arrange
		content := `package example

import "context"
import "github.com"

type Example interface {
	Add(ctx context.Context, id string) error
}`
		// Act
		imports := GetImports(content)

		// Assert
		assert.Equal(t, []string{"context", "github.com"}, imports)
	})
	t.Run("when has one multiple imports, then return correctly", func(t *testing.T) {
		// Arrange
		content := `package example

import (
	"context"
)

type Example interface {
	Add(ctx context.Context, id string) error
}`
		// Act
		imports := GetImports(content)

		// Assert
		assert.Equal(t, []string{"context"}, imports)
	})
	t.Run("when has one multiple imports, then return correctly", func(t *testing.T) {
		// Arrange
		content := `package example

import (
	"context"
	"github.com/blabla"
)

type Example interface {
	Add(ctx context.Context, id string) error
}`
		// Act
		imports := GetImports(content)

		// Assert
		assert.Equal(t, []string{"context", "github.com/blabla"}, imports)
	})
}
