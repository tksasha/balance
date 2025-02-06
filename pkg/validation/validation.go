package validation

import (
	"strconv"
	"time"

	"github.com/tksasha/calculator"
)

const (
	required = "is required"
	invalid  = "is invalid"
)

type Validation struct {
	Errors Errors
}

func New() *Validation {
	return &Validation{
		Errors: Errors{},
	}
}

func (v *Validation) HasErrors() bool {
	return v.Errors.exists()
}

func (v *Validation) Set(attribute, value string) {
	v.Errors.add(attribute, value)
}

func (v *Validation) Presence(attribute, value string) string {
	if value == "" {
		v.Errors.add(attribute, required)
	}

	return value
}

func (v *Validation) Integer(attribute, value string) int {
	if value == "" {
		return 0
	}

	digit, err := strconv.Atoi(value)
	if err != nil {
		v.Errors.add(attribute, invalid)

		return 0
	}

	return digit
}

func (v *Validation) Formula(attribute, formula string) (string, float64) {
	if formula == "" {
		v.Errors.add(attribute, required)

		return formula, 0.0
	}

	sum, err := calculator.Calculate(formula)
	if err != nil {
		v.Errors.add(attribute, invalid)

		return formula, 0.0
	}

	return formula, sum
}

func (v *Validation) Boolean(attribute, value string) bool {
	switch value {
	case "true":
		return true
	case "false", "":
		return false
	default:
		v.Errors.add(attribute, invalid)

		return false
	}
}

func (v *Validation) Date(attribute, value string) time.Time {
	if value == "" {
		v.Errors.add(attribute, required)

		return time.Time{}
	}

	date, err := time.Parse(time.DateOnly, value)
	if err != nil {
		v.Errors.add(attribute, invalid)

		return time.Time{}
	}

	return date
}
