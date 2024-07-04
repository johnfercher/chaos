package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMethod_Call(t *testing.T) {
	t.Run("0 parameters, 0 returns", func(t *testing.T) {
		// Arrange
		m := Method{
			Name: "method",
		}

		// Act
		s := m.Call()

		// Arrange
		assert.Equal(t, "i.inner.method()", s)
	})
	t.Run("1 parameter, 0 returns", func(t *testing.T) {
		// Arrange
		m := Method{
			Name: "method",
			Parameters: []Parameter{
				NewParameter("ctx", "context.Context"),
			},
		}

		// Act
		s := m.Call()

		// Arrange
		assert.Equal(t, "i.inner.method(ctx)", s)
	})
	t.Run("2 parameters, 0 returns", func(t *testing.T) {
		// Arrange
		m := Method{
			Name: "method",
			Parameters: []Parameter{
				NewParameter("ctx", "context.Context"),
				NewParameter("id", "string"),
			},
		}

		// Act
		s := m.Call()

		// Arrange
		assert.Equal(t, "i.inner.method(ctx, id)", s)
	})
	t.Run("3 parameters, 0 returns", func(t *testing.T) {
		// Arrange
		m := Method{
			Name: "method",
			Parameters: []Parameter{
				NewParameter("ctx", "context.Context"),
				NewParameter("id", "string"),
				NewParameter("value", "int"),
			},
		}

		// Act
		s := m.Call()

		// Arrange
		assert.Equal(t, "i.inner.method(ctx, id, value)", s)
	})
	t.Run("0 parameters, 1 return", func(t *testing.T) {
		// Arrange
		m := Method{
			Name: "method",
			Returns: []Parameter{
				NewParameter("result1", "string"),
			},
		}

		// Act
		s := m.Call()

		// Arrange
		assert.Equal(t, "return i.inner.method()", s)
	})
	t.Run("0 parameters, 2 returns", func(t *testing.T) {
		// Arrange
		m := Method{
			Name: "method",
			Returns: []Parameter{
				NewParameter("result1", "string"),
				NewParameter("result2", "int"),
			},
		}

		// Act
		s := m.Call()

		// Arrange
		assert.Equal(t, "return i.inner.method()", s)
	})
	t.Run("2 parameters, 2 returns", func(t *testing.T) {
		// Arrange
		m := Method{
			Name: "method",
			Parameters: []Parameter{
				NewParameter("ctx", "context.Context"),
				NewParameter("id", "string"),
			},
			Returns: []Parameter{
				NewParameter("result1", "string"),
				NewParameter("result2", "int"),
			},
		}

		// Act
		s := m.Call()

		// Arrange
		assert.Equal(t, "return i.inner.method(ctx, id)", s)
	})
}

