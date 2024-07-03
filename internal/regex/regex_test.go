package regex_test

import (
	"github.com/johnfercher/chaos/internal/regex"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetScope(t *testing.T) {
	t.Run("when find interface, when find correctly", func(t *testing.T) {
		// Arrange
		content := `package example

import "context"

type Example interface {
	Add(ctx context.Context, id string) error
}`
		// Act
		text := regex.GetScope(content, regex.InterfaceName, regex.CloseBrackets)

		// Assert
		assert.Equal(t, "type Example interface {\n\tAdd(ctx context.Context, id string) error\n}\n", text)
	})
	t.Run("when find import, when find correctly", func(t *testing.T) {
		// Arrange
		content := `package example

import (
	"context"
)

type Example interface {
	Add(ctx context.Context, id string) error
}`
		// Act
		text := regex.GetScope(content, regex.MultiLineImports, regex.CloseParantesis)

		// Assert
		assert.Equal(t, "import (\n\t\"context\"\n)\n", text)
	})
}

func TestGetPackageName(t *testing.T) {
	t.Run("when has package name, then return it correctly", func(t *testing.T) {
		// Arrange
		content := `package example

import "context"

type Example interface {
	Add(ctx context.Context, id string) error
}`
		// Act
		pkg := regex.GetPackageName(content)

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
		pkg := regex.GetPackageName(content)

		// Assert
		assert.Empty(t, pkg)
	})
}

func TestGetImports(t *testing.T) {
	t.Run("when has one single import, then return correctly", func(t *testing.T) {
		// Arrange
		content := `package example

import "context"

type Example interface {
	Add(ctx context.Context, id string) error
}`
		// Act
		imports := regex.GetImports(content)

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
		imports := regex.GetImports(content)

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
		imports := regex.GetImports(content)

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
		imports := regex.GetImports(content)

		// Assert
		assert.Equal(t, []string{"context", "github.com/blabla"}, imports)
	})
}

func TestGetInterfaces(t *testing.T) {
	t.Run("when has one interface, the return correctly", func(t *testing.T) {
		// Arrange
		content := `package example

import "context"

type Example interface {
	Add(ctx context.Context, id string) error
}`
		// Act
		interfaces := regex.GetInterfaces(content)

		// Assert
		assert.Len(t, interfaces, 1)
		assert.Equal(t, "Example", interfaces[0].Name)
		assert.Len(t, interfaces[0].Methods, 1)
		assert.Equal(t, interfaces[0].Methods[0], "Add(ctx context.Context, id string) error")
	})
	t.Run("when has two interface, the return correctly", func(t *testing.T) {
		// Arrange
		content := `package example

import "context"

type Example interface {
	Add(ctx context.Context, id string) error
}

type Example2 interface {
	Add(ctx context.Context, id string) error
}
`
		// Act
		interfaces := regex.GetInterfaces(content)

		// Assert
		assert.Len(t, interfaces, 2)
		assert.Equal(t, "Example", interfaces[0].Name)
		assert.Len(t, interfaces[0].Methods, 1)
		assert.Equal(t, interfaces[0].Methods[0], "Add(ctx context.Context, id string) error")
		assert.Equal(t, "Example2", interfaces[1].Name)
		assert.Len(t, interfaces[1].Methods, 1)
		assert.Equal(t, interfaces[1].Methods[0], "Add(ctx context.Context, id string) error")
	})
}
