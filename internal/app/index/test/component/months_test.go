package component_test

import (
	"bytes"
	"testing"

	"github.com/tksasha/balance/internal/app/index/component"
	"github.com/tksasha/balance/internal/app/index/test/mocks"
	"github.com/tksasha/balance/internal/common/currency"
	"github.com/tksasha/balance/internal/common/paths/params"
	"go.uber.org/mock/gomock"
	"gotest.tools/v3/golden"
)

func TestMonths(t *testing.T) {
	controller := gomock.NewController(t)

	timeProviderMock := mocks.NewMockTimeProvider(controller)

	timeProviderMock.EXPECT().CurrentYear().Return(2015)

	component := component.New(timeProviderMock)

	testdata := map[currency.Currency]string{
		0:            "months.html",
		currency.UAH: "months-uah.html",
		currency.USD: "months-usd.html",
		currency.EUR: "months-eur.html",
	}

	months := map[int]bool{
		1:  false,
		2:  false,
		3:  false,
		4:  false,
		5:  false,
		6:  false,
		7:  false,
		8:  true,
		9:  false,
		10: false,
		11: false,
		12: false,
	}

	for currency, filename := range testdata {
		for month, flag := range months {
			timeProviderMock.EXPECT().IsCurrentMonth(month).Return(flag)
		}

		writer := bytes.NewBuffer([]byte{})

		params := params.New().WithCurrency(currency)

		if err := component.Months(params).Render(writer); err != nil {
			t.Fatalf("failed to render months: %v", err)
		}

		golden.Assert(t, writer.String(), filename)
	}
}
