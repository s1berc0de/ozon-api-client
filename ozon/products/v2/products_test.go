package v2_test

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/andmetoo/ozon-api-client/internal/auth"
	"github.com/andmetoo/ozon-api-client/internal/test"
	v2 "github.com/andmetoo/ozon-api-client/ozon/products/v2"
	"github.com/stretchr/testify/require"
)

func TestDelete_Success(t *testing.T) {
	c := v2.New(
		test.NewTestClient(
			auth.NewRoundTripper(
				test.RoundTripFunc(func(r *http.Request) *http.Response {
					require.Equal(t, "https://api-seller.ozon.ru/v2/products/delete", test.FullURL(r))
					require.Equal(t, test.ApiKey, r.Header.Get(auth.APIKeyHeader))
					require.Equal(t, test.ClientID, r.Header.Get(auth.ClientIDHeader))
					require.Equal(t, `{"products":[{"offer_id":"033"}]}`, test.Body(t, r))

					return &http.Response{
						StatusCode: http.StatusOK,
						Body: io.NopCloser(bytes.NewBufferString(`{
							"status": [
								{
									"offer_id": "033",
									"is_deleted": true,
									"error": ""
								}
							]
						}`)),
					}
				}),
				test.ClientID,
				test.ApiKey,
			),
		),
		"https://api-seller.ozon.ru/v2/products",
	)
	require.NotNil(t, c)

	resp, httpResp, err := c.Delete(context.Background(), &v2.DeleteRequest{
		Products: []v2.DeleteRequestProduct{
			{
				OfferID: "033",
			},
		},
	})
	require.Nil(t, err)
	require.NotNil(t, httpResp)
	require.Equal(t, httpResp.StatusCode, http.StatusOK)
	require.EqualValues(t, &v2.DeleteResponse{
		Status: []v2.DeleteResponseStatus{
			{
				OfferID:   "033",
				IsDeleted: true,
				Error:     "",
			},
		},
	}, resp)
}

func TestStocks_Success(t *testing.T) {
	c := v2.New(
		test.NewTestClient(
			auth.NewRoundTripper(
				test.RoundTripFunc(func(r *http.Request) *http.Response {
					require.Equal(t, "https://api-seller.ozon.ru/v2/products/stocks", test.FullURL(r))
					require.Equal(t, test.ApiKey, r.Header.Get(auth.APIKeyHeader))
					require.Equal(t, test.ClientID, r.Header.Get(auth.ClientIDHeader))
					require.Equal(t, `{"stocks":[{"offer_id":"PH11042","product_id":313455276,"stock":100,"warehouse_id":22142605386000}]}`, test.Body(t, r))

					return &http.Response{
						StatusCode: http.StatusOK,
						Body: io.NopCloser(bytes.NewBufferString(`{
							"result": [
								{
									"warehouse_id": 22142605386000,
									"product_id": 118597312,
									"offer_id": "PH11042",
									"updated": true,
									"errors": []
								}
							]
						}`)),
					}
				}),
				test.ClientID,
				test.ApiKey,
			),
		),
		"https://api-seller.ozon.ru/v2/products",
	)
	require.NotNil(t, c)

	resp, httpResp, err := c.Stocks(context.Background(), &v2.StocksRequest{
		Stocks: []v2.StocksRequestStock{
			{
				OfferID:     "PH11042",
				ProductID:   313455276,
				Stock:       100,
				WarehouseID: 22142605386000,
			},
		},
	})
	require.Nil(t, err)
	require.NotNil(t, httpResp)
	require.Equal(t, httpResp.StatusCode, http.StatusOK)
	require.EqualValues(t, &v2.StocksResponse{
		Result: []v2.StocksResponseResult{
			{
				WarehouseID: 22142605386000,
				ProductID:   118597312,
				OfferID:     "PH11042",
				Updated:     true,
				Errors:      make([]v2.StocksResponseResultError, 0),
			},
		},
	}, resp)
}
