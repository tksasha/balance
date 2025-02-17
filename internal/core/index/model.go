package index

type Index struct {
	Months []string
}

func New() *Index {
	months := []string{
		"Січень",
		"Лютий",
		"Березень",
		"Квітень",
		"Травень",
		"Червень",
		"Липень",
		"Серпень",
		"Вересень",
		"Жовтень",
		"Листопад",
		"Грудень",
	}

	return &Index{
		Months: months,
	}
}
