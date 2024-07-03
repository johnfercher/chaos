package regex

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMethod(t *testing.T) {
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
		assert.Equal(t, "ctx context.Context", m.Parameters[0].String())
		assert.Equal(t, "ctxPointer *context.Context", m.Parameters[1].String())
		assert.Equal(t, "pointer *string", m.Parameters[2].String())
		assert.Equal(t, "pointers []*string", m.Parameters[3].String())
		assert.Equal(t, "ctxs []context.Context", m.Parameters[4].String())
		assert.Equal(t, "id string", m.Parameters[5].String())
		assert.Equal(t, "arr []string", m.Parameters[6].String())
		assert.Equal(t, "ctxsPointers []*context.Context", m.Parameters[7].String())
		assert.Equal(t, "m map[string]string", m.Parameters[8].String())
		assert.Equal(t, "test ...string", m.Parameters[9].String())
	})
}

func TestGetBuiltinParameter(t *testing.T) {
	cases := []string{
		"Add(id string) error",
		"Add(id string, var ...string)",
		"Add(arr []string, id string)",
		"Add(ctx context.Context, ctxPointer *context.Context, pointer *string, pointerArr []*string, ctxs []context.Context, id string, arr []string, ctxsPointers []*context.Context, m map[string]string, var ...string) error",
	}

	for _, c := range cases {
		// Act
		param := GetBuiltinParameter(c)

		// Assert
		assert.Equal(t, "id string", param[0].String())
	}
}

func TestGetPackageParameter(t *testing.T) {
	cases := []string{
		"Add(ctx context.Context) error",
		"Add(id string, ctx context.Context)",
		"Add(ctx context.Context, var ...string)",
		"Add(ctx context.Context, ctxPointer *context.Context, pointer *string, pointerArr []*string, ctxs []context.Context, id string, arr []string, ctxsPointers []*context.Context, m map[string]string, var ...string) error",
	}

	for _, c := range cases {
		// Act
		param := GetPackageParameter(c)

		// Assert
		assert.Equal(t, "ctx context.Context", param[0].String())
	}
}

func TestGetBuiltinPointerParameter(t *testing.T) {
	cases := []string{
		"Add(pointer *string) error",
		"Add(pointer *string, var ...string)",
		"Add(arr []string, pointer *string)",
		"Add(ctx context.Context, ctxPointer *context.Context, pointer *string, pointerArr []*string, ctxs []context.Context, id string, arr []string, ctxsPointers []*context.Context, m map[string]string, var ...string) error",
	}

	for _, c := range cases {
		// Act
		param := GetBuiltinPointerParameter(c)

		// Assert
		assert.Equal(t, "pointer *string", param[0].String())
	}
}

func TestGetPackagePointerParameter(t *testing.T) {
	cases := []string{
		"Add(ctxPointer *context.Context) error",
		"Add(id string, ctxPointer *context.Context)",
		"Add(ctxPointer *context.Context, var ...string)",
		"Add(ctx context.Context, ctxPointer *context.Context, pointer *string, pointerArr []*string, ctxs []context.Context, id string, arr []string, ctxsPointers []*context.Context, m map[string]string, var ...string) error",
	}

	for _, c := range cases {
		// Act
		param := GetPackagePointerParameter(c)

		// Assert
		assert.Equal(t, "ctxPointer *context.Context", param[0].String())
	}
}

func TestGetBuiltinVarArgParameter(t *testing.T) {
	cases := []string{
		"Add(var ...string) error",
		"Add(id string, var ...string)",
		"Add(arr []string, var ...string)",
		"Add(ctx context.Context, ctxPointer *context.Context, pointer *string, pointerArr []*string, ctxs []context.Context, id string, arr []string, ctxsPointers []*context.Context, m map[string]string, var ...string) error"}

	for _, c := range cases {
		// Act
		param := GetBuiltinVarArgParameter(c)

		// Assert
		assert.Equal(t, "var ...string", param[0].String())
	}
}

func TestGetPackageVarArgParameter(t *testing.T) {
	cases := []string{
		"Add(varArgCtx ...context.Context) error",
		"Add(id string, varArgCtx ...context.Context)",
		"Add(ctx context.Context, ctxPointer *context.Context, pointer *string, pointerArr []*string, ctxs []context.Context, id string, arr []string, ctxsPointers []*context.Context, m map[string]string, varArgCtx ...context.Context) error"}

	for _, c := range cases {
		// Act
		param := GetPackageVarArgParameter(c)

		// Assert
		assert.Equal(t, "varArgCtx ...context.Context", param[0].String())
	}
}

func TestGetBuiltinVarArgPointerParameter(t *testing.T) {
	cases := []string{
		"Add(varPointer ...*string) error",
		"Add(id string, varPointer ...*string)",
		"Add(ctx context.Context, ctxPointer *context.Context, pointer *string, pointerArr []*string, ctxs []context.Context, id string, arr []string, ctxsPointers []*context.Context, m map[string]string, varPointer ...*string) error"}

	for _, c := range cases {
		// Act
		param := GetBuiltinVarArgPointerParameter(c)

		// Assert
		assert.Equal(t, "varPointer ...*string", param[0].String())
	}
}

func TestGetPackageVarArgPointerParameter(t *testing.T) {
	cases := []string{
		"Add(varPointer ...*context.Context) error",
		"Add(id string, varPointer ...*context.Context)",
		"Add(ctx context.Context, ctxPointer *context.Context, pointer *string, pointerArr []*string, ctxs []context.Context, id string, arr []string, ctxsPointers []*context.Context, m map[string]string, varPointer ...*context.Context) error"}

	for _, c := range cases {
		// Act
		param := GetPackageVarArgPointerParameter(c)

		// Assert
		assert.Equal(t, "varPointer ...*context.Context", param[0].String())
	}
}

