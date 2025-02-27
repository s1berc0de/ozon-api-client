package v1

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/s1berc0de/ozon-api-client/internal/request"
	"github.com/s1berc0de/ozon-api-client/ozon/report/v1/postings"
	"net/http"
)

type SubRoutes struct {
	postings *postings.Postings
}

func (c SubRoutes) Postings() *postings.Postings {
	return c.postings
}

func New(
	h *http.Client,
	uri string,
) *Report {
	return &Report{
		h:   h,
		uri: uri,
		subRoutes: &SubRoutes{
			postings: postings.New(h, uri+"/postings"),
		},
	}
}

type Report struct {
	h   *http.Client
	uri string

	subRoutes *SubRoutes
}

func (r Report) SubRoutes() *SubRoutes {
	return r.subRoutes
}

func (r Report) Info(ctx context.Context, req *InfoRequest) (*InfoResponse, *http.Response, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, nil, errors.Wrap(err, "InfoRequest.Marshal")
	}

	reqCtx, err := http.NewRequestWithContext(ctx, http.MethodPost, r.uri+"/info", bytes.NewReader(b))
	if err != nil {
		return nil, nil, errors.Wrap(err, "InfoRequest.NewRequest")
	}

	return request.Send[InfoResponse](r.h, reqCtx, request.ContentTypeApplicationJson)
}
