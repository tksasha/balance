package cash

type CreateRequest struct {
	Formula       string
	Name          string
	Supercategory string
	Favorite      string
}

type UpdateRequest struct {
	ID            string
	Formula       string
	Name          string
	Supercategory string
	Favorite      string
}