func TestGetBuiltinArrayParameter(t *testing.T) {
	cases := []string{
		"Add(arr []string) error",
		"Add(id string, arr []string)",
		"Add(arr []string, var ...string)",
		"Add(ctx context.Context, ctxPointer *context.Context, pointer *string, pointerArr []*string, ctxs []context.Context, id string, arr []string, ctxsPointers []*context.Context, m map[string]string, var ...string) error"}

	for _, c := range cases {
		// Act
		param := GetBuiltinArrayParameter(c)

		// Assert
		assert.Equal(t, "arr []string", param[0].String())
	}
}

func TestGetPackageArrayParameter(t *testing.T) {
	cases := []string{
		"Add(ctxs []context.Context) error",
		"Add(id string, ctxs []context.Context)",
		"Add(ctxs []context.Context, var ...string)",
		"Add(ctx context.Context, ctxPointer *context.Context, pointer *string, pointerArr []*string, ctxs []context.Context, id string, arr []string, ctxsPointers []*context.Context, m map[string]string, var ...string) error"}

	for _, c := range cases {
		// Act
		param := GetPackageArrayParameter(c)

		// Assert
		assert.Equal(t, "ctxs []context.Context", param[0].String())
	}
}

func TestGetBuiltinArrayPointerParameter(t *testing.T) {
	cases := []string{
		"Add(pointerArr []*string) error",
		"Add(id string, pointerArr []*string)",
		"Add(pointerArr []*string, var ...string)",
		"Add(ctx context.Context, ctxPointer *context.Context, pointer *string, pointerArr []*string, ctxs []context.Context, id string, arr []string, ctxsPointers []*context.Context, m map[string]string, var ...string) error"}

	for _, c := range cases {
		// Act
		param := GetBuiltinArrayPointerParameter(c)

		// Assert
		assert.Equal(t, "pointerArr []*string", param[0].String())
	}
}

func TestGetPackageArrayPointerParameter(t *testing.T) {
	cases := []string{
		"Add(ctxsPointers []*context.Context) error",
		"Add(id string, ctxsPointers []*context.Context)",
		"Add(ctxsPointers []*context.Context, var ...string)",
		"Add(ctx context.Context, ctxPointer *context.Context, pointer *string, pointerArr []*string, ctxs []context.Context, id string, arr []string, ctxsPointers []*context.Context, m map[string]string, var ...string) error"}

	for _, c := range cases {
		// Act
		param := GetPackageArrayPointerParameter(c)

		// Assert
		assert.Equal(t, "ctxsPointers []*context.Context", param[0].String())
	}
}

func TestGetTrueBuiltinMapParameter(t *testing.T) {
	cases := []string{
		"Add(m map[string]string) error",
		"Add(id string, m map[string]string)",
		"Add(m map[string]string, var ...string)",
		"Add(ctx context.Context, ctxPointer *context.Context, pointer *string, pointerArr []*string, ctxs []context.Context, id string, arr []string, ctxsPointers []*context.Context, m map[string]string, var ...string) error"}

	for _, c := range cases {
		// Act
		param := GetTrueBuiltinMapParameter(c)

		// Assert
		assert.Equal(t, "m map[string]string", param[0].String())
	}
}

func TestGetKeyPackageMapBuiltinValueParameter(t *testing.T) {
	cases := []string{
		"Add(m2 map[context.Context]string) error",
		"Add(id string, m2 map[context.Context]string)",
		"Add(m2 map[context.Context]string, var ...string)",
		"Add(ctx context.Context, ctxPointer *context.Context, pointer *string, pointerArr []*string, ctxs []context.Context, id string, arr []string, ctxsPointers []*context.Context, m2 map[context.Context]string, var ...string) error"}

	for _, c := range cases {
		// Act
		param := GetKeyPackageMapBuiltinValueParameter(c)

		// Assert
		assert.Equal(t, "m2 map[context.Context]string", param[0].String())
	}
}

func TestGetKeyBuiltinMapPackageValueParameter(t *testing.T) {
	cases := []string{
		"Add(m3 map[string]context.Context) error",
		"Add(id string, m3 map[string]context.Context)",
		"Add(m3 map[string]context.Context, var ...string)",
		"Add(ctx context.Context, ctxPointer *context.Context, pointer *string, pointerArr []*string, ctxs []context.Context, id string, arr []string, ctxsPointers []*context.Context, m3 map[string]context.Context, var ...string) error"}

	for _, c := range cases {
		// Act
		param := GetKeyBuiltinMapPackageValueParameter(c)

		// Assert
		assert.Equal(t, "m3 map[string]context.Context", param[0].String())
	}
}

func TestGetTruePackageMapParameter(t *testing.T) {
	cases := []string{
		"Add(m4 map[context.Context]context.Context) error",
		"Add(id string, m4 map[context.Context]context.Context)",
		"Add(m4 map[context.Context]context.Context, var ...string)",
		"Add(ctx context.Context, ctxPointer *context.Context, pointer *string, pointerArr []*string, ctxs []context.Context, id string, arr []string, ctxsPointers []*context.Context, m4 map[context.Context]context.Context, var ...string) error"}

	for _, c := range cases {
		// Act
		param := GetTruePackageMapParameter(c)

		// Assert
		assert.Equal(t, "m4 map[context.Context]context.Context", param[0].String())
	}
}
