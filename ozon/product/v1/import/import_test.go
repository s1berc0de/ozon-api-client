package _import_test

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	_import "github.com/diphantxm/ozon-api-client/ozon/product/v1/import"

	"github.com/diphantxm/ozon-api-client/internal/auth"
	"github.com/diphantxm/ozon-api-client/internal/test"
	"github.com/stretchr/testify/require"
)

func TestInfo_Success(t *testing.T) {
	c := _import.New(
		test.NewTestClient(
			auth.NewRoundTripper(
				test.RoundTripFunc(func(r *http.Request) *http.Response {
					require.Equal(t, "https://api-seller.ozon.ru/v1/product/import/info", test.FullURL(r))
					require.Equal(t, test.ApiKey, r.Header.Get(auth.APIKeyHeader))
					require.Equal(t, test.ClientID, r.Header.Get(auth.ClientIDHeader))
					require.Equal(t, `{"task_id":"172549793"}`, test.Body(t, r))

					return &http.Response{
						StatusCode: http.StatusOK,
						Body: io.NopCloser(bytes.NewBufferString(`{
							"result": {
								"items": [
									{
										"offer_id": "143210608",
										"product_id": 137285792,
										"status": "imported",
										"errors": [ ]
									}
								],
								"total": 1
							}
						}`)),
					}
				}),
				test.ClientID,
				test.ApiKey,
			),
		),
		"https://api-seller.ozon.ru/v1/product/import",
	)
	require.NotNil(t, c)

	resp, httpResp, err := c.Info(context.Background(), &_import.InfoRequest{
		TaskID: "172549793",
	})
	require.Nil(t, err)
	require.NotNil(t, httpResp)
	require.Equal(t, httpResp.StatusCode, http.StatusOK)
	require.EqualValues(t, &_import.InfoResponse{
		Result: _import.InfoResponseResult{
			Items: []_import.InfoResponseResultItem{
				{
					OfferID:   "143210608",
					ProductID: 137285792,
					Status:    _import.InfoResponseResultItemStatusImported,
					Errors:    make([]_import.InfoResponseResultItemError, 0),
				},
			},
			Total: 1,
		},
	}, resp)
}
