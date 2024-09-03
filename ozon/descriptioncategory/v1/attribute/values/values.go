package values

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/andmetoo/ozon-api-client/internal/request"
	"github.com/pkg/errors"
	"net/http"
)

func New(
	h *http.Client,
	uri string,
) *Values {
	return &Values{
		h:   h,
		uri: uri,
	}
}

type Values struct {
	h   *http.Client
	uri string
}

func (v Values) Search(ctx context.Context, req *SearchRequest) (*SearchResponse, *http.Response, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, nil, errors.Wrap(err, "SearchRequest.Marshal")
	}

	r, err := http.NewRequestWithContext(ctx, http.MethodPost, v.uri+"/search", bytes.NewReader(b))
	if err != nil {
		return nil, nil, errors.Wrap(err, "SearchRequest.NewRequest")
	}

	return request.Send[SearchResponse](v.h, r, request.ContentTypeApplicationJson)
}
