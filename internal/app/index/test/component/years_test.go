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

func TestYears(t *testing.T) {
	controller := gomock.NewController(t)

	timeProviderMock := mocks.NewMockTimeProvider(controller)

	timeProviderMock.EXPECT().CurrentYear().Return(2025)

	component := component.New(timeProviderMock)

	testdata := map[currency.Currency]string{
		0:            "years.html",
		currency.UAH: "years-uah.html",
		currency.USD: "years-usd.html",
		currency.EUR: "years-eur.html",
	}

	for currency, filename := range testdata {
		for n := range 10 {
			timeProviderMock.EXPECT().IsCurrentYear(2015 + n).Return(false)
		}

		timeProviderMock.EXPECT().IsCurrentYear(2025).Return(true)

		w := bytes.NewBuffer([]byte{})

		params := params.New().WithCurrency(currency)

		if err := component.Years(params).Render(w); err != nil {
			t.Fatal(err)
		}

		golden.Assert(t, w.String(), filename)
	}
}
