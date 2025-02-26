package postings

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/s1berc0de/ozon-api-client/internal/request"
	"net/http"
)

type Postings struct {
	h   *http.Client
	uri string
}

func New(
	h *http.Client,
	uri string,
) *Postings {
	return &Postings{
		h:   h,
		uri: uri,
	}
}

func (p Postings) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, *http.Response, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, nil, errors.Wrap(err, "CreateRequest.Marshal")
	}

	r, err := http.NewRequestWithContext(ctx, http.MethodPost, p.uri+"/create", bytes.NewReader(b))
	if err != nil {
		return nil, nil, errors.Wrap(err, "CreateRequest.NewRequest")
	}

	return request.Send[CreateResponse](p.h, r, request.ContentTypeApplicationJson)
}
