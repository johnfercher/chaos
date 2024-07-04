package regex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNamedMethod(t *testing.T) {
	t.Run("when name is simple, then map correctly", func(t *testing.T) {
		// Arrange
		s := "Add(ctx context.Context, id string) error"

		// Act
		m := GetMethod(s)

		// Assert
		assert.Equal(t, "Add", m.Name)
	})
	t.Run("when name has generics 1, then map correctly", func(t *testing.T) {
		// Arrange
		s := "Add[T](ctx context.Context, id string) error"

		// Act
		m := GetMethod(s)

		// Assert
		assert.Equal(t, "Add[T]", m.Name)
	})
	t.Run("when name has generics 1, then map correctly", func(t *testing.T) {
		// Arrange
		s := "Add[T any](ctx context.Context, id string) error"

		// Act
		m := GetMethod(s)

		// Assert
		assert.Equal(t, "Add[T any]", m.Name)
	})
	t.Run("when has different parameters, then map correctly", func(t *testing.T) {
		// Arrange
		s := "Add(ctx context.Context, ctxPointer *context.Context, pointer *string, pointers []*string, ctxs []context.Context, id string, arr []string, ctxsPointers []*context.Context, m map[string]string, test ...string) error"

		// Act
		m := GetMethod(s)

		// Assert
		assert.Len(t, m.Parameters, 10)
		/*assert.Equal(t, "ctx context.Context", m.Parameters[0].NamedSignature())
		assert.Equal(t, "ctxPointer *context.Context", m.Parameters[1].NamedSignature())
		assert.Equal(t, "pointer *string", m.Parameters[2].NamedSignature())
		assert.Equal(t, "pointers []*string", m.Parameters[3].NamedSignature())
		assert.Equal(t, "ctxs []context.Context", m.Parameters[4].NamedSignature())
		assert.Equal(t, "id string", m.Parameters[5].NamedSignature())
		assert.Equal(t, "arr []string", m.Parameters[6].NamedSignature())
		assert.Equal(t, "ctxsPointers []*context.Context", m.Parameters[7].NamedSignature())
		assert.Equal(t, "m map[string]string", m.Parameters[8].NamedSignature())
		assert.Equal(t, "test ...string", m.Parameters[9].NamedSignature())*/
	})
	t.Run("when has one return, then map correctly", func(t *testing.T) {
		// Arrange
		s := "Add(test ...string) error"

		// Act
		m := GetMethod(s)

		// Assert
		assert.Len(t, m.Returns, 1)
		assert.Equal(t, "error", m.Returns[0].Type)
	})
	t.Run("when has one named return, then map correctly", func(t *testing.T) {
		// Arrange
		s := "Add(test ...string) (error)"

		// Act
		m := GetMethod(s)

		// Assert
		assert.Len(t, m.Returns, 1)
		assert.Equal(t, "error", m.Returns[0].Type)
	})
	t.Run("when has one named return, then map correctly", func(t *testing.T) {
		// Arrange
		s := "Add(test ...string) (err error)"

		// Act
		m := GetMethod(s)

		// Assert
		assert.Len(t, m.Returns, 1)
		assert.Equal(t, "error", m.Returns[0].Type)
		assert.Equal(t, "err", m.Returns[0].Name)
	})
	t.Run("when has one named return, then map correctly", func(t *testing.T) {
		// Arrange
		s := "Add(test ...string) (bool, error)"

		// Act
		m := GetMethod(s)

		// Assert
		assert.Len(t, m.Returns, 2)
		assert.Equal(t, "bool", m.Returns[0].Type)
		assert.Equal(t, "error", m.Returns[1].Type)
	})
	t.Run("when has one named return, then map correctly", func(t *testing.T) {
		// Arrange
		s := "Add(test ...string) (b bool, err error)"

		// Act
		m := GetMethod(s)

		// Assert
		assert.Len(t, m.Returns, 2)
		assert.Equal(t, "bool", m.Returns[0].Type)
		assert.Equal(t, "b", m.Returns[0].Name)
		assert.Equal(t, "error", m.Returns[1].Type)
		assert.Equal(t, "err", m.Returns[1].Name)
	})
}
