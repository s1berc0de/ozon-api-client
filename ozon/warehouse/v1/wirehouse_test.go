package v1_test

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/s1berc0de/ozon-api-client/internal/auth"
	"github.com/s1berc0de/ozon-api-client/internal/test"
	v1 "github.com/s1berc0de/ozon-api-client/ozon/warehouse/v1"
	"github.com/stretchr/testify/require"
)

func TestHistory_Success(t *testing.T) {
	c := v1.New(
		test.NewTestClient(
			auth.NewRoundTripper(
				test.RoundTripFunc(func(r *http.Request) *http.Response {
					require.Equal(t, "https://api-seller.ozon.ru/v1/warehouse/list", test.FullURL(r))
					require.Equal(t, test.ApiKey, r.Header.Get(auth.APIKeyHeader))
					require.Equal(t, test.ClientID, r.Header.Get(auth.ClientIDHeader))
					require.Equal(t, `{}`, test.Body(t, r))

					return &http.Response{
						StatusCode: http.StatusOK,
						Body: io.NopCloser(bytes.NewBufferString(`{
							"result": [
								{
									"warehouse_id": 15588127982000,
									"name": "Proffi (Панорама Групп)",
									"is_rfbs": false
								},
								{
									"warehouse_id": 22142605386000,
									"name": "Склад на производственной",
									"is_rfbs": true
								},
								{
									"warehouse_id": 22208673494000,
									"name": "Тест 37349",
									"is_rfbs": true
								},
								{
									"warehouse_id": 22240462819000,
									"name": "Тест12",
									"is_rfbs": true
								}
							]
						}`)),
					}
				}),
				test.ClientID,
				test.ApiKey,
			),
		),
		"https://api-seller.ozon.ru/v1/warehouse",
	)
	require.NotNil(t, c)

	resp, httpResp, err := c.List(context.Background(), &v1.ListRequest{})
	require.Nil(t, err)
	require.NotNil(t, httpResp)
	require.Equal(t, httpResp.StatusCode, http.StatusOK)
	require.EqualValues(t, &v1.ListResponse{
		Result: []v1.ListResponseResult{
			{
				WarehouseID: 15588127982000,
				Name:        "Proffi (Панорама Групп)",
				IsRFBS:      false,
			},
			{
				WarehouseID: 22142605386000,
				Name:        "Склад на производственной",
				IsRFBS:      true,
			},
			{
				WarehouseID: 22208673494000,
				Name:        "Тест 37349",
				IsRFBS:      true,
			},
			{
				WarehouseID: 22240462819000,
				Name:        "Тест12",
				IsRFBS:      true,
			},
		},
	}, resp)
}
