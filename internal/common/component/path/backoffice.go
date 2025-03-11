package path

import "strconv"

func Backoffice() string {
	return "/backoffice"
}

func BackofficeCashes() string {
	return Backoffice() + "/cashes"
}

func BackofficeEditCash(id int) string {
	return Backoffice() + "/" + strconv.Itoa(id) + "/edit"
}

func BackofficeCategories() string {
	return Backoffice() + "/categories"
}
