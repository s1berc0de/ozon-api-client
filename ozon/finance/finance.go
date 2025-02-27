package finance

import (
	v3 "github.com/s1berc0de/ozon-api-client/ozon/finance/v3"
	"net/http"
)

func New(
	h *http.Client,
	uri string,
) *Finance {
	return &Finance{
		v3: v3.New(h, uri+"/v3/finance"),
	}
}

type Finance struct {
	v3 *v3.Finance
}

func (f *Finance) V3() *v3.Finance {
	return f.v3
}
