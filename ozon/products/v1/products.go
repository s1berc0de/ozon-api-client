package v1

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
) *Products {
	return &Products{
		h:   h,
		uri: uri,
	}
}

type Products struct {
	h   *http.Client
	uri string
}

func (c Products) GEORestrictionsCatalogByFilter(ctx context.Context, req *GEORestrictionsCatalogByFilterRequest) (*GEORestrictionsCatalogByFilterResponse, *http.Response, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, nil, errors.Wrap(err, "GEORestrictionsCatalogByFilterRequest.Marshal")
	}

	r, err := http.NewRequestWithContext(ctx, http.MethodPost, c.uri+"/geo-restrictions-catalog-by-filter", bytes.NewReader(b))
	if err != nil {
		return nil, nil, errors.Wrap(err, "GEORestrictionsCatalogByFilterRequest.NewRequest")
	}

	return request.Send[GEORestrictionsCatalogByFilterResponse](c.h, r, request.ContentTypeApplicationJson)
}
