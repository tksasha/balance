package cash

type CreateRequest struct {
	Formula       string
	Name          string
	Supercategory string
	Currency      string
}

type UpdateRequest struct {
	ID            string
	Formula       string
	Name          string
	Supercategory string
	Currency      string
}
