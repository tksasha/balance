package path

import (
	"strconv"
)

func Backoffice() string {
	return "/backoffice"
}

func BackofficeCashes() string {
	return Backoffice() + "/cashes"
}

func BackofficeCash(id int) string {
	return BackofficeCashes() + "/" + strconv.Itoa(id)
}

func BackofficeEditCash(id int) string {
	return BackofficeCash(id) + "/edit"
}

func BackofficeCategories() string {
	return Backoffice() + "/categories"
}
