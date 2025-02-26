package v3

import (
	"github.com/s1berc0de/ozon-api-client/ozon/finance/v3/transaction"
	"net/http"
)

type SubRoutes struct {
	transaction *transaction.Transaction
}

func (s SubRoutes) Transaction() *transaction.Transaction {
	return s.transaction
}

func New(
	h *http.Client,
	uri string,
) *Finance {
	return &Finance{
		h:   h,
		uri: uri,
		subRoutes: &SubRoutes{
			transaction: transaction.New(h, uri+"/transaction"),
		},
	}
}

type Finance struct {
	h   *http.Client
	uri string

	subRoutes *SubRoutes
}

func (f Finance) SubRoutes() *SubRoutes {
	return f.subRoutes
}
