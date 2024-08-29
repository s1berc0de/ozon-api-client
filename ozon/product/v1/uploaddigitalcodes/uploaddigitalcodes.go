package uploaddigitalcodes

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
) *UploadDigitalCodes {
	return &UploadDigitalCodes{
		h:   h,
		uri: uri,
	}
}

type UploadDigitalCodes struct {
	h   *http.Client
	uri string
}

func (c UploadDigitalCodes) Info(ctx context.Context, req *InfoRequest) (*InfoResponse, *http.Response, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, nil, errors.Wrap(err, "InfoRequest.Marshal")
	}

	r, err := http.NewRequestWithContext(ctx, http.MethodPost, c.uri+"/info", bytes.NewReader(b))
	if err != nil {
		return nil, nil, errors.Wrap(err, "InfoRequest.NewRequest")
	}

	return request.Send[InfoResponse](c.h, r, request.ContentTypeApplicationJson)
}
