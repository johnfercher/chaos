package structmodels

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewImport(t *testing.T) {
	t.Run("when import is builtin", func(t *testing.T) {
		// Act
		i := NewImport("context")

		// Assert
		assert.Equal(t, "context", i.Full)
		assert.Equal(t, "context", i.Alias)
	})
	t.Run("when import is not builtin", func(t *testing.T) {
		// Act
		i := NewImport("blabla/context")

		// Assert
		assert.Equal(t, "blabla/context", i.Full)
		assert.Equal(t, "context", i.Alias)
	})
}
