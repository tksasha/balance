package path

import (
	"maps"
	"net/url"
	"strconv"
	"time"

	"github.com/tksasha/balance/internal/common/currency"
)

const itemsPath = "/items"

func Items(params Params, current url.Values) string {
	path := url.URL{Path: itemsPath}

	desired := make(url.Values, len(params)+len(current))

	maps.Copy(desired, current)

	for key, value := range params {
		desired.Set(key, value)
	}

	if _, ok := desired["currency"]; !ok {
		desired.Set("currency", currency.GetCode(currency.Default))
	}

	if _, ok := desired["month"]; !ok {
		desired.Set("month", strconv.Itoa(int(time.Now().Month())))
	}

	if _, ok := desired["year"]; !ok {
		desired.Set("year", strconv.Itoa(time.Now().Year()))
	}

	path.RawQuery = desired.Encode()

	return path.String()
}

func EditItem(id int) string {
	return itemsPath + "/" + strconv.Itoa(id) + "/edit"
}

func UpdateItem(id int) string {
	return itemsPath + "/" + strconv.Itoa(id)
}

func NewItem() string {
	return itemsPath + "/new"
}
