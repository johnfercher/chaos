package regex

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMultiLineScope(t *testing.T) {
	t.Run("when find interface, when find correctly", func(t *testing.T) {
		// Arrange
		content := `package example

import "context"

type Example interface {
	Add(ctx context.Context, id string) error
}`
		// Act
		text := GetMultiLineScope(content, interfaceName, closeBrackets)

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
		text := GetMultiLineScope(content, multiLineImports, closeParantesis)

		// Assert
		assert.Equal(t, "import (\n\t\"context\"\n)\n", text)
	})
}

func TestGetSingleLineScope(t *testing.T) {
	t.Run("", func(t *testing.T) {
		// Arrange
		s := "Add(ctx context.Context, ctxPointer *context.Context, pointer *string, pointerArr []*string, ctxs []context.Context, id string, arr []string, ctxsPointers []*context.Context, m map[string]string, m4 map[context.Context]context.Context, var ...string) error"

		// Act
		scope := GetSingleLineScope(s, "(", ")")

		// Assert
		assert.Equal(t, "(ctx context.Context, ctxPointer *context.Context, pointer *string, pointerArr []*string, ctxs []context.Context, id string, arr []string, ctxsPointers []*context.Context, m map[string]string, m4 map[context.Context]context.Context, var ...string)", scope)
	})
}
