package validation_test

import (
	"testing"
	"time"

	"github.com/tksasha/balance/pkg/validation"
	"gotest.tools/v3/assert"
)

func TestGetterAndSetter(t *testing.T) {
	t.Run("returns empty string when no errors", func(t *testing.T) {
		validation := validation.New()

		assert.Assert(t, !validation.HasErrors())
	})

	t.Run("returns formatted error when set", func(t *testing.T) {
		validation := validation.New()

		validation.Set("name", "must be unique")

		assert.Error(t, validation.Errors, "name: must be unique")
	})
}

func TestValidatePresence(t *testing.T) {
	t.Run("add error for blank value", func(t *testing.T) {
		validate := validation.New()

		_ = validate.Presence("name", "")

		assert.Error(t, validate.Errors, "name: is required")
	})

	t.Run("passes when value is present", func(t *testing.T) {
		validate := validation.New()

		_ = validate.Presence("name", "Bruce Wayne")

		assert.Assert(t, !validate.HasErrors())
	})
}

func TestValidateInteger(t *testing.T) {
	t.Run("add error for invalid value", func(t *testing.T) {
		validate := validation.New()

		res := validate.Integer("age", "abc")

		assert.Equal(t, res, 0)
		assert.Error(t, validate.Errors, "age: is invalid")
	})

	t.Run("converts string to integer when value is valid", func(t *testing.T) {
		validate := validation.New()

		res := validate.Integer("age", "33")

		assert.Equal(t, res, 33)
		assert.Assert(t, !validate.HasErrors())
	})
}

func TestValidateFormula(t *testing.T) {
	t.Run("add error for blank value", func(t *testing.T) {
		validate := validation.New()

		formula, res := validate.Formula("formula", "")

		assert.Equal(t, formula, "")
		assert.Equal(t, res, 0.0)
		assert.Error(t, validate.Errors, "formula: is required")
	})

	t.Run("add error for invalid value", func(t *testing.T) {
		validate := validation.New()

		formula, res := validate.Formula("formula", "abc")

		assert.Equal(t, formula, "abc")
		assert.Equal(t, res, 0.0)
		assert.Error(t, validate.Errors, "formula: is invalid")
	})

	t.Run("calculate result when value is valid", func(t *testing.T) {
		validate := validation.New()

		formula, res := validate.Formula("formula", "2+3")

		assert.Equal(t, formula, "2+3")
		assert.Equal(t, res, 5.0)
		assert.Assert(t, !validate.HasErrors())
	})
}

func TestValidateBoolean(t *testing.T) {
	t.Run("returns true for 'true' string", func(t *testing.T) {
		validate := validation.New()

		res := validate.Boolean("visible", "true")

		assert.Equal(t, res, true)
		assert.Assert(t, !validate.HasErrors())
	})

	t.Run("returns false for 'false' string", func(t *testing.T) {
		validate := validation.New()

		res := validate.Boolean("visible", "false")

		assert.Equal(t, res, false)
		assert.Assert(t, !validate.HasErrors())
	})

	t.Run("returns false for empty string", func(t *testing.T) {
		validate := validation.New()

		res := validate.Boolean("visible", "")

		assert.Equal(t, res, false)
		assert.Assert(t, !validate.HasErrors())
	})

	t.Run("adds error for invalid value", func(t *testing.T) {
		validate := validation.New()

		res := validate.Boolean("visible", "xxx")

		assert.Equal(t, res, false)
		assert.Error(t, validate.Errors, "visible: is invalid")
	})
}

func TestValidateDate(t *testing.T) {
	t.Run("returns error when value is blank", func(t *testing.T) {
		validate := validation.New()

		res := validate.Date("date", "")

		assert.Assert(t, res.IsZero())
		assert.Error(t, validate.Errors, "date: is required")
	})

	t.Run("returns error when value is invalid", func(t *testing.T) {
		validate := validation.New()

		res := validate.Date("date", "abc")

		assert.Assert(t, res.IsZero())
		assert.Error(t, validate.Errors, "date: is invalid")
	})

	t.Run("parse date when value is valid", func(t *testing.T) {
		validate := validation.New()

		res := validate.Date("date", "2025-02-04")

		assert.Equal(t, res, date(t, "2025-02-04"))
		assert.Assert(t, !validate.HasErrors())
	})
}

func date(t *testing.T, value string) time.Time {
	t.Helper()

	date, err := time.Parse(time.DateOnly, value)
	if err != nil {
		t.Fatalf("failed to parse date: %v", err)
	}

	return date
}
