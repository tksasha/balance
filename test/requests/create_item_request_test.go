package requests_test

import (
	"testing"

	"github.com/tksasha/balance/internal/requests"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestGetDate(t *testing.T) {
	t.Run("when date is zero it should return an empty string", func(t *testing.T) {
		req := requests.NewCreateItemRequest()

		req.Parse("", "", "", "")

		res := req.GetDate()

		exp := ""

		assert.Equal(t, res, exp)
	})

	t.Run("when date is not zero it should return formatted date", func(t *testing.T) {
		req := requests.NewCreateItemRequest()

		req.Parse("2024-09-30", "", "", "")

		res := req.GetDate()

		exp := "2024-09-30"

		assert.Equal(t, res, exp)
	})
}

func TestParse(t *testing.T) { //nolint:funlen
	t.Run("when date is blank it should set an error", func(t *testing.T) {
		req := requests.NewCreateItemRequest()

		req.Parse("", "", "", "")

		res := req.Errors.Get("date")

		assert.Assert(t, is.Contains(res, "is invalid"))
		assert.Assert(t, is.Contains(res, "is required"))
	})

	t.Run("when date is invalid it should set an error", func(t *testing.T) {
		req := requests.NewCreateItemRequest()

		req.Parse("1", "", "", "")

		res := req.Errors.Get("date")

		assert.Assert(t, is.Contains(res, "is invalid"))
		assert.Assert(t, is.Contains(res, "is required"))
	})

	t.Run("when formula is blank it should set an error", func(t *testing.T) {
		req := requests.NewCreateItemRequest()

		req.Parse("", "", "", "")

		res := req.Errors.Get("formula")

		assert.Assert(t, is.Contains(res, "is required"))
	})

	t.Run("when formula is invalid it should set an error", func(t *testing.T) {
		req := requests.NewCreateItemRequest()

		req.Parse("", "2/0", "", "")

		res := req.Errors.Get("formula")

		assert.Assert(t, is.Contains(res, "is invalid"))
	})

	t.Run("when formula is zero it should set an error", func(t *testing.T) {
		req := requests.NewCreateItemRequest()

		req.Parse("", "0", "", "")

		res := req.Errors.Get("formula")

		assert.Assert(t, is.Contains(res, "cannot be zero"))
	})

	t.Run("when formula is valid it should calculate sum", func(t *testing.T) {
		req := requests.NewCreateItemRequest()

		req.Parse("", "4.21+6.9", "", "")

		res := req.Errors.Check("formula")

		assert.Assert(t, !res)
		assert.Equal(t, req.Sum, 11.11)
	})

	t.Run("when category id is blank it should set an error", func(t *testing.T) {
		req := requests.NewCreateItemRequest()

		req.Parse("", "", "", "")

		res := req.Errors.Get("category-id")

		assert.Assert(t, is.Contains(res, "is required"))
	})

	t.Run("when category id is invalid it should set an error", func(t *testing.T) {
		req := requests.NewCreateItemRequest()

		req.Parse("", "", "abc", "")

		res := req.Errors.Get("category-id")

		assert.Assert(t, is.Contains(res, "is invalid"))
	})

	t.Run("when category id is zero it should set an error", func(t *testing.T) {
		req := requests.NewCreateItemRequest()

		req.Parse("", "", "0", "")

		res := req.Errors.Get("category-id")

		assert.Assert(t, is.Contains(res, "cannot be zero"))
	})

	t.Run("when description is not blank it should set this", func(t *testing.T) {
		req := requests.NewCreateItemRequest()

		req.Parse("", "", "", "lorem ipsum ...")

		exp := "lorem ipsum ..."

		assert.Equal(t, req.Description, exp)
	})
}
