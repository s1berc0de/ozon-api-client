package transaction

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/s1berc0de/ozon-api-client/internal/request"
	"net/http"
)

type Transaction struct {
	h   *http.Client
	uri string
}

func New(
	h *http.Client,
	uri string,
) *Transaction {
	return &Transaction{
		h:   h,
		uri: uri,
	}
}

func (t Transaction) List(ctx context.Context, req *ListRequest) (*ListResponse, *http.Response, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, nil, errors.Wrap(err, "ListRequest.Marshal")
	}

	r, err := http.NewRequestWithContext(ctx, http.MethodPost, t.uri+"/list", bytes.NewReader(b))
	if err != nil {
		return nil, nil, errors.Wrap(err, "ListRequest.NewRequest")
	}

	return request.Send[ListResponse](t.h, r, request.ContentTypeApplicationJson)
}
