package requestvalidator

type errors map[string][]string

type Base struct {
	errors errors
}

func New() *Base {
	base := new(Base)

	base.errors = errors{}

	return base
}

func (b *Base) AddError(attribute, message string) {
	b.errors[attribute] = append(b.errors[attribute], message)
}

func (b *Base) HasErrors(attribute string) bool {
	return len(b.errors[attribute]) > 0 // TODO: what if there no required key
}

func (b *Base) GetErrors(attribute string) []string {
	return b.errors[attribute] // TODO: what if there no required key
}
