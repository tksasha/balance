package middlewares

func New() []Middleware {
	return []Middleware{
		newCurrencyMiddleware(),
	}
}
