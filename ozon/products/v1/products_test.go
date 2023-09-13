package v1_test

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/diphantxm/ozon-api-client/internal/auth"
	"github.com/diphantxm/ozon-api-client/internal/test"
	v1 "github.com/diphantxm/ozon-api-client/ozon/products/v1"
	"github.com/stretchr/testify/require"
)

func TestDelete_Success(t *testing.T) {
	c := v1.New(
		test.NewTestClient(
			auth.NewRoundTripper(
				test.RoundTripFunc(func(r *http.Request) *http.Response {
					require.Equal(t, "https://api-seller.ozon.ru/v1/products/geo-restrictions-catalog-by-filter", test.FullURL(r))
					require.Equal(t, test.ApiKey, r.Header.Get(auth.APIKeyHeader))
					require.Equal(t, test.ClientID, r.Header.Get(auth.ClientIDHeader))
					require.Equal(t, `{"filter":{"only_visible":true},"last_order_number":0,"limit":3}`, test.Body(t, r))

					return &http.Response{
						StatusCode: http.StatusOK,
						Body: io.NopCloser(bytes.NewBufferString(`{
							"restrictions": [
								{
									"id": "world",
									"name": "Весь Мир",
									"is_visible": true,
									"order_number": 1
								},
								{
									"id": "42fb1c32-0cfe-5c96-9fb5-7f8e8449f28c",
									"name": "Все города РФ",
									"is_visible": true,
									"order_number": 2
								},
								{
									"id": "moscow",
									"name": "Москва",
									"is_visible": true,
									"order_number": 3
								}
							]
						}`)),
					}
				}),
				test.ClientID,
				test.ApiKey,
			),
		),
		"https://api-seller.ozon.ru/v1/products",
	)
	require.NotNil(t, c)

	resp, httpResp, err := c.GEORestrictionsCatalogByFilter(context.Background(), &v1.GEORestrictionsCatalogByFilterRequest{
		Filter: v1.GEORestrictionsCatalogByFilterRequestFilter{
			OnlyVisible: true,
		},
		LastOrderNumber: 0,
		Limit:           3,
	})
	require.Nil(t, err)
	require.NotNil(t, httpResp)
	require.Equal(t, httpResp.StatusCode, http.StatusOK)
	require.EqualValues(t, &v1.GEORestrictionsCatalogByFilterResponse{
		Restrictions: []v1.GEORestrictionsCatalogByFilterResponseRestriction{
			{
				ID:          "world",
				Name:        "Весь Мир",
				IsVisible:   true,
				OrderNumber: 1,
			},
			{
				ID:          "42fb1c32-0cfe-5c96-9fb5-7f8e8449f28c",
				Name:        "Все города РФ",
				IsVisible:   true,
				OrderNumber: 2,
			},
			{
				ID:          "moscow",
				Name:        "Москва",
				IsVisible:   true,
				OrderNumber: 3,
			},
		},
	}, resp)
}
