package handlers

import "net/http"

type CreateCashHandler struct {
	cashService CashService
}

func NewCreateCashHandler(cashService CashService) *CreateCashHandler {
	return &CreateCashHandler{
		cashService: cashService,
	}
}

func (h *CreateCashHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = h.handle(r)
}

func (h *CreateCashHandler) handle(r *http.Request) error {
	_ = r

	return nil
}
