package path

func Backoffice() string {
	return "/backoffice"
}

func BackofficeCashes() string {
	return Backoffice() + "/cashes"
}

func BackofficeCategories() string {
	return Backoffice() + "/categories"
}
