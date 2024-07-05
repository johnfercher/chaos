package regex

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetParameter(t *testing.T) {
	t.Run("id string, value bool", func(t *testing.T) {
		// Arrange
		content := "id string, value bool"

		// Act
		parameters := GetParameters(content)

		// Assert
		assert.Equal(t, "id string", parameters[0].String())
		assert.Equal(t, "value bool", parameters[1].String())
	})
	t.Run("id, id2, id3 bool", func(t *testing.T) {
		// Arrange
		content := "id, id2, id3 bool"

		// Act
		parameters := GetParameters(content)

		// Assert
		assert.Equal(t, "id bool", parameters[0].String())
		assert.Equal(t, "id2 bool", parameters[1].String())
		assert.Equal(t, "id3 bool", parameters[2].String())
	})
}
