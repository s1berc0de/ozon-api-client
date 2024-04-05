package tree

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
	"github.com/s1berc0de/ozon-api-client/internal/request"
)

func New(
	h *http.Client,
	uri string,
) *Tree {
	return &Tree{
		h:   h,
		uri: uri,
	}
}

type Tree struct {
	h   *http.Client
	uri string
}

func (c Tree) Tree(ctx context.Context, req *TreeRequest) (*TreeResponse, *http.Response, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, nil, errors.Wrap(err, "TreeRequest.Marshal")
	}

	r, err := http.NewRequestWithContext(ctx, http.MethodPost, c.uri, bytes.NewReader(b))
	if err != nil {
		return nil, nil, errors.Wrap(err, "TreeRequest.NewRequest")
	}

	return request.Send[TreeResponse](c.h, r, request.ContentTypeApplicationJson)
}
