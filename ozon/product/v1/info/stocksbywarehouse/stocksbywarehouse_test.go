package stocksbywarehouse_test

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/diphantxm/ozon-api-client/ozon/product/v1/info/stocksbywarehouse"

	"github.com/diphantxm/ozon-api-client/internal/auth"
	"github.com/diphantxm/ozon-api-client/internal/test"
	"github.com/stretchr/testify/require"
)

func TestFBS_Success(t *testing.T) {
	c := stocksbywarehouse.New(
		test.NewTestClient(
			auth.NewRoundTripper(
				test.RoundTripFunc(func(r *http.Request) *http.Response {
					require.Equal(t, "https://api-seller.ozon.ru/v1/product/info/stocks-by-warehouse/fbs", test.FullURL(r))
					require.Equal(t, test.ApiKey, r.Header.Get(auth.APIKeyHeader))
					require.Equal(t, test.ClientID, r.Header.Get(auth.ClientIDHeader))
					require.Equal(t, `{"sku":["string"],"fbs_sku":["string"]}`, test.Body(t, r))

					return &http.Response{
						StatusCode: http.StatusOK,
						Body: io.NopCloser(bytes.NewBufferString(`{
							"result": [
								{
									"sku": 0,
									"fbs_sku": 0,
									"present": 0,
									"product_id": 0,
									"reserved": 0,
									"warehouse_id": 0,
									"warehouse_name": "string"
								}
							]
						}`)),
					}
				}),
				test.ClientID,
				test.ApiKey,
			),
		),
		"https://api-seller.ozon.ru/v1/product/info/stocks-by-warehouse",
	)
	require.NotNil(t, c)

	resp, httpResp, err := c.FBS(context.Background(), &stocksbywarehouse.FBSRequest{
		SKU: []string{
			"string",
		},
		FBSSKU: []string{
			"string",
		},
	})
	require.Nil(t, err)
	require.NotNil(t, httpResp)
	require.Equal(t, httpResp.StatusCode, http.StatusOK)
	require.EqualValues(t, &stocksbywarehouse.FBSResponse{
		Result: []stocksbywarehouse.FBSResponseResult{
			{
				SKU:           0,
				FBSSKU:        0,
				Present:       0,
				ProductID:     0,
				Reserved:      0,
				WarehouseID:   0,
				WarehouseName: "string",
			},
		},
	}, resp)
}
