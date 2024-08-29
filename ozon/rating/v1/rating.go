package v1

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/andmetoo/ozon-api-client/internal/request"
	"github.com/pkg/errors"
)

func New(
	h *http.Client,
	uri string,
) *Rating {
	return &Rating{
		h:   h,
		uri: uri,
	}
}

type Rating struct {
	h   *http.Client
	uri string
}

func (c Rating) Summary(ctx context.Context, req *SummaryRequest) (*SummaryResponse, *http.Response, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, nil, errors.Wrap(err, "SummaryRequest.Marshal")
	}

	r, err := http.NewRequestWithContext(ctx, http.MethodPost, c.uri+"/summary", bytes.NewReader(b))
	if err != nil {
		return nil, nil, errors.Wrap(err, "SummaryRequest.NewRequest")
	}

	return request.Send[SummaryResponse](c.h, r, request.ContentTypeApplicationJson)
}

func (c Rating) History(ctx context.Context, req *HistoryRequest) (*HistoryResponse, *http.Response, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, nil, errors.Wrap(err, "HistoryRequest.Marshal")
	}

	r, err := http.NewRequestWithContext(ctx, http.MethodPost, c.uri+"/history", bytes.NewReader(b))
	if err != nil {
		return nil, nil, errors.Wrap(err, "HistoryRequest.NewRequest")
	}

	return request.Send[HistoryResponse](c.h, r, request.ContentTypeApplicationJson)
}
