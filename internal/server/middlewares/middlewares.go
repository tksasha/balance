package middlewares

func New() []Middleware {
	// order matters:
	// 3 -> 2 -> 1 -> 0 for request
	// and
	// 0 -> 1 -> 2 -> 3 for response
	return []Middleware{
		&errorMiddleware{},
		&logMiddleware{},
		&currencyMiddleware{},
		&initMiddleware{},
	}
}
