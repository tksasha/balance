package category

type CreateRequest struct {
	Name          string
	Income        string
	Visible       string
	Supercategory string
}

type UpdateRequest struct {
	ID            string
	Name          string
	Income        string
	Visible       string
	Supercategory string
}
