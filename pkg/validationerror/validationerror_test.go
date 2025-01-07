package validationerror_test

import (
	"slices"
	"testing"

	"github.com/tksasha/balance/pkg/validationerror"
	"gotest.tools/v3/assert"
)

func TestStringSliceMap_New(t *testing.T) {
	err := validationerror.New()

	assert.Assert(t, !err.Exists())
}

func TestStringSliceMap_NewWithMessage(t *testing.T) {
	err := validationerror.NewWithMessage("name", "can't be blank")

	assert.Assert(t, err.Has("name"))
	assert.Error(t, err, "name: can't be blank")
}

func TestStringSliceMap_Error(t *testing.T) {
	err := validationerror.NewWithMessage("name", "can't be blank")

	assert.Error(t, err, "name: can't be blank")
}

func TestStringSliceMap_Set(t *testing.T) {
	t.Run("when attribute is lowercase, it should remain unchanged", func(t *testing.T) {
		err := validationerror.New()

		err.Set("name", "is required")

		assert.Error(t, err, "name: is required")
	})

	t.Run("when attribute is uppercase, it should be stored in lowercase", func(t *testing.T) {
		err := validationerror.New()

		err.Set("Name", "is required")

		assert.Error(t, err, "name: is required")
	})
}

func TestStringSliceMap_Get(t *testing.T) {
	t.Run("when messages are empty, it should return empty array", func(t *testing.T) {
		err := validationerror.New()

		assert.Assert(t, slices.Equal(err.Get("name"), []string{}))
	})

	t.Run("when messages are not empty, it should return slice with those values", func(t *testing.T) {
		err := validationerror.NewWithMessage("name", "can't be blank")

		assert.Assert(t, slices.Equal(err.Get("name"), []string{"can't be blank"}))
	})
}

func TestStringSliceMap_Exists(t *testing.T) {
	t.Run("when there are no attributes, it should return false", func(t *testing.T) {
		err := validationerror.New()

		assert.Assert(t, !err.Exists())
	})

	t.Run("when there is attribute, it should return true", func(t *testing.T) {
		err := validationerror.NewWithMessage("name", "can't be blank")

		assert.Assert(t, err.Exists())
	})
}

func TestStringSliceMap_Has(t *testing.T) {
	err := validationerror.NewWithMessage("foo", "bar")

	t.Run("when there are messages for attribute, it should return true", func(t *testing.T) {
		assert.Assert(t, err.Has("foo"))
	})

	t.Run("when there no messages for attribute, it should return false", func(t *testing.T) {
		assert.Assert(t, !err.Has("jar"))
	})
}
