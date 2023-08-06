package v1

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/diphantxm/ozon-api-client/ozon/product/v1/uploaddigitalcodes"

	"github.com/diphantxm/ozon-api-client/ozon/product/v1/update"

	"github.com/diphantxm/ozon-api-client/ozon/product/v1/info"

	"github.com/diphantxm/ozon-api-client/internal/request"
	"github.com/diphantxm/ozon-api-client/ozon/product/v1/attributes"
	_import "github.com/diphantxm/ozon-api-client/ozon/product/v1/import"
	"github.com/diphantxm/ozon-api-client/ozon/product/v1/pictures"

	"github.com/pkg/errors"
)

type SubRoutes struct {
	_import            *_import.Import
	attributes         *attributes.Attributes
	pictures           *pictures.Pictures
	info               *info.Info
	update             *update.Update
	uploadDigitalCodes *uploaddigitalcodes.UploadDigitalCodes
}

func (c SubRoutes) Import() *_import.Import {
	return c._import
}

func (c SubRoutes) Attributes() *attributes.Attributes {
	return c.attributes
}

func (c SubRoutes) Pictures() *pictures.Pictures {
	return c.pictures
}

func (c SubRoutes) Info() *info.Info {
	return c.info
}

func (c SubRoutes) Update() *update.Update {
	return c.update
}

func (c SubRoutes) UploadDigitalCodes() *uploaddigitalcodes.UploadDigitalCodes {
	return c.uploadDigitalCodes
}

func New(
	h *http.Client,
	uri string,
) *Product {
	return &Product{
		h:   h,
		uri: uri,
		subRoutes: &SubRoutes{
			_import:            _import.New(h, uri+"/import"),
			attributes:         attributes.New(h, uri+"/attributes"),
			pictures:           pictures.New(h, uri+"/pictures"),
			info:               info.New(h, uri+"/info"),
			update:             update.New(h, uri+"/update"),
			uploadDigitalCodes: uploaddigitalcodes.New(h, uri+"/upload_digital_codes"),
		},
	}
}

type Product struct {
	h   *http.Client
	uri string

	subRoutes *SubRoutes
}

func (c Product) SubRoutes() *SubRoutes {
	return c.subRoutes
}

func (c Product) ImportBySKU(ctx context.Context, req *ImportBySKURequest) (*ImportBySKUResponse, *http.Response, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, nil, errors.Wrap(err, "ImportBySKURequest.Marshal")
	}

	r, err := http.NewRequestWithContext(ctx, http.MethodPost, c.uri+"/import-by-sku", bytes.NewReader(b))
	if err != nil {
		return nil, nil, errors.Wrap(err, "ImportBySKURequest.NewRequest")
	}

	return request.Send[ImportBySKUResponse](c.h, r, request.ContentTypeApplicationJson)
}

func (c Product) RatingBySKU(ctx context.Context, req *RatingBySKURequest) (*RatingBySKUResponse, *http.Response, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, nil, errors.Wrap(err, "RatingBySKURequest.Marshal")
	}

	r, err := http.NewRequestWithContext(ctx, http.MethodPost, c.uri+"/rating-by-sku", bytes.NewReader(b))
	if err != nil {
		return nil, nil, errors.Wrap(err, "RatingBySKURequest.NewRequest")
	}

	return request.Send[RatingBySKUResponse](c.h, r, request.ContentTypeApplicationJson)
}

func (c Product) Archive(ctx context.Context, req *ArchiveRequest) (*ArchiveResponse, *http.Response, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, nil, errors.Wrap(err, "ArchiveRequest.Marshal")
	}

	r, err := http.NewRequestWithContext(ctx, http.MethodPost, c.uri+"/archive", bytes.NewReader(b))
	if err != nil {
		return nil, nil, errors.Wrap(err, "ArchiveRequest.NewRequest")
	}

	return request.Send[ArchiveResponse](c.h, r, request.ContentTypeApplicationJson)
}

func (c Product) UnArchive(ctx context.Context, req *UnArchiveRequest) (*UnArchiveResponse, *http.Response, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, nil, errors.Wrap(err, "UnArchiveRequest.Marshal")
	}

	r, err := http.NewRequestWithContext(ctx, http.MethodPost, c.uri+"/unarchive", bytes.NewReader(b))
	if err != nil {
		return nil, nil, errors.Wrap(err, "UnArchiveRequest.NewRequest")
	}

	return request.Send[UnArchiveResponse](c.h, r, request.ContentTypeApplicationJson)
}

func (c Product) UploadDigitalCodes(ctx context.Context, req *UploadDigitalCodesRequest) (*UploadDigitalCodesResponse, *http.Response, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, nil, errors.Wrap(err, "UploadDigitalCodesRequest.Marshal")
	}

	r, err := http.NewRequestWithContext(ctx, http.MethodPost, c.uri+"/upload_digital_codes", bytes.NewReader(b))
	if err != nil {
		return nil, nil, errors.Wrap(err, "UploadDigitalCodesRequest.NewRequest")
	}

	return request.Send[UploadDigitalCodesResponse](c.h, r, request.ContentTypeApplicationJson)
}
