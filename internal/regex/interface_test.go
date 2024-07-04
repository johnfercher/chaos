package regex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInterfaces(t *testing.T) {
	t.Run("when has one interface, the return correctly", func(t *testing.T) {
		// Arrange
		content := `package example

import "context"

type Example interface {
	Add(ctx context.Context, id string) error
}`
		// Act
		interfaces := GetInterfaces(content)

		// Assert
		assert.Len(t, interfaces, 1)
		assert.Equal(t, "Example", interfaces[0].Name)
		assert.Len(t, interfaces[0].Methods, 1)
		assert.Equal(t, interfaces[0].Methods[0].Signature(), "Add(ctx context.Context, id string) error")
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
		interfaces := GetInterfaces(content)

		// Assert
		assert.Len(t, interfaces, 2)
		assert.Equal(t, "Example", interfaces[0].Name)
		assert.Len(t, interfaces[0].Methods, 1)
		assert.Equal(t, interfaces[0].Methods[0].Signature(), "Add(ctx context.Context, id string) error")
		assert.Equal(t, "Example2", interfaces[1].Name)
		assert.Len(t, interfaces[1].Methods, 1)
		assert.Equal(t, interfaces[1].Methods[0].Signature(), "Add(ctx context.Context, id string) error")
	})
}