func TestMethod_Signature(t *testing.T) {
	t.Run("0 parameters, 0 returns", func(t *testing.T) {
		// Arrange
		m := Method{
			Name: "method",
		}

		// Act
		s := m.Signature()

		// Arrange
		assert.Equal(t, "method()", s)
	})
	t.Run("1 parameter, 0 returns", func(t *testing.T) {
		// Arrange
		m := Method{
			Name: "method",
			Parameters: []Parameter{
				NewParameter("ctx", "context.Context"),
			},
		}

		// Act
		s := m.Signature()

		// Arrange
		assert.Equal(t, "method(ctx context.Context)", s)
	})
	t.Run("2 parameters, 0 returns", func(t *testing.T) {
		// Arrange
		m := Method{
			Name: "method",
			Parameters: []Parameter{
				NewParameter("ctx", "context.Context"),
				NewParameter("id", "string"),
			},
		}

		// Act
		s := m.Signature()

		// Arrange
		assert.Equal(t, "method(ctx context.Context, id string)", s)
	})
	t.Run("3 parameters, 0 returns", func(t *testing.T) {
		// Arrange
		m := Method{
			Name: "method",
			Parameters: []Parameter{
				NewParameter("ctx", "context.Context"),
				NewParameter("id", "string"),
				NewParameter("value", "int"),
			},
		}

		// Act
		s := m.Signature()

		// Arrange
		assert.Equal(t, "method(ctx context.Context, id string, value int)", s)
	})
	t.Run("0 parameters, 1 return", func(t *testing.T) {
		// Arrange
		m := Method{
			Name: "method",
			Returns: []Parameter{
				NewParameter("result1", "string"),
			},
		}

		// Act
		s := m.Signature()

		// Arrange
		assert.Equal(t, "method() string", s)
	})
	t.Run("0 parameters, 2 returns", func(t *testing.T) {
		// Arrange
		m := Method{
			Name: "method",
			Returns: []Parameter{
				NewParameter("result1", "string"),
				NewParameter("result2", "int"),
			},
		}

		// Act
		s := m.Signature()

		// Arrange
		assert.Equal(t, "method() (string, int)", s)
	})
	t.Run("2 parameters, 2 returns", func(t *testing.T) {
		// Arrange
		m := Method{
			Name: "method",
			Parameters: []Parameter{
				NewParameter("ctx", "context.Context"),
				NewParameter("id", "string"),
			},
			Returns: []Parameter{
				NewParameter("result1", "string"),
				NewParameter("result2", "int"),
			},
		}

		// Act
		s := m.Signature()

		// Arrange
		assert.Equal(t, "method(ctx context.Context, id string) (string, int)", s)
	})
}

func TestMethod_CallReturn(t *testing.T) {
	t.Run("0 parameters, 0 returns", func(t *testing.T) {
		// Arrange
		m := Method{
			Name: "method",
		}

		// Act
		s := m.CallReturn("err")

		// Arrange
		assert.Equal(t, "return", s)
	})
	t.Run("1 parameter, 0 returns", func(t *testing.T) {
		// Arrange
		m := Method{
			Name: "method",
			Parameters: []Parameter{
				NewParameter("ctx", "context.Context"),
			},
		}

		// Act
		s := m.CallReturn("err")

		// Arrange
		assert.Equal(t, "return", s)
	})
	t.Run("2 parameters, 0 returns", func(t *testing.T) {
		// Arrange
		m := Method{
			Name: "method",
			Parameters: []Parameter{
				NewParameter("ctx", "context.Context"),
				NewParameter("id", "string"),
			},
		}

		// Act
		s := m.CallReturn("err")

		// Arrange
		assert.Equal(t, "return", s)
	})
	t.Run("3 parameters, 0 returns", func(t *testing.T) {
		// Arrange
		m := Method{
			Name: "method",
			Parameters: []Parameter{
				NewParameter("ctx", "context.Context"),
				NewParameter("id", "string"),
				NewParameter("value", "int"),
			},
		}

		// Act
		s := m.CallReturn("err")

		// Arrange
		assert.Equal(t, "return", s)
	})
	t.Run("0 parameters, 1 return", func(t *testing.T) {
		// Arrange
		m := Method{
			Name: "method",
			Returns: []Parameter{
				NewParameter("result1", "string"),
			},
		}

		// Act
		s := m.CallReturn("err")

		// Arrange
		assert.Equal(t, `return ""`, s)
	})
	t.Run("0 parameters, 2 returns", func(t *testing.T) {
		// Arrange
		m := Method{
			Name: "method",
			Returns: []Parameter{
				NewParameter("result1", "string"),
				NewParameter("result2", "error"),
			},
		}

		// Act
		s := m.CallReturn("err")

		// Arrange
		assert.Equal(t, `return "", err`, s)
	})
	t.Run("2 parameters, 2 returns", func(t *testing.T) {
		// Arrange
		m := Method{
			Name: "method",
			Parameters: []Parameter{
				NewParameter("ctx", "context.Context"),
				NewParameter("id", "string"),
			},
			Returns: []Parameter{
				NewParameter("result1", "string"),
				NewParameter("result2", "error"),
			},
		}

		// Act
		s := m.CallReturn("err")

		// Arrange
		assert.Equal(t, `return "", err`, s)
	})
}
