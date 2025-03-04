package categoryreport

type Entity struct {
	CategoryName string
	CategorySlug string
	Sum          float64
}

type Entities []*Entity
