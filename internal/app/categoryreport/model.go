package categoryreport

type Entity struct {
	CategoryName  string
	CategorySlug  string
	Sum           float64
	Supercategory int
}

type Entities []*Entity

type MappedEntities map[int]Entities
