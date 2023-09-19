package info_test

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/s1berc0de/ozon-api-client/ozon/product/v3/info"

	"github.com/s1berc0de/ozon-api-client/internal/auth"
	"github.com/s1berc0de/ozon-api-client/internal/test"
	"github.com/stretchr/testify/require"
)

func TestStocks_Success(t *testing.T) {
	c := info.New(
		test.NewTestClient(
			auth.NewRoundTripper(
				test.RoundTripFunc(func(r *http.Request) *http.Response {
					require.Equal(t, "https://api-seller.ozon.ru/v3/product/info/stocks", test.FullURL(r))
					require.Equal(t, test.ApiKey, r.Header.Get(auth.APIKeyHeader))
					require.Equal(t, test.ClientID, r.Header.Get(auth.ClientIDHeader))
					require.Equal(t, `{"filter":{"offer_id":["136834"],"product_id":["214887921"],"visibility":"ALL"},"last_id":"","limit":100}`, test.Body(t, r))

					return &http.Response{
						StatusCode: http.StatusOK,
						Body: io.NopCloser(bytes.NewBufferString(`{
							"result": {
								"items": [
									{
										"product_id": 214887921,
										"offer_id": "136834",
										"stocks": [
											{
												"type": "fbs",
												"present": 170,
												"reserved": 0
											},
											{
												"type": "fbo",
												"present": 0,
												"reserved": 0
											}
										]
									}
								],
								"total": 1,
								"last_id": "anVsbA=="
							}
						}`)),
					}
				}),
				test.ClientID,
				test.ApiKey,
			),
		),
		"https://api-seller.ozon.ru/v3/product/info",
	)
	require.NotNil(t, c)

	resp, httpResp, err := c.Stocks(context.Background(), &info.StocksRequest{
		Filter: info.StocksRequestFilter{
			OfferID: []string{
				"136834",
			},
			ProductID: []string{
				"214887921",
			},
			Visibility: info.StocksRequestFilterVisibilityALL,
		},
		LastID: "",
		Limit:  100,
	})
	require.Nil(t, err)
	require.NotNil(t, httpResp)
	require.Equal(t, httpResp.StatusCode, http.StatusOK)
	require.EqualValues(t, &info.StocksResponse{
		Result: info.StocksResponseResult{
			Items: []info.StocksResponseResultItem{
				{
					ProductID: 214887921,
					OfferID:   "136834",
					Stocks: []info.StocksResponseResultItemStock{
						{
							Type:     "fbs",
							Present:  170,
							Reserved: 0,
						},
						{
							Type:     "fbo",
							Present:  0,
							Reserved: 0,
						},
					},
				},
			},
			Total:  1,
			LastID: "anVsbA==",
		},
	}, resp)
}
